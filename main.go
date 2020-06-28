package main

import (
	"fmt"
	"log"

	"github.com/SoundRequest/OAuth2Server/loader"

	"github.com/SoundRequest/OAuth2Server/handler"
)

func main() {
	/**
	* Initializer
	**/
	dbConfig, errLoadConfig := loader.LoadConfig("./config.json")
	if errLoadConfig != nil {
		log.Fatal(errLoadConfig)
	}
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Schema)
	fmt.Println("Connection To " + connectionInfo)
	var _, errInitDB = handler.InitDB(dbConfig.DBType, connectionInfo)
	if errInitDB != nil {
		log.Fatal(errInitDB)
	}
}
