package users_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Invalid user params: %w", err)
	}
	user, err := s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("Save user repository error: %w", err)
	}
	return user, nil
}
