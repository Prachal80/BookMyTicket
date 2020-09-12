package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()

	router.HandleFunc("/", PingEndpoint).Methods("GET")
	router.HandleFunc("/theater", CreateTheaterEndpoint).Methods("POST")
	router.HandleFunc("/theaters", GetAllTheatersEndpoint).Methods("GET")
	router.HandleFunc("/theater/{name}", GetTheaterEndpoint).Methods("GET")
	router.HandleFunc("/theater/{name}", DeleteTheaterEndpoint).Methods("POST")
	router.HandleFunc("/theater/{name}", UpdateTheaterEndpoint).Methods("PUT")

	http.ListenAndServe(":"+port, router)

}
