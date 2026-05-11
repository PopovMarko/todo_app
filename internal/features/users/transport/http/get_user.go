package users_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary 		Get user by ID
// @Description 	Get user information by user ID
// @Tags 			Users
// @Accept 			json
// @Produce 		json
// @Param			id path int true "User ID"
// @Success 		200 {object} GetUserResponse "User information retrieved successfully"
// @Failure			400 {object} core_http_response.ErrorResponse "Invalid request"
// @Failure			404 {object} core_http_response.ErrorResponse "User not found"
// @Failure			500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router			/users/{id} [get]
func (h *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathParams(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("get user id from request: %w", err)
		return
	}

	userDomain, err := h.userService.GetUser(ctx, *userID)
	if err != nil {
		responseHandler.ErrorResponse("get user from service: %w", err)
		return
	}

	userResponse := GetUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JsonResponse(userResponse, http.StatusOK)
}
