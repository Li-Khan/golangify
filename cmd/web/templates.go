package main

import (
	"path/filepath"
	"text/template"

	"github.com/Li-Khan/golangify/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}
	return cache, nil
}
