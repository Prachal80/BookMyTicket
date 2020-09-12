package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoURI = "mongodb://10.0.1.183:27017"
// var mongoURI = "mongodb://localhost:27017"
var client *mongo.Client

// PingEndpoint to check pinging.
func PingEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response).Encode("{'Response' : 'Ping successful'}")
}

// GetTheaterEndpoint to get theater by name.
func GetTheaterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["name"]
	var theater Theater
	collection := client.Database("Theater").Collection("Theaters")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Theater{Name: name}).Decode(&theater)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(theater)
}

// CreateTheaterEndpoint to register theater.
func CreateTheaterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var theater Theater
	_ = json.NewDecoder(request.Body).Decode(&theater)

	collection := client.Database("Theater").Collection("Theaters")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, theater)
	if result == nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode("{'Response' : 'Registration Successful'}")
}

// DeleteTheaterEndpoint to delete theater.
func DeleteTheaterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	name, _ := params["name"]
	filter := bson.D{{"name", name}}
	var theater Theater
	collection := client.Database("Theater").Collection("Theaters")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Theater{Name: name}).Decode(&theater)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if deleteResult == nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode("{'Response' : 'Deleted Successfully'}")
}

// UpdateTheaterEndpoint to udate details of theater.
func UpdateTheaterEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var theater Theater
	params := mux.Vars(request)
	name, _ := params["name"]
	_ = json.NewDecoder(request.Body).Decode(&theater)
	filter := bson.D{{"name", name}}

	update := bson.D{
		{"$set", bson.D{
			{"name", theater.Name},
			{"rating", theater.Rating},
			{"address", theater.Address},
			{"screens", theater.Screens},
		}},
	}

	collection := client.Database("Theater").Collection("Theaters")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Theater{Name: name}).Decode(&theater)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	errGet := collection.FindOne(ctx, Theater{Name: name}).Decode(&theater)
	if errGet != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "Does not exists" }`))
		return
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if updateResult == nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode("{'Response' : 'Updated Successfully'}")
}

// GetAllTheatersEndpoint to get all theaters registered to this platform.
func GetAllTheatersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var theaters []Theater
	collection := client.Database("Theater").Collection("Theaters")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var theater Theater
		cursor.Decode(&theater)
		theaters = append(theaters, theater)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Response": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(theaters)
}
