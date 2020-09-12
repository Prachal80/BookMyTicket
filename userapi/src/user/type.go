package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//Types
type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"   bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}
