package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID      primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Theater string             `json:"theater,omitempty" bson:"theater,omitempty"`
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}
