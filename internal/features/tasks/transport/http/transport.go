package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
)

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, userID, limit, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
	PatchTask(ctx context.Context, taskID int, patch domain.TaskPatch) (domain.Task, error)
}

type TasksHTTPHandler struct {
	tasksService TasksService
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/tasks/{id}",
			Handler: h.PatchTask,
		},
	}
}
