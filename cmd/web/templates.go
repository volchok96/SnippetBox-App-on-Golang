package main

import "volchok96.com/snippetbox/pkg/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}