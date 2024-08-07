package main

import (
	"html/template"
	"path/filepath"

	"volchok96.com/snippetbox/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map that will store the cache.
	appCache := map[string]*template.Template{}

	// Use the filepath.Glob function to get a slice of all file paths with
    // the '.page.tmpl' extension. Essentially, we will get a list of all template files for the pages
    // of our web application.
	pageTmpls, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, pageTmpl := range pageTmpls {
		// Extract the final file name (e.g., 'home.page.tmpl') from the full file path
        // and assign it to the variable name.
		name := filepath.Base(pageTmpl)

		tmpls, err := template.ParseFiles(pageTmpl)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlob method to add all layout templates.
		tmpls, err = tmpls.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		tmpls, err = tmpls.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add the resulting set of templates to the cache, using the page name
        // (e.g., home.page.tmpl) as the key for our map.
		appCache[name] = tmpls
	}

	return appCache, nil
}