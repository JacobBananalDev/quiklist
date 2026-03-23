package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jacobbananaldev/quiklist/internal/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Routes
	r.Get("/health", handlers.HealthHandler)

	return r
}