package users_service

import (
	"context"
	"fmt"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, userID int, userPatch domain.UserPatch) (domain.User, error) {
	// Get user from repository
	user, err := s.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	// Try to apply patch to user
	if err := user.ApplyPatch(userPatch); err != nil {
		return domain.User{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	// Save patched user to repository
	user, err = s.UserRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to save patched user: %w", err)
	}

	return user, nil

}
