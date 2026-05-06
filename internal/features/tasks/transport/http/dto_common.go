package tasks_transport_http

import (
	"github.com/PopovMarko/todo_app/internal/core/domain"
	"time"
)

type TaskHTTPRequest struct {
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	AuthorUserID int     `json:"autnor_user_id"`
}

func domainTaskFromDTO(dto TaskHTTPRequest) domain.Task {
	return domain.NewUninitializeTask(dto.Title, dto.Description, dto.AuthorUserID)
}

type TaskHTTPResponse struct {
	ID           int        `json:"task_id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func dtoTaskFromDomain(task domain.Task) TaskHTTPResponse {
	return TaskHTTPResponse{
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
