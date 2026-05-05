package users_transport_http

import (
	"net/http"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse("limit or offset are invalid", core_errors.ErrInvalidArgument)
		return
	}

	usersDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse("failed to get users", err)
	}
	userResponse := GetUsersResponse(usersDTOFromDomains(usersDomains))
	responseHandler.JsonResponse(userResponse, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParams(r, "limit")
	if err != nil {
		return nil, nil, err
	}
	offset, err := core_http_utils.GetIntQueryParams(r, "offset")
	if err != nil {
		return nil, nil, err
	}

	return limit, offset, nil
}
