.PHONY: build run test test-coverage clean swagger deps help

# Application configuration
APP_NAME=1inch-testtask

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/main.go

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	@go run cmd/main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Generate swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@which swag > /dev/null || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	@swag init -g cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test ./...
	@golangci-lint run