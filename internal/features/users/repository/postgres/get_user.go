package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_postgres_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	WHERE id = $1
	`
	row := r.Pool.QueryRow(ctx, query, userID)

	var userModel UserModel

	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with ID %d - not found: %w", userID, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("failed to scan user: %w", err)
	}

	return userDomainFromModel(userModel), nil
}
