package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home is the handler for root, displays a 404 error.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Currently at \"/\""))
}

func zettelsView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Fprint to streamline the string conversion and write to w.
	fmt.Fprintf(w, "Display the zettel with ID %d", id)
}

func zettelsCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// set the header before using the writer.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new zettel"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/zettels/view", zettelsView)
	mux.HandleFunc("/zettels/create", zettelsCreate)

	log.Println("Starting server on port :8181")
	err := http.ListenAndServe(":8181", mux)
	log.Fatal(err)
}
