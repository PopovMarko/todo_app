package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"
)

func (r *TasksRepository) PatchTask(ctx context.Context, taskID int, patch domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.tasks	
		SET 
		version = version + 1,
		title = $1, 
		description = $2, 
		completed = $3,
		completed_at = $4
		WHERE id = $5 AND version = $6
		RETURNING id, version, title, description, completed, created_at, completed_at, author_user_id;
	`

	var patchedTask TaskModel
	row := r.Pool.QueryRow(
		ctx,
		query,
		patch.Title,
		patch.Description,
		patch.Completed,
		patch.CompletedAt,
		taskID,
		patch.Version,
	)

	if err := row.Scan(
		&patchedTask.ID,
		&patchedTask.Version,
		&patchedTask.Title,
		&patchedTask.Description,
		&patchedTask.Completed,
		&patchedTask.CreatedAt,
		&patchedTask.CompletedAt,
		&patchedTask.AuthorUserID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with ID %d concurrently accessed: %v: %w",
				taskID, err, core_errors.ErrConflict,
			)
		}
		return domain.Task{}, fmt.Errorf(
			"scan patched task: %w",
			err,
		)
	}
	return modelToDomain(patchedTask), nil
}
