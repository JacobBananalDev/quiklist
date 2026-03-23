package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jacobbananaldev/quiklist/internal/handlers"
	"github.com/jacobbananaldev/quiklist/internal/repository"
	"github.com/jacobbananaldev/quiklist/internal/services"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Routes
	r.Get("/health", handlers.HealthHandler)

	// Dependencies wiring
	taskRepo := repository.NewTaskRepository()
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r.Get("/task", taskHandler.GetTasks)

	return r
}