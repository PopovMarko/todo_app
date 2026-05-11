# TODO App

A RESTful HTTP API for managing users, tasks, and task statistics. Built with Go, PostgreSQL, and structured logging. All services — the application, database, migrations, and documentation generator — run inside Docker containers.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Makefile Reference](#makefile-reference)
- [API Reference](#api-reference)
- [Database Schema](#database-schema)
- [Project Structure](#project-structure)

---

## Features

- Full CRUD for users and tasks
- Partial updates (PATCH) with three-state field logic: omit / set / null
- Task statistics with filtering by user and date range
- Optimistic locking via `version` fields
- Structured JSON logging with `zap`
- Swagger UI served by the application itself
- Graceful shutdown on SIGINT / SIGTERM

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.26 |
| Database | PostgreSQL 18 |
| Driver | `pgx/v5` |
| Logging | `go.uber.org/zap` |
| Validation | `go-playground/validator/v10` |
| Config | `kelseyhightower/envconfig` |
| API Docs | `swaggo/swag` |
| Migrations | `migrate/migrate` |
| Containerization | Docker + Docker Compose |

---

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) 24+
- [Docker Compose](https://docs.docker.com/compose/) v2+
- `make`

No local Go installation is required. Everything runs in containers.

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/PopovMarko/todo_app.git
cd todo_app
```

### 2. Configure environment variables

Copy the example file and fill in the required values:

```bash
cp .env_example .env
```

Edit `.env` — at minimum set the Postgres credentials:

```dotenv
POSTGRES_USER=todo
POSTGRES_PASSWORD=secret
POSTGRES_DB=todo_app
```

See [Environment Variables](#environment-variables) for the full reference.

### 3. Start the database

```bash
make env-up
```

### 4. Run migrations

```bash
make migrate-up
```

### 5. Start the application

```bash
make todoapp-deploy
```

The API is now available at `http://localhost:5050`.  
Swagger UI is available at `http://localhost:5050/swagger/`.

---

## Environment Variables

All variables are loaded from the `.env` file in the project root.

| Variable | Required | Default | Description |
|---|---|---|---|
| `HTTP_ADDRESS` | yes | `:5050` | Address the HTTP server binds to |
| `HTTP_SHUTDOWN_TIMEOUT` | yes | `30s` | Graceful shutdown timeout |
| `POSTGRES_HOST` | yes | — | Database host |
| `POSTGRES_PORT` | no | `5432` | Database port |
| `POSTGRES_USER` | yes | — | Database user |
| `POSTGRES_PASSWORD` | yes | — | Database password |
| `POSTGRES_DB` | yes | — | Database name |
| `POSTGRES_TIMEOUT` | yes | `30s` | Connection timeout |
| `LOGGER_LEVEL` | no | `DEBUG` | Log level (`DEBUG`, `INFO`, `WARN`, `ERROR`) |
| `LOGGER_FOLDER` | yes | — | Directory for log file output |

---

## Makefile Reference

### Environment

| Command | Description |
|---|---|
| `make env-up` | Start the PostgreSQL container |
| `make env-down` | Stop the PostgreSQL container |
| `make env-cleanup` | Stop Postgres and delete all volume data (prompts for confirmation) |
| `make port-forward` | Expose Postgres on `127.0.0.1:5432` via socat (useful for local DB tools) |
| `make port-forward-close` | Close the port-forward container |

### Migrations

| Command | Description |
|---|---|
| `make migrate-up` | Apply all pending migrations |
| `make migrate-down` | Roll back the last migration |
| `make migrate-create seq=<name>` | Create a new sequential migration pair, e.g. `make migrate-create seq=add_index` |
| `make migrate-action action=<action>` | Run any `migrate` CLI action directly |

### Application

| Command | Description |
|---|---|
| `make todoapp-deploy` | Build the application image and start the container |
| `make ps` | List all running Compose services |
| `make log-cleanup` | Delete all log files from `out/logs/` (prompts for confirmation) |

### Documentation

| Command | Description |
|---|---|
| `make swagger-gen` | Regenerate Swagger docs from code annotations into `docs/` |

---

## API Reference

All endpoints are prefixed with `/api/v1`. Interactive documentation is available at `/swagger/` when the application is running.

### Users

| Method | Path | Description | Success |
|---|---|---|---|
| `POST` | `/users` | Create a user | `201` |
| `GET` | `/users` | List users | `200` |
| `GET` | `/users/{id}` | Get a user by ID | `200` |
| `PATCH` | `/users/{id}` | Partially update a user | `200` |
| `DELETE` | `/users/{id}` | Delete a user | `204` |

**GET /users — query parameters**

| Parameter | Type | Description |
|---|---|---|
| `limit` | integer | Page size |
| `offset` | integer | Page offset |

**PATCH /users/{id} — body**

Fields are three-state: omitting a key leaves the field unchanged; setting it to `null` clears it (where allowed); providing a value updates it.

```json
{
  "full_name": "Jane Doe",
  "phone_number": "+12025550147"
}
```

> `full_name` cannot be set to `null`. `phone_number` can be nulled.

---

### Tasks

| Method | Path | Description | Success |
|---|---|---|---|
| `POST` | `/tasks` | Create a task | `201` |
| `GET` | `/tasks` | List tasks | `200` |
| `GET` | `/tasks/{id}` | Get a task by ID | `200` |
| `PATCH` | `/tasks/{id}` | Partially update a task | `200` |
| `DELETE` | `/tasks/{id}` | Delete a task | `204` |

**GET /tasks — query parameters**

| Parameter | Type | Description |
|---|---|---|
| `id` | integer | Filter by author user ID |
| `limit` | integer | Page size |
| `offset` | integer | Page offset |

**PATCH /tasks/{id} — body**

```json
{
  "title": "Updated title",
  "description": null,
  "completed": true
}
```

> `title` cannot be set to `null`. `description` can be nulled. `completed` must be a boolean if provided.

---

### Statistics

| Method | Path | Description | Success |
|---|---|---|---|
| `GET` | `/statistics` | Get task statistics | `200` |

**Query parameters**

| Parameter | Type | Description |
|---|---|---|
| `id` | integer | Filter by user ID |
| `from` | date | Start of date range |
| `to` | date | End of date range |

**Response**

```json
{
  "completed_tasks": 12,
  "total_tasks": 20,
  "tasks_completion_rate": 0.6,
  "tasks_average_completion_time": "2h30m"
}
```

---

## Database Schema

Migrations live in `migrations/` and are applied with `make migrate-up`.

### `todoapp.users`

| Column | Type | Constraints |
|---|---|---|
| `id` | `SERIAL` | Primary key |
| `version` | `BIGINT` | Optimistic lock, default `1` |
| `full_name` | `VARCHAR(100)` | Not null, length 3–100 |
| `phone_number` | `VARCHAR(15)` | Nullable, format `+[digits]`, length 10–15 |

### `todoapp.tasks`

| Column | Type | Constraints |
|---|---|---|
| `id` | `SERIAL` | Primary key |
| `version` | `BIGINT` | Optimistic lock, default `1` |
| `title` | `VARCHAR(100)` | Not null, length 1–100 |
| `description` | `VARCHAR(1000)` | Nullable, length 1–1000 |
| `completed` | `BOOL` | Not null |
| `created_at` | `TIMESTAMPTZ` | Not null |
| `completed_at` | `TIMESTAMPTZ` | Nullable; required when `completed = true`; must be ≥ `created_at` |
| `author_user_id` | `INTEGER` | Foreign key → `users.id` |

---

## Project Structure

```
.
├── cmd/
│   └── todo/
│       ├── main.go          # Entry point, dependency wiring
│       └── Dockerfile       # Multi-stage build (Go builder + Alpine runtime)
├── internal/
│   ├── core/                # Shared, feature-agnostic code
│   │   ├── domain/          # Domain models and patch logic
│   │   ├── errors/          # Sentinel errors
│   │   ├── logger/          # Zap logger setup
│   │   ├── repository/      # Database connection pool
│   │   └── transport/http/  # Server, middleware, request/response helpers
│   └── features/            # Feature modules (users, tasks, statistics)
│       └── <feature>/
│           ├── repository/postgres/  # SQL queries
│           ├── service/              # Business logic
│           └── transport/http/       # HTTP handlers and DTOs
├── migrations/              # SQL migration files
├── docs/                    # Auto-generated Swagger/OpenAPI specs
├── out/
│   ├── logs/                # Application log output
│   └── pgdata/              # Postgres data volume
├── docker-compose.yaml
├── Makefile
├── .env_example
├── go.mod
└── go.sum
```
