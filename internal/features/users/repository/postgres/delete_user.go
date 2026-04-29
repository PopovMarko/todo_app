package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.users
	WHERE id = $1
	`
	commandTag, err := r.Pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("User with ID %d not found: %w", userID, core_errors.ErrNotFound)
	}

	return nil
}
