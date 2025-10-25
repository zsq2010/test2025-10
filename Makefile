.PHONY: all build run clean test fmt lint help

BINARY_NAME=server
BINARY_PATH=bin/$(BINARY_NAME)
MAIN_PATH=./cmd/server

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	@go build -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_PATH)"

run:
	@echo "Running $(BINARY_NAME)..."
	@go run $(MAIN_PATH)

dev: build
	@echo "Starting $(BINARY_NAME) in development mode..."
	@./$(BINARY_PATH)

clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean
	@echo "Clean complete"

test:
	@echo "Running tests..."
	@go test -v ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, skipping..."; \
		go vet ./...; \
	fi

tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

deps:
	@echo "Downloading dependencies..."
	@go mod download

help:
	@echo "Available targets:"
	@echo "  build    - Build the server binary"
	@echo "  run      - Run the server directly with go run"
	@echo "  dev      - Build and run the server binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format code"
	@echo "  lint     - Run linter"
	@echo "  tidy     - Tidy dependencies"
	@echo "  deps     - Download dependencies"
	@echo "  help     - Show this help message"
