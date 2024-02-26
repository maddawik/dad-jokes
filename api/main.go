package main

import (
	"context"
	"encoding/json"
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
	http.ListenAndServe(":6969", nil)
}
