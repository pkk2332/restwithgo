package main

import (
	"restwithgo/controller/book"

	"github.com/gin-gonic/gin"
)

//Book is the representation of

func main() {

	// r := mux.NewRouter()

	// r.HandleFunc("/api/books", book.GetBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", book.GetBook).Methods("GET")
	// r.HandleFunc("/api/books", book.CreateBook).Methods("POST")
	// r.HandleFunc("/api/books/", book.UpdateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", book.DeleteBook).Methods("DELETE")
	// log.Fatal(http.ListenAndServe(":8001", r))
	r := gin.Default()
	r.GET("books", book.GinGetBooks)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "asdasdasasd",
		})
	})
	r.Run()
}
