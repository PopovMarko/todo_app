package core_http_response

import "net/http"

var (
	UninitializedStatus = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     UninitializedStatus,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriter) GetStatusCode() int {
	if w.statusCode == UninitializedStatus {
		return http.StatusOK
	}
	return w.statusCode
}
