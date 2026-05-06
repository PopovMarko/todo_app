package tasks_service

import (
	"context"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (h *TasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}
