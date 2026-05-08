package core_postgres_pool

import "errors"

var (
	ErrNoRows             = errors.New("no rows affected")
	ErrViolatesForeignKey = errors.New("violates foreign key")
	ErrUnnown             = errors.New("unnown error")
)
