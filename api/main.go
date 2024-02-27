package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

func NewServer(c *mongo.Client) *Server {
	return &Server{
		client: c,
	}
}

func (s *Server) handleGetLatestJoke(w http.ResponseWriter, r *http.Request) {
	// var id primitive.ObjectID
	coll := s.client.Database("dadjokes").Collection("jokes")
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{}, opts).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println(err)
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", result)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

	log.Printf("Receieved request for latest joke from %v\n", r.RemoteAddr)
}

func (s *Server) handleGetAllJokes(w http.ResponseWriter, r *http.Request) {
	coll := s.client.Database("dadjokes").Collection("jokes")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)

	log.Printf("Receieved request for all jokes from %v\n", r.RemoteAddr)
}

func main() {
	mongodb_uri := os.Getenv("MONGODB_URI")
	clientOpts := options.Client().ApplyURI(mongodb_uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	server := NewServer(client)
	http.HandleFunc("/jokes", server.handleGetAllJokes)
	http.HandleFunc("/joke", server.handleGetLatestJoke)
	http.ListenAndServe(":6969", nil)
}
