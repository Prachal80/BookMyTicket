package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"time"

	"net/http"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
	"github.com/jordan-wright/email"
	"github.com/subosito/twilio"
)

var client *mongo.Client

var mongodb_server = "mongodb://admin:password@10.0.1.247:27017"

//var mongodb_server = "mongodb://localhost:27017"
var mongodb_server_1 = "52.37.128.85:27017"
var mongodb_database = "movies"
var mongodb_collection = "people"
var mongodb_collection1 = "submissions"
var mongodb_username = "jay"
var mongodb_password = "jay"

// type Person struct {
// 	UserId string `json:"UserId,omitempty" bson:"UserId,omitempty"`
// 	Name   string `json:"name,omitempty" bson:"name,omitempty"`
// 	Email  string `json:"email,omitempty" bson:"email,omitempty"`
// 	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
// }

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
func ping(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response).Encode("{message: Ping Check Passed}")
}

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	fmt.Println("Name", person.Name)
	collection := client.Database(mongodb_database).Collection("people")
	fmt.Println("Name", person.Name)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	var person Person
	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{Email: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func RemovePersonEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	var person Person

	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := Person{Email: id}
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println(result.DeletedCount)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	var people []Person
	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func sendmail(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Sending........")

	var person Person

	params := mux.Vars(request)
	id := (params["id"])

	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{Email: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	to2 := id

	e := email.NewEmail()
	e.From = "Confirmation@sandbox1251b125b80941c1a31f4c8a505f8afc.mailgun.org"
	e.To = []string{to2}
	e.Subject = "Ticket Confirmation"
	e.Text = []byte(`Yayy Ticket Booked..!!

	Your Ticket for movie "` + person.Movie + `" is booked at the theater  "` + person.Theater + `"	 ENJOY... :)`)
	err = e.Send("smtp.mailgun.org:587", smtp.PlainAuth("", "postmaster@sandbox1251b125b80941c1a31f4c8a505f8afc.mailgun.org", "f62247ebb7a9f704dbb58a4721546c7c-09001d55-d50d34a8", "smtp.mailgun.org"))
	if err != nil {
		json.NewEncoder(response).Encode("{message : Email not sent}")
	} else {
		json.NewEncoder(response).Encode("{message : Email sent...}")
	}
}

func sendsms(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Sending sms")
	AccountSid := "AC6e3ee3aa9f0e15dc727f715cf9d05838"
	AuthToken := "01ae941c8c7f80e1a1f07a7a874d6d03"
	From := "+12015146384"

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	no := (params["no"])
	fmt.Println(id)
	fmt.Println(no)
	var person Person

	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//c := session.DB(mongodb_database).C(mongodb_collection)
	response.Header().Set("content-type", "application/json")
	fmt.Println(id)
	err := collection.FindOne(ctx, Person{Email: id}).Decode(&person)

	response.Header().Set("content-type", "application/json")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	to2 := no
	To := to2
	// Initialize twilio Client
	c2 := twilio.NewClient(AccountSid, AuthToken, nil)

	// Send Message
	params2 := twilio.MessageParams{
		Body: "Your Ticket for movie \"" + person.Movie + "\"  is booked at the theater \"" + person.Theater + "\"	ENJOY...",
	}

	s, resp, err := c2.Messages.Send(From, To, params2)
	log.Println("Send:", s)
	log.Println("Response:", resp)
	log.Println("Err:", err)
	json.NewEncoder(response).Encode("{'message' : 'Message sent'}")

}
