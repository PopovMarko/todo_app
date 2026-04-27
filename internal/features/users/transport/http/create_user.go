package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	"go.uber.org/zap"
)

// DTO for parse user from request and get to service layer
type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

// DTO for get user from service layer and send to http
type CreateUserResponse struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

// Method of user handler that in transport.go
func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	logger.Debug("Create User method called")

	var request CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("create user", zap.Error(err))
	}
	w.WriteHeader(http.StatusOK)

}
