package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_request "github.com/PopovMarko/todo_app/internal/core/transport/http/request"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_types "github.com/PopovMarko/todo_app/internal/core/transport/http/types"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (p *PatchTaskRequest) Validate() error {
	if p.Title.Set {
		if p.Title.Value == nil {
			return fmt.Errorf("title can't patched to NULL: %w", core_errors.ErrInvalidArgument)
		}
		titlelen := len([]rune(*p.Title.Value))
		if titlelen < 3 || titlelen > 100 {
			return fmt.Errorf(
				"title length must be between 3 and 100 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.Description.Set && p.Description.Value != nil {
		descriptionlen := len([]rune(*p.Description.Value))
		if descriptionlen < 1 || descriptionlen > 1000 {
			return fmt.Errorf(
				"description length must be between 1 and 1000 characters: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf(
			"completed can't be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

// PatchTask 	godoc
// @Summary 	Patch task
// @Description	Patch task with new information
// @Description ### Logic of fields patch (three-state logic)
// @Description	1. **field not sent** `description` ignored, value in DB not change
// @Description	2. **field value sent explicitly** `"description": "any new description"` sets new value in DB
// @Description	3. **field set to null** `"description": null` clears value in DB
// @Description **Restriction** `title` can't be set to null
// @Tags 		Tasks
// @Accept		json
// @Produce		json
// @Param		id path int true "Task ID"
// @Param		request body PatchTaskRequest true "Patch information"
// @Success		200 {object} TaskDTOResponse "Patched task copy"
// @Failure		404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 	500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		/tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	taskID, err := core_http_utils.GetIntPathParams(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("failed to get task ID", err)
		return
	}

	var taskDto PatchTaskRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &taskDto); err != nil {
		responseHandler.ErrorResponse("decode and validate request: %w", err)
		return
	}

	taskPatch := domain.NewTaskPatch(
		taskDto.Title.ToDomain(),
		taskDto.Description.ToDomain(),
		taskDto.Completed.ToDomain(),
	)

	domainTask, err := h.tasksService.PatchTask(ctx, *taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse("patch task", err)
		return
	}

	responsePatch := TaskDTOResponse(dtoTaskFromDomain(domainTask))
	responseHandler.JsonResponse(responsePatch, http.StatusOK)

}
