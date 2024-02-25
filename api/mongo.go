package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DadJokesWorker struct {
	client *mongo.Client
}

func NewDadJokesWorker(c *mongo.Client) *DadJokesWorker {
	return &DadJokesWorker{
		client: c,
	}
}

func (djw *DadJokesWorker) start() error {
	coll := djw.client.Database("dadjokes").Collection("jokes")
	ticker := time.NewTicker(10 * time.Second)

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	client := &http.Client{}

	for {
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		dadJoke := bson.M{}

		if err := json.NewDecoder(resp.Body).Decode(&dadJoke); err != nil {
			return err
		}

		_, err = coll.InsertOne(context.TODO(), dadJoke)
		if err != nil {
			return err
		}

		fmt.Println(dadJoke)
		<-ticker.C
	}
}
