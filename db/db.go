package db

import (
	"context"
	"fmt"
	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"time"
)

var db *mongo.Database

func init() {

	coptions := options.Client().ApplyURI("mongodb://mongo2")
	client, err := mongo.NewClient(coptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("restgo")
	fmt.Println(db)
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("connecting to MongoDB!")

}

//Start is return the collection name
func Start(book string) *mongo.Collection {
	fmt.Println(db)
	collection := db.Collection(book)
	return collection

}
