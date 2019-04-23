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

func init() {

	coptions := options.Client().ApplyURI("mongodb://foo:bar@localhost:27017")
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
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("connecting to MongoDB!")

}

//Start is db start print
func Start() string {
	return "connected"
}
