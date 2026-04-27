package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

// Func to reconize errors and establish status code
func (h *HTTPResponseHandler) ErrorResponse(msg string, err error) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))
	h.errorResponse(msg, statusCode, err)
}

func (h *HTTPResponseHandler) PanicResponse(msg string, p any) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.errorResponse(msg, statusCode, err)
}

func (h *HTTPResponseHandler) errorResponse(msg string, statusCode int, err error) {
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}

}
