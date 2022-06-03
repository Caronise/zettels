package main

import "net/http"

// the routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", fileServer)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/zettel/view", app.zettelView)
	mux.HandleFunc("/zettel/create", app.zettelCreate)

	return mux
}
