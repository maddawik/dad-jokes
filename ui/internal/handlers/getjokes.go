package handlers

import (
	"net/http"

	"github.com/mawdac/go-docker-api-test/internal/store"
	"github.com/mawdac/go-docker-api-test/internal/store/dbstore"
	"github.com/mawdac/go-docker-api-test/internal/templates"
)

type GetJokesHandler struct {
	jokeStore store.JokeStore
}

func NewGetJokesHandler(j *dbstore.JokeStore) *GetJokesHandler {
	return &GetJokesHandler{
		jokeStore: j,
	}
}

func (h *GetJokesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jokes, err := h.jokeStore.GetJokes()
	if err != nil {
		http.Error(w, "Error getting jokes", http.StatusInternalServerError)
		return
	}

	c := templates.Jokes(jokes)
	err = c.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
