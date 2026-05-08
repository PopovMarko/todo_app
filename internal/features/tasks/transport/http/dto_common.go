package tasks_transport_http

import (
	"github.com/PopovMarko/todo_app/internal/core/domain"
	"time"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=3,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type TaskDTOResponse struct {
	ID           int        `json:"task_id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func dtoTaskFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func dtoTasksFromDomain(tasks []domain.Task) []TaskDTOResponse {
	resp := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		resp[i] = dtoTaskFromDomain(task)
	}
	return resp
}
