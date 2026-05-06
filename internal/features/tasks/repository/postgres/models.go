package tasks_postgres_repository

import (
	"time"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func modelToDomain(model TaskModel) domain.Task {
	return domain.Task{
		ID:           model.ID,
		Version:      model.Version,
		Title:        model.Title,
		Description:  model.Description,
		Completed:    model.Completed,
		CreatedAt:    model.CreatedAt,
		CompletedAt:  model.CompletedAt,
		AuthorUserID: model.AuthorUserID,
	}
}
