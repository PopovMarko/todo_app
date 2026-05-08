package tasks_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, taskID int, patch domain.TaskPatch) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}

	if err := task.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply patch: %w", err)
	}

	patchedTask, err := s.tasksRepository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("failed to save patched task: %w", err)
	}
	return patchedTask, nil
}
