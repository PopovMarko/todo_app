package core_http_middleware

import (
	"net/http"
	"time"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestUserIDHeader = "X-Request-ID"

// Middleware function to check user ID and add it if not provided
func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestUserIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestUserIDHeader, requestID)
			w.Header().Set(requestUserIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

// Middleware function to put the Logger into Context of the Request
func Logger(l *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestUserIDHeader)

			//Logger preconfigure for particular request with request id and url
			log := l.With(
				zap.String("request ID: ", requestID),
				zap.String("URL: ", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), log)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Middleware function to trace latency and status code
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.LogFromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()
			logger.Debug(">>> Incomin HTTP request",
				zap.String("Method", r.Method),
				zap.Time("Time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			// Get status code from custom ResponseWriter after it returns from end handler
			statusCode := rw.GetStatusCode()
			logger.Debug("<<< Outgoing HTTP response",
				zap.Duration("Latency", time.Since(before)),
				zap.Int("status", statusCode),
			)
		})
	}
}

// Middleware function to panic recover
func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.LogFromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse("during handle HTTP request got some panic", p)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
