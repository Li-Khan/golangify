package main

import "github.com/Li-Khan/golangify/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
