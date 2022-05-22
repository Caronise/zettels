package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/zettel/view", zettelView)
	mux.HandleFunc("/zettel/create", zettelCreate)

	log.Println("Starting server on port :8181")
	err := http.ListenAndServe(":8181", mux)
	log.Fatal(err)
}
