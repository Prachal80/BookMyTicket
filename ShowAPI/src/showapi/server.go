package main

import(
    
    "context"
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "time"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    //"go.mongodb.org/mongo-driver/mongo/options"
    //"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"


)
var client *mongo.Client

func CreateShowEndpoint(response http.ResponseWriter,request *http.Request){
    response.Header().Set("content-type", "application/json")
    var show Show
    _ = json.NewDecoder(request.Body).Decode(&show)
    collection := client.Database("Show").Collection("Shows")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    result, _ := collection.InsertOne(ctx, show)
    if result ==nil{
       json.NewEncoder(response).Encode("{ message:Show already exists!}")
    }else{
    json.NewEncoder(response).Encode(result)
    }
}


func GetShowEndpoint(response http.ResponseWriter,request *http.Request){
   response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id:= (params["id"])
    var show Show
    collection := client.Database("Show").Collection("Shows")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, Show{ShowID: id}).Decode(&show)
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(show)
}


func GetAllShowsEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var shows []Show
    collection := client.Database("Show").Collection("Shows")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var show Show
        cursor.Decode(&show)
        shows = append(shows, show)
    }
    if err := cursor.Err(); err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(shows)
}


func DeleteShowEndpoint(response http.ResponseWriter, request *http.Request){
    response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id, _ := (params["id"])
    //var movie Movie
    filter := bson.D{{"ShowID",id}}
    collection := client.Database("Show").Collection("Shows")
    ///////
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    deleteResult, err := collection.DeleteOne(ctx, filter)
    if err != nil {
    log.Fatal(err)
    }
    fmt.Printf("Deleted %v documents in the Movies collection\n", deleteResult.DeletedCount)

    json.NewEncoder(response).Encode(deleteResult)
}


func UpdateShowEndpoint(response http.ResponseWriter, request *http.Request){
    response.Header().Set("content-type", "application/json")

    var show Show
    params := mux.Vars(request)
    id := params["id"]
    
    _ = json.NewDecoder(request.Body).Decode(&show)

    fmt.Printf("Show : %q" , show.ShowID )
    fmt.Printf("Theatre : %q" , show.TheatreID )
    fmt.Printf("Movie : %q" , show.MovieID)

    filter := bson.D{{"ShowID",id}}

    update := bson.D{
    { "$set",  bson.D{
        {"ShowID", show.ShowID } ,{"TheatreID", show.TheatreID } ,{"MovieID", show.MovieID },
    }},
    }

    collection := client.Database("Show").Collection("Shows")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    ////
    updateResult, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
    log.Fatal(err)
    }

    fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
    

    json.NewEncoder(response).Encode(updateResult)
}



func CreateBookingEndpoint(response http.ResponseWriter,request *http.Request){
    response.Header().Set("content-type", "application/json")
    var booking Bookings
    _ = json.NewDecoder(request.Body).Decode(&booking)
    collection := client.Database("Show").Collection("Bookings")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    result, _ := collection.InsertOne(ctx, booking)

   
    
    if result ==nil{
       json.NewEncoder(response).Encode("{ message:Show already exists!}")
    }else{
    json.NewEncoder(response).Encode(result)
    }
}

func BookShowEndpoint(response http.ResponseWriter,request *http.Request){
    response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id:= (params["id"])
    var bookings Bookings
    var currentBooking InCommingRequest
    _=json.NewDecoder(request.Body).Decode(&currentBooking)
    collection:= client.Database("Show").Collection("Bookings")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, Bookings{ShowID: id}).Decode(&bookings)
    if err != nil {
        fmt.Println("Error while fetching stored show bookings")
        return
    }
    fmt.Println("Booking fetched!")
    addedUsers:=append(bookings.Users,currentBooking.User)
    fmt.Println(addedUsers)
    filter := bson.D{{"ShowID",id}}
    update := bson.D{
    {"$set", bson.D{
        {"Users", addedUsers},
    }},
    }
    fmt.Println("BSON updated!")
    //bookings.Users=addedUsers
    updatedBooking, err1 := collection.UpdateOne(ctx, filter, update)
    if err1 != nil {
    log.Fatal(err1)
    }
    json.NewEncoder(response).Encode(updatedBooking)
}



func GetBookingEndpoint(response http.ResponseWriter,request *http.Request){
   response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id:= (params["id"])
    var bookings Bookings
    collection := client.Database("Show").Collection("Bookings")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, Bookings{ShowID: id}).Decode(&bookings)
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(bookings)
}


func GetAllBookingEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var bookings []Bookings
    collection := client.Database("Show").Collection("Bookings")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var booking Bookings
        cursor.Decode(&booking)
        bookings = append(bookings, booking)
    }
    if err := cursor.Err(); err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(bookings)
}
