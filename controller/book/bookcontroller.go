package book

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/pkkcode/restgo/db"

	"github.com/gorilla/mux"
	"github.com/pkkcode/restgo/model"
	// . "github.com/pkkcode/restgo/model"
)

var books []model.Book

func init() {
	aa := db.Start()
	fmt.Println(aa)
}

//GetBooks is
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//GetBook is
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Book{})
}

//CreateBook is
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// for _, b := range book {
	book.ID = strconv.Itoa(rand.Intn(1000000))
	// }
	fmt.Println(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

//UpdateBook is
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			CreateBook(w, r)
			break
		}
	}
	//	json.NewEncoder(w).Encode(books)
}

//DeleteBook is
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}
