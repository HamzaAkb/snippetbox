package main

import "github.com/HamzaAkb/snippetbox/internal/models"

type templateData struct {
	Snippet 	*models.Snippet
	Snippets	[]*models.Snippet
}