package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Book sd
type Book struct {
	ID     *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string              `json:"title"`
	Author Author              `json:"author"`
}
