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

	router.HandleFunc("/movie", CreateMovieEndpoint).Methods("POST")
	router.HandleFunc("/movies", GetAllMoviesEndpoint).Methods("GET")
	router.HandleFunc("/movie/{name}", GetMovieEndpoint).Methods("GET")
	router.HandleFunc("/movie/{name}", UpdateMovieEndpoint).Methods("PUT")
	router.HandleFunc("/movie/{name}", DeleteMovieEndpoint).Methods("POST")
	router.HandleFunc("/", PingCheckEndpoint).Methods("GET")
	
	http.ListenAndServe(":"+port, router)

	}