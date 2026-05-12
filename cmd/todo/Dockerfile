FROM golang:1.26.3-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/todo/todo /app/cmd/todo

FROM alpine:3.22.4
WORKDIR /app
COPY --from=builder /app/web /app/cmd/todo/web
COPY --from=builder /app/cmd/todo /app/cmd/todo
ENV PROJECT_ROOT=/app/cmd/todo
CMD ["/app/cmd/todo/todo"]
