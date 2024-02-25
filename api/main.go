package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	worker := NewDadJokesWorker(client)
	go worker.start()
	// if err = worker.start(); err != nil {
	// 	log.Fatal(err)
	// }
	server := NewServer(client)
	http.HandleFunc("/jokes", server.handleGetAllJokes)
	http.ListenAndServe(":6969", nil)
}
