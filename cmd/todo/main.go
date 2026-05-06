package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_pgx_pool "github.com/PopovMarko/todo_app/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/PopovMarko/todo_app/internal/core/transport/http/middleware"
	core_http_server "github.com/PopovMarko/todo_app/internal/core/transport/http/server"
	tasks_postgres_repository "github.com/PopovMarko/todo_app/internal/features/tasks/repository/postgres"
	tasks_service "github.com/PopovMarko/todo_app/internal/features/tasks/service"
	tasks_transport_http "github.com/PopovMarko/todo_app/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/PopovMarko/todo_app/internal/features/users/repository/postgres"
	users_service "github.com/PopovMarko/todo_app/internal/features/users/service"
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

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger.Debug("Initializing connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to initialize connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("Initializing feature", zap.String("feature", "users"))

	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	userService := users_service.NewService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUserHTTPHandler(userService)

	logger.Debug("Initializing feature", zap.String("feature", "tasks"))

	taskRepository := tasks_postgres_repository.NewTaskRepository(pool)
	tasksService := tasks_service.NewTasksService(taskRepository)
	tasksHTTPHandler := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("Initializing HTTP server")

	config := core_http_server.NewConfigMust()
	httpServer := core_http_server.NewHTTPServer(
		config,
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouter.RegisterRoutes(tasksHTTPHandler.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	logger.Debug("Starting todo application")

	if err := httpServer.Run(ctx); err != nil {
		logger.Info("Server error", zap.Error(err))
	}

}
