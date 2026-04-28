include .env 
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d todoapp-postgres
	
env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Clean up all volumes environment. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres todoapp-postgres-port-forwarder && \
		rm -rf out/pgdata && \
		echo "Volumes cleaned up"; \
	else \
		echo "Volumes cleanup cancelled"; \
	fi

port-forward:
	@docker compose up -d todoapp-postgres-port-forwarder 
port-forward-close:
	@docker compose down todoapp-postgres-port-forwarder
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Error: No seq parameter. Usage: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Error: no action parameter. Usage make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down
todo-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ./cmd/todo/main.go

