package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iyilmaz24/CMD-X-V.git/internal/models"
)

type application struct {
	errorLog *log.Logger	
	infoLog *log.Logger	
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":4000"
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		errorLog.Fatal("DSN not set in environment variables")
	}

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache();
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %v", port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
