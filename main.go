package main

import (
	"log"
	"net/http"

	"github.com/SoundRequest/OAuth2Server/RequestHandler"
)

func main() {
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler.Redirect("/protected", w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
