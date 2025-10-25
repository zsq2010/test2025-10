# Contributing Guidelines

Thank you for considering contributing to this project!

## Development Setup

1. **Prerequisites**
   - Go 1.23 or higher
   - Git
   - Make (optional but recommended)

2. **Clone and Setup**
   ```bash
   git clone <repository-url>
   cd project
   go mod download
   ```

3. **Build and Run**
   ```bash
   make build
   make run
   ```

## Code Style

### Formatting
- All code must be formatted with `gofmt`
- Run `make fmt` before committing
- Use standard Go naming conventions

### Code Organization
- Keep business logic in `/internal` packages
- Use dependency injection for testability
- Pass `context.Context` for cancellable operations
- Use structured logging with key-value pairs

### Logging
```go
logger.Info("operation completed",
    "duration", elapsed,
    "status", "success",
)
```

### Error Handling
```go
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

## Testing

- Write unit tests for new functionality
- Run tests with `make test` or `go test ./...`
- Aim for meaningful test coverage
- Use table-driven tests where appropriate

## Pull Request Process

1. **Create a Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Changes**
   - Write clean, documented code
   - Add tests for new functionality
   - Update documentation as needed

3. **Verify**
   ```bash
   make fmt
   make lint
   make test
   make build
   ```

4. **Commit**
   - Use clear, descriptive commit messages
   - Reference issue numbers if applicable

5. **Submit PR**
   - Provide a clear description of changes
   - Link related issues
   - Ensure all checks pass

## Project Structure

- `/cmd/server` - Main application entry point
- `/internal/api` - HTTP handlers and routing
- `/internal/config` - Configuration management
- `/internal/engine` - Core engine logic
- `/internal/tasks` - Task management
- `/internal/logging` - Logging setup
- `/pkg` - Public packages (import safe)
- `/web` - Web assets
- `/tasks` - Task definitions
- `/examples` - Example configurations

## Adding New Features

### New API Endpoint
1. Add handler in `/internal/api/handlers.go`
2. Register route in `/internal/api/router.go`
3. Add tests
4. Update API documentation

### New Configuration Option
1. Add field to config structs in `/internal/config/config.go`
2. Add environment variable handling
3. Update validation
4. Document in README.md

### New Task Type
1. Define task structure
2. Implement task execution logic
3. Add to task manager
4. Create example in `/examples`

## Dependencies

- Use `go get` to add new dependencies
- Run `go mod tidy` to clean up
- Justify new dependencies in PR description

## Questions?

Feel free to open an issue for discussion before starting significant work.
