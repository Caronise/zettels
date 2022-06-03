package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	port := flag.String("port", ":8181", "Port to run the server on.")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog,
		infoLog,
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(), // Call app.routes() method to get the servemux
	}

	infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
