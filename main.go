package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

type Book struct{
	ID string `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}

var books []Book

func main(){
 router := mux.NewRouter()
 router.HandleFunc("/books",getBooks).Methods("GET")
 router.HandleFunc("/books/{id}",getBook).Methods("GET")
 router.HandleFunc("/books",addBook).Methods("POST")
 router.HandleFunc("/books",updateBook).Methods("PUT")
 router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

 log.Fatal(http.ListenAndServe(":8000",router))

}

func getBooks (w http.ResponseWriter, r *http.Request){
	log.Println("Get all books")
}

func getBook (w http.ResponseWriter, r *http.Request){
	log.Println("Get all book")
}

func addBook (w http.ResponseWriter, r *http.Request){
	log.Println("Add one book")
}

func updateBook (w http.ResponseWriter, r *http.Request){
	log.Println("Updates a book")
}

func removeBook (w http.ResponseWriter, r *http.Request){
	log.Println("Remove a book")
}