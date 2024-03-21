// Package store is for general data stores
package store

// Joke is a joke from icanhazdadjoke.com
type Joke struct {
	ID   string `json:"id,omitempty"`
	Joke string `json:"joke,omitempty"`
}

// JokeStore is an interface for storing jokes
type JokeStore interface {
	GetJoke(id string) (*Joke, error)
	GetJokes() ([]*Joke, error)
}
