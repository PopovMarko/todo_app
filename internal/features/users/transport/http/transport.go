package users_transport_http

import (
	"context"
	"net/http"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
)

// Interface of the Service layer. Transport delend on
type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

// Handl requests for User on transport layer
// Represent user transport
type UserHTTPHandler struct {
	userService UserService
}

// Handler constructor
func NewUserHTTPHandler(userService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

// Method returns list of routes of the feature functionality
func (h *UserHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
