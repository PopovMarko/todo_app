package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_middleware "github.com/PopovMarko/todo_app/internal/core/transport/http/middleware"
	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
	users_transport_http "github.com/PopovMarko/todo_app/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	loggerConfig := core_logger.NewLoggerConfigMust()
	logger, err := core_logger.NewLogger(loggerConfig)
	if err != nil {
		fmt.Printf("Logger configuration failed: %v", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Starting todo application")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	usersTransportHTTP := users_transport_http.NewUserHTTPHandler(nil)
	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersRoutes...)

	config := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		config,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Info("Server error", zap.Error(err))
	}

}
