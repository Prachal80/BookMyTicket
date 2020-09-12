package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Main Function
func main() {

	//mongoURI := "mongodb://cmpe281:cmpe281@10.0.1.244:27017"

	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/", PingCheckEndpoint).Methods("GET")
	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user", GetAllUsersEndpoint).Methods("GET")
	router.HandleFunc("/user/{email}", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/user/{email}", DeleteUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{email}", UpdateUserEndpoint).Methods("PUT")

	http.ListenAndServe(":3000", router)
	//"mongodb://localhost:27017"
}
