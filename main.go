package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkkcode/restgo/controller/book"
)

//Book is the representation of

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/books", book.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", book.GetBook).Methods("GET")
	r.HandleFunc("/api/books", book.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/", book.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", book.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", r))

}
