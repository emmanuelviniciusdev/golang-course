package controllers

import (
	"25mongoapi/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "netflix-movies"
const collectionName = "watchlist"

var collection *mongo.Collection

func Init()  {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(databaseName).Collection(collectionName)

	fmt.Println("MongoDB connection was successful")
}

func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie models.MovieNetflixInput

	if r.Body == nil {
		fmt.Println("Body is null")
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewDecoder(r.Body).Decode(&newMovie)

	result := insertOneMovie(newMovie)

	json.NewEncoder(w).Encode(result.InsertedID)
}

func UpdateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedMovie models.MovieNetflixInput

	if r.Body == nil {
		fmt.Println("Body is null")
		json.NewEncoder(w).Encode(nil)
		return
	}

	json.NewDecoder(r.Body).Decode(&updatedMovie)

	params := mux.Vars(r)

	updateOneMovie(params["movieId"], updatedMovie.MovieName, updatedMovie.Watched)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	deleteOneMovie(params["movieId"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deleteAllMovies()
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func getAllMovies() []bson.M {
	ctx := context.TODO()

	filterQuery := bson.M{}

	cursor, err := collection.Find(ctx, filterQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var movies []bson.M

	for cursor.Next(ctx) {
		var movie bson.M

		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	fmt.Println("Quantity of movies being returned:", len(movies))

	return movies
}

func deleteAllMovies() *mongo.DeleteResult {
	filterQuery := bson.M{}

	result, err := collection.DeleteMany(context.TODO(), filterQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)

	return result
}

func deleteOneMovie(movieId string) *mongo.DeleteResult {
	movieObjectId, _ := primitive.ObjectIDFromHex(movieId)

	filterQuery := bson.M{"_id": movieObjectId}

	result, err := collection.DeleteOne(context.TODO(), filterQuery)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Deleted count:", result.DeletedCount)

	return result
}

func updateOneMovie(movieId string, movieName string, watched bool) {
	movieObjectId, _ := primitive.ObjectIDFromHex(movieId)

	filterQuery := bson.M{"_id": movieObjectId}
	updateQuery := bson.M{"$set": bson.M{"movieName": movieName, "watched": watched}}

	result, err := collection.UpdateOne(context.TODO(), filterQuery, updateQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

func insertOneMovie(movie models.MovieNetflixInput) *mongo.InsertOneResult {
	result, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie inserted successfully. ID:", result.InsertedID)

	return result
}
