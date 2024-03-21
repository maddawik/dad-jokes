package handlers

import (
	"net/http"

	"github.com/mawdac/go-docker-api-test/internal/store"
	"github.com/mawdac/go-docker-api-test/internal/store/dbstore"
	"github.com/mawdac/go-docker-api-test/internal/templates"
)

type HomeHandler struct {
	jokeStore store.JokeStore
}

func NewHomeHandler(j *dbstore.JokeStore) *HomeHandler {
	return &HomeHandler{
		jokeStore: j,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jokes, err := h.jokeStore.GetJokes()
	if err != nil {
		http.Error(w, "Error getting jokes", http.StatusInternalServerError)
		return
	}

	c := templates.Index(jokes)
	err = templates.Base(c, "My website").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
