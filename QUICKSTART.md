# Quick Start Guide

Get up and running with the Go backend server in minutes.

## Prerequisites

- Go 1.23+ (Go 1.24 recommended)
- Make (optional)
- curl (for testing)

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd project

# Download dependencies
go mod download

# Build the server
make build
# or
go build -o bin/server ./cmd/server
```

## Running the Server

### Option 1: Using Make

```bash
make run
```

### Option 2: Using Go directly

```bash
go run ./cmd/server
```

### Option 3: Using the compiled binary

```bash
./bin/server
```

## Testing the Server

In a new terminal window:

```bash
# Health check
curl http://localhost:8080/health

# Expected response:
# {"status":"healthy","version":"0.1.0"}
```

## Configuration

### Using Environment Variables

```bash
SERVER_PORT=9090 LOG_LEVEL=debug go run ./cmd/server
```

### Using a Config File

```bash
# Copy the example config
cp config.example.yaml config.yaml

# Edit config.yaml as needed
# Then run:
./bin/server -config config.yaml
```

## Common Tasks

### Build

```bash
make build
```

### Clean build artifacts

```bash
make clean
```

### Format code

```bash
make fmt
```

### Run linter

```bash
make lint
```

### Run tests

```bash
make test
```

## Configuration Options

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `SERVER_HOST` | `0.0.0.0` | Server host address |
| `SERVER_PORT` | `8080` | Server port |
| `LOG_LEVEL` | `info` | Logging level (debug, info, warn, error) |
| `LOG_FORMAT` | `json` | Log format (json, text) |
| `TASKS_DIR` | `./tasks` | Task definitions directory |
| `WEB_DIR` | `./web` | Web assets directory |
| `EXAMPLES_DIR` | `./examples` | Examples directory |

## Stopping the Server

Press `Ctrl+C` to stop the server. It will shut down gracefully.

## Next Steps

- Read [README.md](README.md) for detailed documentation
- Check [ARCHITECTURE.md](ARCHITECTURE.md) to understand the design
- See [CONTRIBUTING.md](CONTRIBUTING.md) to contribute
- Explore [examples/](examples/) for sample configurations

## Troubleshooting

### Port already in use

If you see an error about the port being in use:

```bash
# Change the port
SERVER_PORT=9090 go run ./cmd/server
```

### Go not found

Make sure Go is installed and in your PATH:

```bash
go version
```

### Dependencies not found

Run:

```bash
go mod download
go mod tidy
```

## Support

- Open an issue for bugs or feature requests
- See documentation in `/docs` (coming soon)
- Check examples in `/examples`
