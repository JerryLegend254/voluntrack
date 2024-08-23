package main

import (
	"bytes"
	"log"
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	// Get template from cache using name
	ts, ok := app.templateCache[name]

	// If template not found in cache, log error and return
	if !ok {
		log.Fatal("Could not get template from cache")
		return
	}

	// Create a buffer to write the template to
	buf := new(bytes.Buffer)

	// Execute the template with the template data
	err := ts.Execute(buf, td)

	// If there was an error executing the template, log error and return
	if err != nil {
		log.Fatal("Could not execute template")
		return
	}

	// Write the buffer to the response writer
	buf.WriteTo(w)
}
