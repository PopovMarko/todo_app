package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

// DeleteTask  	godoc
// @Summary 	Delete task
// @Description Delete task by task ID
// @Tags 		Tasks
// @Accept		json
// @Produce 	json
// @Param		id path int true "Task ID"
// @Success		204 "Task deleted successfully"
// @Failure		404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure		500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router		/tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	taskID, err := core_http_utils.GetIntPathParams(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("failed to get task ID", err)
		return
	}

	if err := h.tasksService.DeleteTask(ctx, *taskID); err != nil {
		responseHandler.ErrorResponse("failed to delete task", err)
		return
	}

	responseHandler.NoContentResponse(http.StatusNoContent)
}
