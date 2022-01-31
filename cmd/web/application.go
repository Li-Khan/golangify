package main

import (
	"log"

	"github.com/Li-Khan/golangify/pkg/models/sqlite"
)

// Application ...
type Application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	snippets *sqlite.SnippetModel
}
