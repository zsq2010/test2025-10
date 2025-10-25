package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Logging   LoggingConfig   `yaml:"logging"`
	BasePaths BasePathsConfig `yaml:"base_paths"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type BasePathsConfig struct {
	TasksDir    string `yaml:"tasks_dir"`
	WebDir      string `yaml:"web_dir"`
	ExamplesDir string `yaml:"examples_dir"`
}

func Load() (*Config, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to configuration file")
	flag.Parse()

	cfg := &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvAsInt("SERVER_PORT", 8080),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		BasePaths: BasePathsConfig{
			TasksDir:    getEnv("TASKS_DIR", "./tasks"),
			WebDir:      getEnv("WEB_DIR", "./web"),
			ExamplesDir: getEnv("EXAMPLES_DIR", "./examples"),
		},
	}

	if configPath != "" {
		if err := loadConfigFile(configPath, cfg); err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	} else if defaultPath := getEnv("CONFIG_FILE", ""); defaultPath != "" {
		if err := loadConfigFile(defaultPath, cfg); err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

func loadConfigFile(path string, cfg *Config) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve config path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

func (c *Config) validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLevels[c.Logging.Level] {
		return fmt.Errorf("invalid log level: %s", c.Logging.Level)
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
