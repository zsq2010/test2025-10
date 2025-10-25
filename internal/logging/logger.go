package logging

import (
	"log/slog"
	"os"
	"strings"

	"github.com/example/project/internal/config"
)

func New(cfg config.LoggingConfig) *slog.Logger {
	level := parseLevel(cfg.Level)

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
	}

	if strings.ToLower(cfg.Format) == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
