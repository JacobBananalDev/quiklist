package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jacobbananaldev/quiklist/internal/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := h.service.GetTasks()

	json.NewEncoder(w).Encode(tasks)
}