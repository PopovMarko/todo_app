package users_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

func (h *UserHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathParams(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("delete user from request", err)
		return
	}

	err = h.userService.DeleteUser(ctx, *userID)
	if err != nil {
		responseHandler.ErrorResponse("delete user from service", err)
		return
	}

	responseHandler.NoContentResponse(http.StatusNoContent)
}
