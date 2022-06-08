package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "Golang pointers1", Author: "Mr. Kasun1", Year: "2010"},
		Book{ID: 2, Title: "Golang pointers2", Author: "Mr. Kasun2", Year: "2011"},
		Book{ID: 3, Title: "Golang pointers3", Author: "Mr. Kasun3", Year: "2012"},
		Book{ID: 4, Title: "Golang pointers4", Author: "Mr. Kasun4", Year: "2013"},
		Book{ID: 5, Title: "Golang pointers5", Author: "Mr. Kasun5", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all book")
	params := mux.Vars(r)
	i, _ :=strconv.Atoi(params["id"])
    
	for _, book :=range books{
		if book.ID == i{
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Add one book")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books,book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	//log.Println("Updates a book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i ,item := range books {
		if item.ID == book.ID{
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book")
}
