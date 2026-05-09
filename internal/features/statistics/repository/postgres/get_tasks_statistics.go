package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (r *StatisticsRepository) GetTasksStatistics(ctx context.Context, userID *int, from *time.Time, to *time.Time) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT id, version, title, description, completed, created_at, completed_at, author_user_id
	FROM todoapp.tasks
	`)

	var (
		whereString []string
		args        []any
	)
	if userID != nil {
		args = append(args, *userID)
		whereString = append(whereString, fmt.Sprintf("author_user_id=$%d", len(args)))
	}
	if from != nil {
		args = append(args, *from)
		whereString = append(whereString, fmt.Sprintf("created_at>=$%d", len(args)))
	}
	if to != nil {
		args = append(args, *to)
		whereString = append(whereString, fmt.Sprintf("created_at<$%d", len(args)))
	}

	if len(whereString) > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" WHERE %s ", strings.Join(whereString, " AND ")))
	}

	rows, err := r.Pool.Query(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks statistics: %w", err)
	}
	defer rows.Close()

	var tasks []TaskModel

	for rows.Next() {
		var task TaskModel
		if err := rows.Scan(
			&task.ID,
			&task.Version,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
			&task.AuthorUserID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to rows next: %w", err)
	}
	domainTasks := modelsToDomains(tasks)

	return domainTasks, nil
}
