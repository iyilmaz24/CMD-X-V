package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger	
	infoLog *log.Logger	
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":4000"
	}

	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %v", port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
