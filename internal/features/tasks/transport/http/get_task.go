package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

// GetTask 		godoc
// @Summary 	Get task
// @Description Get task by ID
// @Tags 		Tasks
// @Produce		json
// @Param		id path int true "Task ID"
// @Success		200 {object} TaskDTOResponse "Task information"
// @Failure		404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 	400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router		/tasks/{id} [get]
func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	const taskPathParamKey = "id"
	taskID, err := core_http_utils.GetIntPathParams(r, taskPathParamKey)
	if err != nil {
		responseHandler.ErrorResponse("failed to get task ID from path param", err)
		return
	}

	type getTaskResponse TaskDTOResponse

	taskDomain, err := h.tasksService.GetTask(ctx, *taskID)
	if err != nil {
		responseHandler.ErrorResponse("failed to get task from service", err)
		return
	}

	taskResponse := getTaskResponse(dtoTaskFromDomain(taskDomain))
	responseHandler.JsonResponse(taskResponse, http.StatusOK)
}
