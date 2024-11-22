package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view?id=1", snippetView)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":4000"
	}

	log.Printf("Starting server on %v", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
