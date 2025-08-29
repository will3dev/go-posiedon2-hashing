# Makefile for Poseidon2 Hashing Project

.PHONY: help build run test clean fmt vet lint deps

# Default target
help:
	@echo "Available commands:"
	@echo "  build   - Build the application"
	@echo "  run     - Run the application"
	@echo "  test    - Run tests"
	@echo "  clean   - Clean build artifacts"
	@echo "  fmt     - Format code"
	@echo "  vet     - Run go vet"
	@echo "  lint    - Run linting tools"
	@echo "  deps    - Download dependencies"

# Build the application
build:
	@echo "Building..."
	go build -o poseidon2-hashing main.go

# Run tests
test: 
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts 
clean:
	@echo "Cleaning..."
	rm -f poseidon2-hashing
	go clean

# Format code
fmt:
	@echo "formatting code..."
	go fmt ./...

#Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

#Download dependencies 
deps:
	@echo "Downloading dependencies..."
	go mod tidy
	go mod download

