package tasks_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (s *TasksService) GetTask(ctx context.Context, taskID int) (domain.Task, error) {

	domainTask, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to get task from repository: %w", err)
	}
	return domainTask, nil
}
