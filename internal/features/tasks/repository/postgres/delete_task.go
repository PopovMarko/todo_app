package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, taskID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.tasks 
	WHERE id = $1
	`
	commandTag, err := r.Pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task with ID %d: %w", taskID, err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("task with ID %d not found: %w", taskID, core_errors.ErrNotFound)
	}

	return nil
}
