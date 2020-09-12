package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Theater represents theater structure with fields defined below.
type Theater struct {
	ID      primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"   bson:"name,omitempty"`
	Rating  string             `json:"rating,omitempty" bson:"rating,omitempty"`
	Address string             `json:"address,omitempty"    bson:"address,omitempty"`
	Screens string             `json:"screens,omitempty"    bson:"screens,omitempty"`
}

// Theaters is an array of type Theater
type Theaters []Theater
