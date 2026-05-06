package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_request "github.com/PopovMarko/todo_app/internal/core/transport/http/request"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
)

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	var task = TaskHTTPRequest{}
	if err := core_http_request.DecodeAndValidateRequest(r, task); err != nil {
		responseHandler.ErrorResponse("decode and validate request", err)
	}

	domainTask := domainTaskFromDTO(task)

	domainTask, err := h.tasksService.CreateTask(ctx, domainTask)
	if err != nil {
		responseHandler.ErrorResponse("cerate task", err)
	}
	taskResponse := dtoTaskFromDomain(domainTask)
	responseHandler.JsonResponse(taskResponse, http.StatusCreated)
}
