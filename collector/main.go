package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DadJokesWorker struct {
	client *mongo.Client
}

func NewDadJokesWorker(c *mongo.Client) *DadJokesWorker {
	return &DadJokesWorker{
		client: c,
	}
}

func (djw *DadJokesWorker) start() {
	coll := djw.client.Database("dadjokes").Collection("jokes")
	ticker := time.NewTicker(10 * time.Second)

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	client := &http.Client{}

	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		dadJoke := bson.M{}

		if err := json.NewDecoder(resp.Body).Decode(&dadJoke); err != nil {
			log.Fatal(err)
		}

		_, err = coll.InsertOne(context.TODO(), dadJoke)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n\n", dadJoke["joke"])
		<-ticker.C
	}
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
	worker := NewDadJokesWorker(client)
	worker.start()
}
