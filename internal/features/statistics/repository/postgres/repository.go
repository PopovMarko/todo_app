package statistics_postgres_repository

import core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"

type StatisticsRepository struct {
	Pool core_postgres_pool.Pool
}

func NewStatisticsRepository(pool core_postgres_pool.Pool) *StatisticsRepository {
	return &StatisticsRepository{
		Pool: pool,
	}
}
