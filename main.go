package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv := &http.Server{
		Addr:    ":" + "8000",
		Handler: mux,
	}
	log.Println("serving on port 8000")
	log.Fatal(srv.ListenAndServe())

}
