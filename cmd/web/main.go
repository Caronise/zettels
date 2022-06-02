package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", ":8181", "Port to run the server on.")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", fileServer)

	mux.HandleFunc("/", home)
	mux.HandleFunc("/zettel/view", zettelView)
	mux.HandleFunc("/zettel/create", zettelCreate)

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
