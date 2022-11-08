package main

import (
	"html/template"
	"path/filepath"

	"github.com/hrshshrma/lets-go-snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("/home/harsh/projects/lets-go/snippetbox/ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		/*files := []string{
			"/home/harsh/projects/lets-go/snippetbox/ui/html/base.html",
			"/home/harsh/projects/lets-go/snippetbox/ui/html/partials/nav.html",
			page,
		}*/

		//for base
		ts, err := template.ParseFiles("/home/harsh/projects/lets-go/snippetbox/ui/html/base.html")
		if err != nil {
			return nil, err
		}

		//for partials
		ts, err = ts.ParseGlob("/home/harsh/projects/lets-go/snippetbox/ui/html/partials/nav.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
