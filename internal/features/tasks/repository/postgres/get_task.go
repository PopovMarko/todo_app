package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) GetTask(ctx context.Context, taskID int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	WHERE id = $1
	`

	var model TaskModel
	row := r.Pool.QueryRow(ctx, query, taskID)

	if err := row.Scan(
		&model.ID,
		&model.Version,
		&model.Title,
		&model.Description,
		&model.Completed,
		&model.CreatedAt,
		&model.CompletedAt,
		&model.AuthorUserID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with ID %d not found: %v: %w",
				taskID, err, core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("Scan task: %w", err)
	}

	return modelToDomain(model), nil
}
