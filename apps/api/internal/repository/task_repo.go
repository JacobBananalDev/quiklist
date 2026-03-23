package repository

import (
	"time"

	"github.com/jacobbananaldev/quiklist/internal/models"
)

type TaskRepository struct {
	tasks []models.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: []models.Task{
			{
				ID:        1,
				Title:     "Learn Go",
				Completed: false,
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "Build QuikList",
				Completed: false,
				CreatedAt: time.Now(),
			},
		},
	}
}

func (r *TaskRepository) GetAll() []models.Task {
	return r.tasks
}