package api

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(logger *slog.Logger) *mux.Router {
	handler := NewHandler(logger)
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health).Methods("GET")

	router.Use(loggingMiddleware(logger))

	return router
}

func loggingMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("incoming request",
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
			)
			next.ServeHTTP(w, r)
		})
	}
}
