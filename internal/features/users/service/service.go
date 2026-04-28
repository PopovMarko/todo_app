package users_service

import (
	"context"

	"github.com/PopovMarko/todo_app/internal/core/domain"
)

// Repository interface, service layer depends on
type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

// Represent service layer
type UsersService struct {
	UserRepository UserRepository
}

func NewService(userRepository UserRepository) *UsersService {
	return &UsersService{
		UserRepository: userRepository,
	}
}
