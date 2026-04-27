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

func (w *ResponseWriter) GetStatusCodeOrPanic() int {
	if w.statusCode == UninitializedStatus {
		panic("Status code uninitialized")
	}
	return w.statusCode
}
