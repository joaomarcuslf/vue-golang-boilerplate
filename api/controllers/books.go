package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	helpers "my_library_app/helpers"
	"my_library_app/models"
)

func ListBooks(connection *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var books []models.Book

		cur, err := connection.Collection("books").Find(context.TODO(), bson.M{})

		if err != nil {
			helpers.JSONError(err, w, 404)
			return
		}

		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {
			var book models.Book
			err := cur.Decode(&book)

			if err != nil {
				log.Fatal(err)
				helpers.JSONError(err, w, 500)
				return
			}

			books = append(books, book)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(books)
	}
}

func GetBook(connection *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var book models.Book

		var params = mux.Vars(r)

		id, _ := primitive.ObjectIDFromHex(params["id"])

		filter := bson.M{"_id": id}

		err := connection.Collection("books").FindOne(context.TODO(), filter).Decode(&book)

		if err != nil {
			fmt.Println("Error", err)
			helpers.JSONError(err, w, 404)
			return
		}

		json.NewEncoder(w).Encode(book)
	}
}

func CreateBook(connection *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var book models.Book

		_ = json.NewDecoder(r.Body).Decode(&book)

		result, err := connection.Collection("books").InsertOne(context.TODO(), book)

		if err != nil {
			helpers.JSONError(err, w, 400)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}

func UpdateBook(connection *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)

		id, _ := primitive.ObjectIDFromHex(params["id"])

		var book models.Book

		_ = json.NewDecoder(r.Body).Decode(&book)

		filter := bson.M{"_id": id}
		update := bson.M{
			"$set": bson.M{
				"title":     book.Title,
				"author":    book.Author,
				"country":   book.Country,
				"imageLink": book.ImageLink,
				"language":  book.Language,
				"link":      book.Link,
				"year":      book.Year,
				"pages":     book.Pages,
			},
		}

		_, err := connection.Collection("books").UpdateOne(context.TODO(), filter, update)

		if err != nil {
			helpers.JSONError(err, w, 400)
			return
		}

		book.ID = id

		json.NewEncoder(w).Encode(book)
	}
}

func DeleteBook(connection *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var params = mux.Vars(r)

		id, err := primitive.ObjectIDFromHex(params["id"])

		filter := bson.M{"_id": id}

		deleteResult, err := connection.Collection("books").DeleteOne(context.TODO(), filter)

		if err != nil {
			helpers.JSONError(err, w, 500)
			return
		}

		json.NewEncoder(w).Encode(deleteResult)
	}
}
