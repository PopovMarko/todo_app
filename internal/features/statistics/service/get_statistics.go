package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil && to.Before(*from) {
		return domain.Statistics{}, fmt.Errorf(
			"Invalid time range: 'to' time is before 'from' time: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	domainTasks, err := s.statisticsRepository.GetTasksStatistics(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("Failed to get tasks for statistics: %w", err)
	}

	domainStat, err := statCalc(domainTasks)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("failed to calculate statistics %w", err)
	}

	return domainStat, nil
}

func statCalc(domainTasks []domain.Task) (domain.Statistics, error) {
	totalTasks := len(domainTasks)
	if totalTasks == 0 {
		return *domain.NewStatistics(0, 0, nil, nil), nil
	}

	var (
		completedTasks             int
		completedTasksTime         time.Duration
		tasksAverageCompletionTime time.Duration
	)

	for _, task := range domainTasks {
		if task.Completed {
			completedTasks++
			if task.CompletedAt != nil {
				completedTasksTime += task.CompletedAt.Sub(task.CreatedAt)
			}
		}
	}

	tasksCompletionRate := float64(completedTasks) / float64(totalTasks) * 100
	if completedTasks > 0 {
		averageCompletionTime := completedTasksTime / time.Duration(completedTasks)
		tasksAverageCompletionTime = time.Duration(averageCompletionTime)
	}

	return *domain.NewStatistics(
		totalTasks,
		completedTasks,
		&tasksCompletionRate,
		&tasksAverageCompletionTime,
	), nil
}
