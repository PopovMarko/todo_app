package tasks_postgres_repository

import core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"

type TasksRepository struct {
	Pool core_postgres_pool.Pool
}

func NewTaskRepository(pool core_postgres_pool.Pool) *TasksRepository {
	return &TasksRepository{
		Pool: pool,
	}
}
