package statistics_transport_http

import "github.com/PopovMarko/todo_app/internal/core/domain"

type StatisticsDTOResponse struct {
	TotalTasks                int      `json:"total_tasks"`
	CompletedTasks            int      `json:"completed_tasks"`
	TasksCompletionRate       *float64 `json:"tasks_completion_rate"`
	TasksAverageCompetionTime *string  `json:"tasks_average_completion_time"`
}

func dtoStatisticsFromDomain(domainStatistics domain.Statistics) StatisticsDTOResponse {
	var tasksAverageCompletionString *string
	if domainStatistics.TasksAverageCompletionTime != nil {
		duration := domainStatistics.TasksAverageCompletionTime.String()
		tasksAverageCompletionString = &duration
	}
	return StatisticsDTOResponse{
		TotalTasks:                domainStatistics.TotalTasks,
		CompletedTasks:            domainStatistics.CompletedTasks,
		TasksCompletionRate:       domainStatistics.TasksCompletionRate,
		TasksAverageCompetionTime: tasksAverageCompletionString,
	}
}
