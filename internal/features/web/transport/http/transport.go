package web_transport_http

import (
	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
)

type WebService interface {
	GetHTMLPage() ([]byte, error)
}
type WebHTTPHandler struct {
	webService WebService
}

func NewWebHTTPHandler(webService WebService) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetHTMLPage,
		},
	}
}
