# Architecture Overview

## System Design

This is a Go-based backend server designed for task automation and browser automation using chromedp.

## Components

### 1. HTTP Server (`internal/api`)
- **Router**: Uses gorilla/mux for HTTP routing
- **Handlers**: RESTful API endpoints
- **Middleware**: Request logging and other cross-cutting concerns
- **Endpoints**:
  - `GET /health` - Health check endpoint

### 2. Configuration Management (`internal/config`)
- Supports multiple configuration sources:
  - Environment variables (highest priority)
  - YAML configuration files
  - Command-line flags
  - Default values (lowest priority)
- Validates configuration on load
- Type-safe configuration structs

### 3. Logging (`internal/logging`)
- Uses Go 1.21+ `log/slog` for structured logging
- Supports multiple formats:
  - JSON (default, for production)
  - Text (for development)
- Configurable log levels: debug, info, warn, error
- Key-value pair logging for easy parsing

### 4. Task Manager (`internal/tasks`)
- Manages task definitions and execution (stub)
- Loads tasks from YAML files
- Future: Task scheduling, execution, and monitoring

### 5. Engine (`internal/engine`)
- Core automation engine (stub)
- Will use chromedp for browser automation
- Future: Headless browser control, screenshot capture, web scraping

## Data Flow

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ HTTP Request
       ▼
┌─────────────┐
│   Router    │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Handlers   │
└──────┬──────┘
       │
       ▼
┌─────────────┐     ┌─────────────┐
│   Engine    │────▶│Task Manager │
└─────────────┘     └─────────────┘
       │                   │
       ▼                   ▼
┌─────────────┐     ┌─────────────┐
│  Chromedp   │     │Task Workers │
└─────────────┘     └─────────────┘
```

## Startup Sequence

1. Load configuration from env vars, files, and flags
2. Initialize structured logger
3. Create task manager and engine instances
4. Start task manager
5. Start engine
6. Create HTTP router and register handlers
7. Start HTTP server
8. Wait for shutdown signal

## Shutdown Sequence

1. Receive SIGINT or SIGTERM signal
2. Stop accepting new HTTP connections
3. Wait for active requests to complete (30s timeout)
4. Stop engine
5. Stop task manager
6. Exit cleanly

## Configuration Sources (Priority Order)

1. Environment variables (highest)
2. YAML configuration file specified via `-config` flag
3. YAML configuration file specified via `CONFIG_FILE` env var
4. Default values (lowest)

## Security Considerations

- Server binds to 0.0.0.0 by default (configurable)
- No authentication implemented yet (TODO)
- No rate limiting implemented yet (TODO)
- CORS not configured yet (TODO)

## Scalability

- Stateless design allows horizontal scaling
- Task manager can be extended for distributed task processing
- Configuration via environment variables supports containerization

## Future Enhancements

- [ ] API authentication and authorization
- [ ] Task scheduling and cron support
- [ ] Task result storage and retrieval
- [ ] WebSocket support for real-time updates
- [ ] Metrics and monitoring (Prometheus)
- [ ] Distributed tracing (OpenTelemetry)
- [ ] Database integration for task persistence
- [ ] Queue system for task distribution (Redis, RabbitMQ)
- [ ] Admin UI for task management
- [ ] Docker support
- [ ] Kubernetes deployment manifests
