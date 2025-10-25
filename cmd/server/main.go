package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/project/internal/api"
	"github.com/example/project/internal/config"
	"github.com/example/project/internal/engine"
	"github.com/example/project/internal/logging"
	"github.com/example/project/internal/tasks"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	logger := logging.New(cfg.Logging)
	slog.SetDefault(logger)

	logger.Info("starting server",
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
		"log_level", cfg.Logging.Level,
	)

	taskManager := tasks.NewManager(logger, cfg.BasePaths.TasksDir)
	eng := engine.New(logger)

	ctx := context.Background()
	if err := taskManager.Start(ctx); err != nil {
		return fmt.Errorf("failed to start task manager: %w", err)
	}

	if err := eng.Start(ctx); err != nil {
		return fmt.Errorf("failed to start engine: %w", err)
	}

	router := api.NewRouter(logger)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("server listening", "address", addr)
		serverErrors <- srv.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		logger.Info("shutdown signal received", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
			return fmt.Errorf("failed to gracefully shutdown server: %w", err)
		}

		logger.Info("shutting down components")
		if err := eng.Stop(ctx); err != nil {
			logger.Error("failed to stop engine", "error", err)
		}
		if err := taskManager.Stop(ctx); err != nil {
			logger.Error("failed to stop task manager", "error", err)
		}

		logger.Info("server stopped gracefully")
	}

	return nil
}
