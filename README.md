# Poseidon2 Hashing

A Go project implementing Poseidon2 hashing algorithms.

## Project Structure

```
poseidon2-hashing/
├── cmd/                    # Command line applications
├── internal/               # Private application code
├── pkg/                    # Public library code
├── api/                    # API definitions
├── docs/                   # Documentation
├── test/                   # Additional test files
├── scripts/                # Build and deployment scripts
├── go.mod                  # Go module file
├── go.sum                  # Go module checksums
├── main.go                 # Main application entry point
├── Makefile                # Build automation
└── README.md               # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd poseidon2-hashing
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

### Building

```bash
# Build for current platform
go build -o poseidon2-hashing

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o poseidon2-hashing-linux
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
go vet ./...
```

## License

[Add your license here]
