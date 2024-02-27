package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type DadJokesWorker struct {
	client *http.Client
}

func NewDadJokesWorker(c *http.Client) *DadJokesWorker {
	return &DadJokesWorker{
		client: c,
	}
}

func main() {
	client := &http.Client{}
	worker := NewDadJokesWorker(client)
	worker.start()
}

func (djw *DadJokesWorker) start() {
	// Get joke from icanhazdadjoke
	jokeReq, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	jokeReq.Header.Set("Accept", "application/json")

	jokeResp, err := djw.client.Do(jokeReq)
	if err != nil {
		log.Fatal(err)
	}

	jokeB, err := httputil.DumpResponse(jokeResp, true)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("-----------Joke Response------------")
	fmt.Println(string(jokeB))

	// Insert into local jokes database via API
	apiReq, err := http.NewRequest("POST", "http://api:8080/jokes", jokeResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	apiReq.Header.Set("Content-Type", "application/json")

	apiResp, err := djw.client.Do(apiReq)
	if err != nil {
		log.Fatal(err)
	}
	defer apiResp.Body.Close()

	// Print out the response from the API
	apiB, err := httputil.DumpResponse(apiResp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------API Response-------------")
	fmt.Println(string(apiB))
}
