package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A global variable that is incremented everytime a book is added.
// Used for providing a unique ID to each book
var count int

// A slice that will contain books
var books []Book

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Initializing with some dummy data
	count = 0
	books = append(books, Book{ID: count, Isbn: "123456", Title: "Book 1", Author: &Author{FirstName: "Andrew", LastName: "S"}})
	count++
	books = append(books, Book{ID: count, Isbn: "456280", Title: "Book 2", Author: &Author{FirstName: "Alex", LastName: "B"}})
	count++

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", addBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":7777", router))
}
