package main

import (
	"log"
	"text/template"

	"github.com/Li-Khan/golangify/pkg/models/sqlite"
)

// Application ...
type Application struct {
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	snippets      *sqlite.SnippetModel
	templateCache map[string]*template.Template
}
