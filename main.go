package main

import (
	"boot-rest-api/handlers"
	"log"
	"net/http"
)

func main() {
	var bh handlers.BookHandler

	http.Handle("/books", &bh)

	// this handles the /book/id
	http.Handle("/books/", &bh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
