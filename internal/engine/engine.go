package engine

import (
	"context"
	"log/slog"

	_ "github.com/chromedp/chromedp"
)

type Engine struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Engine {
	return &Engine{
		logger: logger,
	}
}

func (e *Engine) Start(ctx context.Context) error {
	e.logger.Info("engine started")
	return nil
}

func (e *Engine) Stop(ctx context.Context) error {
	e.logger.Info("engine stopped")
	return nil
}
