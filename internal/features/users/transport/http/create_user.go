package users_transport_http

import (
	"fmt"
	"net/http"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_request "github.com/PopovMarko/todo_app/internal/core/transport/http/request"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
)

// CreateUser godoc
// @Summary 		Create a new user
// @Description 	Create a new user with the provided information
// @Tags 			Users
// @Accept 			json
// @Produce 		json
// @Param 			request body UserDTORequest true "User information"
// @Success 		201 {object} UserDTOResponse "User created successfully"
// @Failure 		400 {object} core_http_response.ErrorResponse "Invalid request"
// @Failure 		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 			/users [post]
func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	type (
		CreateUserResponse UserDTOResponse
		CreateUserRequest  UserDTORequest
	)
	var requestUser CreateUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &requestUser); err != nil {
		err = fmt.Errorf("%w", core_errors.ErrInvalidArgument)
		responseHandler.ErrorResponse("decode or validate user DTO error", err)

		return
	}
	user := domainFromDTO(UserDTORequest(requestUser))
	user, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		responseHandler.ErrorResponse("failed to create user", err)
		return
	}

	userResponseDTO := CreateUserResponse(userDTOFromDomain(user))
	responseHandler.JsonResponse(userResponseDTO, http.StatusCreated)
}
