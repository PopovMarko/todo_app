package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func GetIntPathParams(r *http.Request, key string) (*int, error) {
	value := r.PathValue(key)
	if value == "" {
		return nil, fmt.Errorf("empty path parameter: %w", core_errors.ErrInvalidArgument)
	}
	param, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("invalid path parameter: %v, %w", err, core_errors.ErrInvalidArgument)
	}

	return &param, nil
}
