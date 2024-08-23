package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {

	tc, err := newTemplateCache("./ui/html/")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{templateCache: tc}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Server started on port 8080")
	log.Fatal(srv.ListenAndServe())

}
