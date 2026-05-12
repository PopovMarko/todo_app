package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PopovMarko/todo_app/docs"
	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_middleware "github.com/PopovMarko/todo_app/internal/core/transport/http/middleware"
	"go.uber.org/zap"

	"github.com/swaggo/http-swagger"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	logger     *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middleware ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		logger:     log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.Version)

		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}

func (s *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		s.mux.Handle(pattern, route.Handler)
	}
}
func (s *HTTPServer) RegisterSwagger() {
	s.mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	s.mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		swaggerDoc := docs.SwaggerInfo.ReadDoc()
		_, _ = w.Write([]byte(swaggerDoc))
	})

}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middleware...)
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	s.logger.Warn("Starting HTTP server...", zap.String("on addr", s.config.Addr))

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}

	}()

	// Select blocks and wait for error from error channel or context cancelation
	select {
	case err := <-ch:
		if err != nil {
			s.logger.Error("during starting HTTP server", zap.Error(err))
			return fmt.Errorf("Listen and serve HTTP: %w", err)
		}

	case <-ctx.Done():
		s.logger.Warn("Shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			s.logger.Error("shutting down HTTP server error", zap.Error(err))
			return fmt.Errorf("shutting down HTTP server error: %w", err)
		}
		s.logger.Info("Server closed gracefully")
	}

	return nil
}
