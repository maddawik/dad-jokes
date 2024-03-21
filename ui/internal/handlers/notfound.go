package handlers

import (
	"net/http"

	"github.com/mawdac/go-docker-api-test/internal/templates"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.NotFound()
	w.WriteHeader(http.StatusNotFound)
	err := templates.Base(c, "Not Found").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
