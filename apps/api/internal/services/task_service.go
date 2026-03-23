package services

import (
	"github.com/jacobbananaldev/quiklist/internal/models"
	"github.com/jacobbananaldev/quiklist/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() []models.Task {
	return s.repo.GetAll()
}