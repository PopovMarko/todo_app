package tasks_service

import (
	"context"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

type TasksService struct {
	tasksRepository TasksRepository
}

func NewTaskService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
