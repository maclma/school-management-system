.PHONY: build run test clean migrate setup dev lint format

# Build the application
build:
	@echo "Building..."
	go build -o bin/server ./cmd/server

# Run the application
run: build
	@echo "Starting server..."
	./bin/server

# Run in development mode with live reload (requires air)
dev:
	@if command -v air >/dev/null 2>&1; then \
		air -c .air.toml; \
	else \
		echo "Air not installed. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air -c .air.toml; \
	fi

# Run tests
test:
	@echo "Running tests..."
	go test ./... -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf logs/*.log
	rm -rf coverage.out

# Setup the project
setup:
	@chmod +x scripts/setup.sh
	./scripts/setup.sh

# Run migrations
migrate:
	@echo "Running migrations..."
	go run ./cmd/migrate

# Lint the code
lint:
	@echo "Linting..."
	golangci-lint run

# Format the code
format:
	@echo "Formatting..."
	gofmt -w .

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cosmtrek/air@latest

# Create admin user
create-admin:
	@echo "Creating admin user..."
	go run scripts/create_admin.go

help:
	@echo "Available commands:"
	@echo "  make setup    - Setup the project (database, dependencies)"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make dev      - Run with live reload (requires air)"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linter"
	@echo "  make format   - Format code"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make deps     - Install dependencies"