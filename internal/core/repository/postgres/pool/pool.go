package core_postgres_pool

import (
	"context"
	"time"
)

// Connection pool interface, postgres pool depends on
type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error)
	Close()

	OpTimeout() time.Duration
}

// Interface that will be returned in Query method of the Pool interface
// for reverse dependencies of pgx library
type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Scan(dest ...any) error
}

type CommandTag interface {
	RowsAffected() int64
}
