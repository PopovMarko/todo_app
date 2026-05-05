package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (r *UsersRepository) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.Pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, phone_number
	FROM todoapp.users
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from database: %w", err)
	}
	defer rows.Close()

	var modelsUsers []UserModel

	for rows.Next() {
		var userModel UserModel
		if err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
		); err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}
		modelsUsers = append(modelsUsers, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan user model: %w", err)
	}

	return usersDomainsFromModels(modelsUsers), nil
}
