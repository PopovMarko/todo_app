package users_transport_http

import (
	"net/http"

	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
)

// Interface of the Service layer. Transport delend on
type UserService interface {
}

// Handl requests for User on transport layer
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
