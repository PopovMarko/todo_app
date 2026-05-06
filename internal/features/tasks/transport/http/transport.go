package tasks_transport_http

import (
	"context"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
}

type TasksHTTPHandler struct {
	tasksService TasksService
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}
