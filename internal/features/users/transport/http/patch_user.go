package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PopovMarko/todo_app/internal/core/domain"
	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_request "github.com/PopovMarko/todo_app/internal/core/transport/http/request"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_types "github.com/PopovMarko/todo_app/internal/core/transport/http/types"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

type (
	PatchUserResponse UserDTOResponse
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (p *PatchUserRequest) Validate() error {
	if p.FullName.Set {
		if p.FullName.Value == nil {
			return fmt.Errorf("transport custom validate full name can't be NULL: %w", core_errors.ErrInvalidArgument)
		}
		fullNameLen := len([]rune(*p.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("full name length must be between 3 and 100 characters: %w", core_errors.ErrInvalidArgument)
		}
	}

	if p.PhoneNumber.Value != nil {
		phoneNumberLen := len([]rune(*p.PhoneNumber.Value))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("phone number must be between 10 and 15 characters: %w", core_errors.ErrInvalidArgument)
		}
		if !strings.HasPrefix(*p.PhoneNumber.Value, "+") {
			return fmt.Errorf("phone humber must start with +: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

func (h *UserHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, err := core_http_utils.GetIntPathParams(r, "id")
	if err != nil {
		responseHandler.ErrorResponse("Failed to get user id from request", err)
	}

	var request PatchUserRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse("decode and validate request:", err)
		return
	}

	userPatch := userPatchFromRequest(request)
	user, err := h.userService.PatchUser(ctx, *userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse("patch User", err)
		return
	}
	response := PatchUserResponse(userDTOFromDomain(user))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
