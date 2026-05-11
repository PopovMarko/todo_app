package domain

import (
	"time"
)

type Statistics struct {
	TotalTasks                 int
	CompletedTasks             int
	TasksCompletionRate        *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	totalTasks int,
	completedTasks int,
	tasksCompletionRate *float64,
	tasksAverageCompletionTime *time.Duration,
) *Statistics {
	return &Statistics{
		TotalTasks:                 totalTasks,
		CompletedTasks:             completedTasks,
		TasksCompletionRate:        tasksCompletionRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
