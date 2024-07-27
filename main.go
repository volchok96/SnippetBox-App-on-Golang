package main

import (
	"log"
	"net/http"
)

// Handler for the home page.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hi Snippetbox"))
}

// Handler for displaying the content of a note.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show a Snippet"))
}

// Handler for creating a new note.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check if the request uses the POST method
	if r.Method != http.MethodPost {
		// Use the Header().Set() method to add the 'Allow: POST' header to
		// the HTTP header map. The first parameter is the name of the header, and
		// the second parameter is the value of the header.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "GET Method forbidden!\n", 405)
		return
	}
	w.Write([]byte("Add new Snippet"))
}

func main() {
	// Register two new handlers and their corresponding URL patterns in
	// the servemux router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Running the web server on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
