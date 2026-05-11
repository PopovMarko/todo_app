package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

// GetTasks 	godoc
// @Summary 	Get tasks
// @Description Get tasks with optional pagination (user, limit, offset)
// @Tags		Tasks
// @Produce 	json
// @Param		id query int false "User ID to filter tasks by User"
// @Param 		limit query int false "window size"
// @Param		offset query int false "window offset"
// @Success		200 {object} TaskDTOResponse "List of filtered tasks"
// @Failure		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks [get]
func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, limit, offset, err := getUserLimitOffsetQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse("failed to get query params", err)
		return
	}

	domainTasks, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse("get tasks: %w", err)
		return
	}

	type GetTasksDTOResponse []TaskDTOResponse

	respons := GetTasksDTOResponse(dtoTasksFromDomain(domainTasks))
	responseHandler.JsonResponse(respons, http.StatusOK)

}

func getUserLimitOffsetQueryParam(r *http.Request) (*int, *int, *int, error) {
	const (
		userIdQueryParamKey = "id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	user, err := core_http_utils.GetIntQueryParams(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to get user id by key %s: %w",
			userIdQueryParamKey, err,
		)
	}

	limit, err := core_http_utils.GetIntQueryParams(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to get limit by key %s: %w",
			limitQueryParamKey, err,
		)
	}

	offset, err := core_http_utils.GetIntQueryParams(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"faild to get offset by key %s: %w",
			offsetQueryParamKey, err,
		)
	}

	return user, limit, offset, nil
}
