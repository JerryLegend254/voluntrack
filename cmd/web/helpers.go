package main

import (
	"bytes"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Get template from cache using name
	ts, ok := app.templateCache[name]

	// If template not found in cache, log error and return
	if !ok {
		app.errorLog.Fatal("Could not get template from cache")
		return
	}

	// Create a buffer to write the template to
	buf := new(bytes.Buffer)

	// Execute the template with the template data
	err := ts.Execute(buf, td)

	// If there was an error executing the template, log error and return
	if err != nil {
		app.errorLog.Fatal("Could not execute template")
		return
	}

	// Write the buffer to the response writer
	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Println(err)

	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, "Bad request", status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
