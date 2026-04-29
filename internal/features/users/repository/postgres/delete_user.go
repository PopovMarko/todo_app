package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM todoapp.users
	WHERE id = $1
	`
	_, err := r.Pool.Exec(ctx, query, userID)

	if err != nil {
		return fmt.Errorf("failed to scan user: %w", err)
	}

	return nil
}
