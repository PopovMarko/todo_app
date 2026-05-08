package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (r *TasksRepository) GetTasks(ctx context.Context, userID, limit, offset *int) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	%s
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	args := []any{limit, offset}
	if userID == nil {
		query = fmt.Sprintf(query, "")
	} else {
		args = append(args, userID)
		query = fmt.Sprintf(query, "WHERE author_user_id = $3")
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()

	var modelsTasks []TaskModel

	for rows.Next() {
		var model TaskModel

		err := rows.Scan(
			&model.ID,
			&model.Version,
			&model.Title,
			&model.Description,
			&model.Completed,
			&model.CreatedAt,
			&model.CompletedAt,
			&model.AuthorUserID,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan tasks: %w", err)
		}

		modelsTasks = append(modelsTasks, model)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to scan tasks: %w", err)
	}

	return modelsToDomains(modelsTasks), nil
}
