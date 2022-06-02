package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":8181", "Port to run the server on.")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", fileServer)

	mux.HandleFunc("/", home)
	mux.HandleFunc("/zettel/view", zettelView)
	mux.HandleFunc("/zettel/create", zettelCreate)

	log.Printf("Starting server on %s", *port)
	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}
