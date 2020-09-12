package main

//import "go.mongodb.org/mongo-driver/bson/primitive"

type Show struct {
    ShowID string `json:"ShowID,omitempty"    bson:"ShowID,omitempty"`
    TheatreID  string `json:"TheatreID,omitempty"    bson:"TheatreID,omitempty"`
    MovieID string `json:"MovieID,omitempty"    bson:"MovieID,omitempty"`
   
}

type Bookings struct{
    ShowID string `json:"ShowID,omitempty"    bson:"ShowID,omitempty"`
    Users []string `json:"Users,omitempty'     bson:Users,omitempty` 
}

type InCommingRequest struct{
    ShowID string `json:"ShowID,omitempty"    bson:"ShowID,omitempty"`
    User string `json:"User,omitempty'     bson:User,omitempty` 
}