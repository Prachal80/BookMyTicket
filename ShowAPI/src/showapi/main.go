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


func main() {

    fmt.Println("Starting the application")
    ctx,_:= context.WithTimeout(context.Background(),10*time.Second)

    
    clientOptions := options.Client().ApplyURI("mongodb://admin:password@10.0.1.200:27017")
    client,_=mongo.Connect(ctx,clientOptions)
    router:=mux.NewRouter()
    router.HandleFunc("/show", CreateShowEndpoint).Methods("POST")
    router.HandleFunc("/show/{id}", DeleteShowEndpoint).Methods("DELETE")
    router.HandleFunc("/show/{id}", UpdateShowEndpoint).Methods("PUT")
    router.HandleFunc("/show/{id}", GetShowEndpoint).Methods("GET")
    router.HandleFunc("/shows", GetAllShowsEndpoint).Methods("GET")
    router.HandleFunc("/createbook/{id}", CreateBookingEndpoint).Methods("POST")
    router.HandleFunc("/book/{id}", BookShowEndpoint).Methods("POST")
    router.HandleFunc("/book/{id}", GetBookingEndpoint).Methods("GET")
    router.HandleFunc("/bookings", GetAllBookingEndpoint).Methods("GET")
    http.ListenAndServe(":3000",router)
}