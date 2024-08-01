package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Handler for the home page.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Handler for displaying the content of a note.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the parameter id from the URL and try
	// to convert the string to an integer using the strconv.Atoi() function. If it cannot
	// be converted to an integer, or if the value is less than 1, return a response
	// 404 - page not found!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the fmt.Fprintf() function to insert the value from id into the response string
	// and write it to http.ResponseWriter.
	fmt.Fprintf(w, "Displaying the selected note with ID %d...\n", id)
}

// Handler for creating a new note
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check if the request uses the POST method
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add the 'Allow: POST' header to
		// the HTTP header map. The first parameter is the name of the header, and
		// the second parameter is the value of the header.
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // 405 Error
		return
	}

	// test variables
	title := "Convenient Dependency Management"
	content := `Go has a built-in dependency management system using 
	the go mod tool. This makes it easy to manage versions of libraries 
	and packages used in a project. The system also simplifies the build 
	and deployment processes for applications.`
	expires := "7"

	// pass the data to the SnippetModel.Insert() method
	// get back ID of the newly created record to the DB
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
