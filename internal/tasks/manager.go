package tasks

import (
	"context"
	"log/slog"
)

type Manager struct {
	logger   *slog.Logger
	tasksDir string
}

func NewManager(logger *slog.Logger, tasksDir string) *Manager {
	return &Manager{
		logger:   logger,
		tasksDir: tasksDir,
	}
}

func (m *Manager) Start(ctx context.Context) error {
	m.logger.Info("task manager started", "tasks_dir", m.tasksDir)
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	m.logger.Info("task manager stopped")
	return nil
}
