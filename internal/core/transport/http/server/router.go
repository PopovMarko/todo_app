package core_http_server

import (
	"fmt"
	"net/http"
)

type APIVersionRouter struct {
	*http.ServeMux
	Version APIVersion
}

type APIVersion string

var (
	APIVersion1 = APIVersion("v1")
	AtiVersion2 = APIVersion("v2")
)

func NewAPIVersionRouter(version APIVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux: http.NewServeMux(),
		Version:  version,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		r.Handle(pattern, route.Handler)
	}

}
