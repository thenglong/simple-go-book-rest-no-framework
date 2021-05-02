package main

import (
	"boot-rest-api/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var bh handlers.BookHandler

	http.Handle("/books", &bh)

	// this handles the /book/id
	http.Handle("/books/", &bh)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
