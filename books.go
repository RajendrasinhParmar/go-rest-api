package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Give all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(books)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Jay Mahadev </h1>")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)

	for _, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = count
	count++

	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Update existing books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)

	var tempBook Book

	for index, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			_ = json.NewDecoder(r.Body).Decode(&tempBook)
			tempBook.ID = index
			books[index] = tempBook
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Delete the book with ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)

	var tempBook Book
	for index, book := range books {
		if strconv.Itoa(book.ID) == vars["id"] {
			tempBook = books[index]
			books = append(books[:index], books[index+1:]...)
			_ = json.NewEncoder(w).Encode(tempBook)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}
