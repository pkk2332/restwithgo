package book

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
	"github.com/pkkcode/restgo/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkkcode/restgo/model"
	// . "github.com/pkkcode/restgo/model"
)

var books []model.Book
var b *mongo.Collection

// var b11 bson.TypeObjectID

func init() {

	b = db.Start("books")
	fmt.Println(b)
	//	fmt.Println(b11)

}

//GetBooks is
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	cur, err := b.Find(context.Background(), bson.D{})
	var results []model.Book
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := model.Book{}
		err := cur.Decode(&result)
		results = append(results, result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the raw bson bytes use cursor.Current
		//raw := cur.Current
		// do something with raw...
	}
	defer cur.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(results)
}

//GetBook is
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	rid, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rid)
	cur := b.FindOne(context.Background(), bson.M{"_id": rid})

	var b3 model.Book
	err1 := cur.Decode(&b3)
	if err1 != nil {
		log.Fatal(err1)
	}
	//	mongo.SingleResult
	fmt.Println(b3)
	json.NewEncoder(w).Encode(b3)
}

//CreateBook is
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	res, err := b.InsertOne(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}

//UpdateBook is
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)
	update := bson.D{
		{"$set", book},
	}
	// fmt.Println(book)
	res, err := b.UpdateOne(context.Background(), bson.D{{"_id", book.ID}}, update)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	//json.NewEncoder(w).Encode(res)
}

//DeleteBook is
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// to change the string id into binary json id in mongodb
	rid, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(params["id"])
	res, err := b.DeleteOne(context.Background(), bson.M{"_id": rid})
	if err != nil {
		log.Fatal(err)
	}
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		break
	// 	}
	// }
	fmt.Println(res.DeletedCount)
	json.NewEncoder(w).Encode(true)
}
