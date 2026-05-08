package core_pgx_pool

import (
	"errors"
	"fmt"

	core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {

		return mapError(err)
	}

	return nil
}

type pgconnCommandTag struct {
	pgconn.CommandTag
}

// Helper func for map pgx errors to custom todo_app errors
func mapError(err error) error {
	const pgxViolateForeignKeyCode = "23503"
	if err != nil {
		if pgErr, found := errors.AsType[*pgconn.PgError](err); found {
			if pgErr.Code == pgxViolateForeignKeyCode {
				return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrViolatesForeignKey)
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrNoRows
		}
	}

	return fmt.Errorf("%v: %w", err, core_postgres_pool.ErrUnnown)
}
