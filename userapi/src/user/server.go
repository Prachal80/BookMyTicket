package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var mongodb_database = "Users"
var mongodb_collection = "people"

//var mongoURI = "mongodb://localhost:27017"

//var mongoURI = "mongodb://host.docker.internal:27017"
var mongoURI = "mongodb://cmpe281:cmpe281@10.0.1.244:27017"

type Users []User

//Ping Check
func PingCheckEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	json.NewEncoder(response).Encode("{'Response':'API is UP!'}")

}

// Get User by Email
func GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	email, _ := params["email"]
	var user User
	collection := client.Database(mongodb_database).Collection(mongodb_collection)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{Email: email}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(user)
}

//Create a User
func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database(mongodb_database).Collection(mongodb_collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(result)

}

//Delete a User
func DeleteUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	params := mux.Vars(request)
	email, _ := params["email"]
	//var movie Movie
	filter := bson.D{{"email", email}}
	collection := client.Database(mongodb_database).Collection(mongodb_collection)
	///////
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, User{Email: email}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	deleteResult, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the Movies collection\n", deleteResult.DeletedCount)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(deleteResult)
}

//Update User Information
func UpdateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var user User
	params := mux.Vars(request)
	email, _ := params["email"]

	_ = json.NewDecoder(request.Body).Decode(&user)

	fmt.Println("newName : %q", user.Name)
	fmt.Println("newEmail : %q", user.Email)
	fmt.Println("newPassword : %q", user.Password)

	filter := bson.D{{"email", email}}

	update := bson.D{
		{"$set", bson.D{
			{"name", user.Name}, {"email", user.Email}, {"password", user.Password},
		}},
	}

	collection := client.Database(mongodb_database).Collection(mongodb_collection)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	////
	err := collection.FindOne(ctx, User{Email: email}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	json.NewEncoder(response).Encode(updateResult)
}

// Get All Users from Database
func GetAllUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []User
	collection := client.Database(mongodb_database).Collection(mongodb_collection)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}
