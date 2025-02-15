.PHONY: all build run test test-unit test-e2e clean lint coverage help

# Default target
all: build

# Build the application
build:
	@echo "Building..."
	@go build -o bin/app main.go

# Run the application
run:
	@echo "Starting application..."
	@go run main.go

# Run all tests
test: test-unit test-e2e
	@echo "All tests completed"

# Run unit tests
test-unit:
	@echo "Running unit tests..."
	@go test ./... -v -count=1

# Run unit tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

# Run e2e tests
test-e2e:
	@echo "Running E2E tests..."
	@go test ./tests/e2e/... -tags=e2e -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Help command
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run all tests"
	@echo "  make test-unit  - Run unit tests"
	@echo "  make test-e2e   - Run E2E tests"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make lint       - Run linter"