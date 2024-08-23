package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello VolunTrack!"))
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started on port 8080")
	log.Fatal(srv.ListenAndServe())

}
