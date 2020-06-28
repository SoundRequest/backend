package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SoundRequest/OAuth2Server/loader"

	"github.com/SoundRequest/OAuth2Server/checker"
	"github.com/SoundRequest/OAuth2Server/handler"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	/**
	* Initializer
	**/
	dbConfig, err := loader.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)
	fmt.Println("Connection To " + connectionInfo)
	db, err := handler.InitDB(dbConfig.DBType, connectionInfo)
	if err != nil {
		log.Fatal(err)
	}

	var tempQuery string

	db.QueryRow("SELECT 1").Scan(&tempQuery)
	fmt.Println(tempQuery)

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.Redirect("/protected", w, r)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		clientId, clientSecret, _ := handler.GenerateCredential(clientStore)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
	})

	http.HandleFunc("/protected", checker.CheckIsTokenValid(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Okay. You've authenticated successfully"))
	}, srv))

	log.Fatal(http.ListenAndServe(":9096", nil))
}
