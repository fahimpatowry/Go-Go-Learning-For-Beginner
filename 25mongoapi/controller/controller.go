package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	models "mongoapi/modle"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://tecnonextfahim:123@cluster0.wr4cdml.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// most important
var collection *mongo.Collection

// connect with mongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDb
	// client, err := mongo.Connect(context.TODO(), clientOption)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// connect to MongoDB (v2 style)
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	// optional but recommended
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("mono connection successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection is ready")
}

// mongodb helpers - file

// insert 1 record
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted 1 ", inserted.InsertedID)
}

// update record
func updateOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	// bson.M // if not wared about uppercase or lower case
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result.ModifiedCount)
}

// update record
func deleteOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	// bson.M // if not wared about uppercase or lower case
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result.DeletedCount)
}

// update record
func deleteAll(movieId string) int64 {

	result, err := collection.DeleteOne(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result.DeletedCount)

	return result.DeletedCount
}

// get all
func getAll() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie primitive.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

	fmt.Println("a;;data", movies)

	return movies
}

// Actual controller -file
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/x-ww-from-urlencode")

	allMovies := getAll()
	json.NewEncoder(w).Encode(allMovies)

}

// Actual controller -file
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/x-ww-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

// Actual controller -file
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/x-ww-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	param := mux.Vars(r)
	updateOne(param["id"])

	json.NewEncoder(w).Encode(param["id"])
}

// Actual controller -file
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type", "application/x-ww-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	param := mux.Vars(r)
	deleteAll(param["id"])

	json.NewEncoder(w).Encode(param["id"])
}

// Actual controller -file
// func DeleteAll(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-type", "application/x-ww-from-urlencode")
// 	w.Header().Set("Allow-Control-Allow-Methods", "POST")

// 	param := mux.Vars(r)
// 	deleteAll()

// 	json.NewEncoder(w).Encode(param["id"])
// }
