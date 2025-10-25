# Go Backend Server

A Go-based backend server with configuration management, structured logging, and task execution capabilities.

## Features

- HTTP server with graceful shutdown
- Structured logging with `log/slog`
- Configuration via environment variables and YAML files
- Health check endpoint
- Task management system (stub)
- Browser automation engine with chromedp (stub)
- Modular architecture with clean separation of concerns

## Project Structure

```
.
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── api/             # HTTP handlers and routing
│   ├── config/          # Configuration management
│   ├── engine/          # Engine logic (stub)
│   ├── logging/         # Logging setup
│   └── tasks/           # Task management (stub)
├── pkg/                 # Public libraries (if needed)
├── web/                 # Web assets
├── tasks/               # Task definitions
├── examples/            # Example configurations and tasks
├── Makefile             # Build and development tasks
└── go.mod               # Go module definition
```

## Requirements

- Go 1.23 or higher (currently using Go 1.24)
- Make (optional, for using Makefile)

## Quick Start

### Using Make

```bash
# Build the server
make build

# Run the server directly
make run

# Build and run
make dev

# Format code
make fmt

# Run tests
make test

# Clean build artifacts
make clean
```

### Using Go commands

```bash
# Run the server
go run ./cmd/server

# Build the server
go build -o bin/server ./cmd/server

# Run the binary
./bin/server
```

## Configuration

The server can be configured via environment variables or a YAML configuration file.

### Environment Variables

- `SERVER_HOST` - Server host address (default: `0.0.0.0`)
- `SERVER_PORT` - Server port (default: `8080`)
- `LOG_LEVEL` - Logging level: debug, info, warn, error (default: `info`)
- `LOG_FORMAT` - Log format: json, text (default: `json`)
- `TASKS_DIR` - Directory for task definitions (default: `./tasks`)
- `WEB_DIR` - Directory for web assets (default: `./web`)
- `EXAMPLES_DIR` - Directory for examples (default: `./examples`)
- `CONFIG_FILE` - Path to YAML configuration file (optional)

### Configuration File

You can also use a YAML configuration file:

```yaml
server:
  host: 0.0.0.0
  port: 8080

logging:
  level: info
  format: json

base_paths:
  tasks_dir: ./tasks
  web_dir: ./web
  examples_dir: ./examples
```

Pass the configuration file using the `-config` flag:

```bash
./bin/server -config config.yaml
```

## API Endpoints

### Health Check

```bash
GET /health
```

Returns the health status of the server:

```json
{
  "status": "healthy",
  "version": "0.1.0"
}
```

Example:

```bash
curl http://localhost:8080/health
```

## Development

### Running the server

```bash
# With default configuration
go run ./cmd/server

# With custom port
SERVER_PORT=3000 go run ./cmd/server

# With debug logging
LOG_LEVEL=debug go run ./cmd/server

# With configuration file
go run ./cmd/server -config config.yaml
```

### Code formatting

```bash
go fmt ./...
# or
make fmt
```

### Linting

```bash
# If golangci-lint is installed
golangci-lint run ./...

# Otherwise, use go vet
go vet ./...

# or
make lint
```

### Testing

```bash
go test ./...
# or
make test
```

## Graceful Shutdown

The server handles `SIGINT` (Ctrl+C) and `SIGTERM` signals gracefully:

1. Stops accepting new connections
2. Waits for active requests to complete (up to 30 seconds)
3. Shuts down the engine and task manager
4. Exits cleanly

## Architecture

### Components

- **Config Manager**: Loads and validates configuration from files and environment variables
- **Logger**: Structured logging with `log/slog`, supports JSON and text formats
- **API Router**: HTTP request routing using `gorilla/mux`
- **Engine**: Core engine logic (stub implementation)
- **Task Manager**: Task execution and management (stub implementation)

### Dependency Injection

The main function wires up all dependencies and passes them to the components that need them. This makes the codebase testable and maintainable.

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and URL matcher
- [yaml.v3](https://gopkg.in/yaml.v3) - YAML parser
- [chromedp](https://github.com/chromedp/chromedp) - Browser automation library

## License

[Add your license here]

## Contributing

[Add contribution guidelines here]
