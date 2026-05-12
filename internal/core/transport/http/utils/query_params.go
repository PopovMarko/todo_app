package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

func GetIntQueryParams(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("parameter %s, by key %s - not a valid integer: %v: %w",
			param, key, err, core_errors.ErrInvalidArgument)
	}

	return &val, nil
}

func GetTimeQueryParams(r *http.Request, key string) (*time.Time, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	layout := "01-02-2006"
	val, err := time.Parse(layout, param)
	if err != nil {
		return nil, fmt.Errorf("param %s, by key %s - not a valid time: %v: %w",
			param, key, err, core_errors.ErrInvalidArgument,
		)
	}
	return &val, nil
}
