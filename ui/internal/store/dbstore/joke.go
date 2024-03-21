package dbstore

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mawdac/go-docker-api-test/internal/store"
)

type JokeStore struct {
	jokes []store.Joke
}

func NewJokeStore() *JokeStore {
	return &JokeStore{
		jokes: []store.Joke{},
	}
}

func (s *JokeStore) GetJoke(id string) (*store.Joke, error) {
	url := "http://localhost:42069/joke/" + id
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	var joke *store.Joke

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return nil, err
	}

	return joke, nil
}

func (s *JokeStore) GetJokes() ([]*store.Joke, error) {
	url := "http://localhost:42069/jokes"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}
	//
	jokes := make([]*store.Joke, 0)
	//
	err = json.Unmarshal(body, &jokes)
	if err != nil {
		return nil, err
	}

	// jokes = append(jokes, &store.Joke{ID: "123", Joke: "abc"})

	return jokes, nil
}
