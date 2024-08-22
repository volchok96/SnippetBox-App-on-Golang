package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper logs the error message to errorLog and
// then sends a 500 "Internal Server Error" response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	if outputErr := app.errorLog.Output(2, trace); outputErr != nil {
		fmt.Printf("error logging to errorLog: %v\n", outputErr)
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code
// and corresponding description to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Implementation of the notFound helper.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, name string, tmpldata *templateData) {
	tmpls, ok := app.tmplCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("template %s does not exist", name))
		return
	}

	err := tmpls.Execute(w, tmpldata)
	if err != nil {
		app.serverError(w, err)
	}
}