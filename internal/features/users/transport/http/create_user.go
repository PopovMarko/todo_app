package users_transport_http

import (
	"fmt"
	"net/http"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_request "github.com/PopovMarko/todo_app/internal/core/transport/http/request"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
)

// DTO for parse user from request and get to service layer
type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100" `
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

// DTO for get user from service layer and send to http
type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

// Method of user handler that in transport.go
func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	var requestUser CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &requestUser); err != nil {
		err = fmt.Errorf("%w", core_errors.ErrInvalidArgument)
		responseHandler.ErrorResponse("decode or validate user DTO error", err)

		return
	}
	user := domainFromDTO(requestUser)
	user, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		responseHandler.ErrorResponse("failed to create user", err)
		return
	}

	userResponseDTO := dtoFromDomain(user)
	responseHandler.JsonResponse(userResponseDTO, http.StatusCreated)

}

// Helper func to connect domain and transport without
// importing each other
func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}
