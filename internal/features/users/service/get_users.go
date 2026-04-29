package users_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit can't be negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset can't be negative: %w", core_errors.ErrInvalidArgument)
	}

	domainUsers, err := s.UserRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users from repository: %w", err)
	}

	return domainUsers, nil
}
