package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/unrolled/render"
)

func main() {

	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	// 	port = "12345"
	// }

	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://cmpe281:cmpe281@3.89.47.220:27017")
	clientOptions := options.Client().ApplyURI(mongodb_server)
	fmt.Println("Client Options set...")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Mongo Connected...")
	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person/{id}", RemovePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/sendmail/{id}", sendmail).Methods("POST")
	router.HandleFunc("/sendsms/{id},{no}", sendsms).Methods("POST")
	router.HandleFunc("/", ping).Methods("GET")
	http.ListenAndServe(":3000", router)

}
