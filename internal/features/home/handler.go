package home

import (
	"go2start/internal/templates/pages"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		// Home : GET /
		Home(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service string
	}
)

func NewHandler(svc string) Handler {
	return &handler{service: svc}
}

func Mount(r chi.Router, h Handler) {
	r.Get("/", h.Home)
}

func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	if err := pages.HomePage().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
