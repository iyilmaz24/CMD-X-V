package main

import (
	"path/filepath"
	"text/template"

	"github.com/iyilmaz24/CMD-X-V.git/internal/models"
)


type templateData struct { 
	Snippet *models.Snippet 
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) { 
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl") // get a slice of all filepaths with the .tmpl extension in the // ui/html/pages directory.
	if err != nil {
		return nil, err 
	}

	for _, page := range pages {
		name := filepath.Base(page) // extract the base name from the page filepath.

		files := []string{
		"./ui/html/base.tmpl", "./ui/html/partials/nav.tmpl", page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err 
		}

		cache[name] = ts // add the template set to the cache, using the base name as the key.
	}

	return cache, nil 
}