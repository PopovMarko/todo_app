package web_transport_http

import (
	"net/http"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
)

func (h *WebHTTPHandler) GetHTMLPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	htmlPage, err := h.webService.GetHTMLPage()
	if err != nil {
		responseHandler.ErrorResponse("Failed to get index.html page: %w", err)
		return
	}

	responseHandler.HTMLResponse(htmlPage)
}
