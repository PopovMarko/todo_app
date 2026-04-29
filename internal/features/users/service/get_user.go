package users_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, userID int) (domain.User, error) {
	userDomain, err := s.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from reporitory: %w", err)
	}
	return userDomain, nil
}
