# AGENTS.md - Development Guide for Agentic Coding Agents

This document provides essential information for agentic coding agents working with this Go backend repository. It covers build commands, testing procedures, and code style guidelines.

## Build and Run Commands

### Basic Build
```bash
go build .
```

### Run the Application
```bash
# Run the API server
go run . serve

# Load fixtures to database
go run . load

# Run database migrations
go run . migrate

# Index events (continuous process)
go run . index

# Analyze UniswapV3 data
go run . analyze
```

### Start Dependencies
```bash
# Start PostgreSQL database
./startdb.sh

# Start Redis cache
./startredis.sh

# Start both database and Redis
./startall.sh
```

## Testing Commands

### Running Tests
Since this codebase currently doesn't have specific test files, you can run tests with:
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./trade/api

# Run a single test function (if it existed)
go test -run TestFunctionName ./package/path
```

### Test Structure
When adding tests, follow this pattern:
- Place test files in the same directory as the code they test
- Name test files with `_test.go` suffix
- Use table-driven tests when appropriate
- Follow the existing code style for tests

## Linting and Formatting

### Go Formatting
```bash
# Format all Go files
go fmt ./...

# Format specific file
go fmt trade/api/api.go
```

### Linting with golangci-lint
```bash
# Install golangci-lint if not already installed
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linting on entire codebase
golangci-lint run

# Run linting on specific directory
golangci-lint run trade/api/...
```

## Code Style Guidelines

### General Principles
1. Follow the official Go Code Review Comments guide
2. Write clean, readable, and maintainable code
3. Prefer explicit over implicit code
4. Handle errors appropriately - don't ignore them
5. Use context.Context for cancellation and timeouts
6. Write meaningful comments for exported functions and types

### Imports
1. Use grouped imports with blank lines between groups:
   - Standard library
   - Third-party packages
   - Local packages
2. Avoid import aliases unless there are conflicts
3. Remove unused imports

### Naming Conventions
1. Use camelCase for variables and functions
2. Use PascalCase for exported names
3. Use short variable names for short scopes (e.g., `i`, `j`)
4. Use descriptive names for longer scopes
5. Use meaningful names that convey purpose

### Types
1. Use structs to group related data
2. Define interfaces with only required methods
3. Prefer pointer receivers for methods that modify the receiver
4. Use value receivers for small types and when methods don't modify the receiver

### Error Handling
1. Always handle errors returned by functions
2. Don't use `_` to ignore errors unless explicitly intended
3. Wrap errors with context when passing them up the stack
4. Use `fmt.Errorf("operation: %w", err)` to wrap errors
5. Define custom error types for specific error conditions

### Functions
1. Keep functions small and focused on a single task
2. Return early to reduce nesting
3. Use meaningful parameter names
4. Document exported functions with comments
5. Prefer returning values over modifying parameters

### Comments
1. Comment all exported functions, types, and variables
2. Use full sentences starting with the name of the thing being documented
3. Focus on why something is done, not what is done
4. Keep comments up to date with code changes

### Logging
1. Use the `log/slog` package for structured logging
2. Include relevant context in log messages
3. Use appropriate log levels (debug, info, warn, error)
4. Avoid logging sensitive information

### Database Operations
1. Use GORM for database operations
2. Handle database errors appropriately
3. Use transactions when performing multiple related operations
4. Use connection pooling appropriately

### API Design
1. Follow RESTful conventions when possible
2. Use appropriate HTTP status codes
3. Return JSON responses with consistent structure
4. Validate input parameters
5. Handle API errors gracefully with meaningful messages

### Concurrency
1. Use goroutines and channels appropriately
2. Avoid data races with proper synchronization
3. Use context for cancellation
4. Limit the number of concurrent operations when necessary

### Dependencies
1. Keep dependencies minimal and well-maintained
2. Use go.mod and go.sum for dependency management
3. Regularly update dependencies
4. Pin specific versions for stability

## Git Workflow

### Commit Messages
1. Use clear, concise commit messages
2. Start with a capital letter
3. Use imperative mood ("Add feature" not "Added feature")
4. Keep first line under 50 characters
5. Use body for detailed explanations when necessary

### Branching Strategy
1. Use feature branches for new development
2. Create pull requests for code review
3. Keep branches up to date with main branch
4. Delete branches after merging

## Development Environment

### Required Tools
1. Go 1.24 or later
2. Docker (for running PostgreSQL and Redis)
3. PostgreSQL database
4. Redis server

### Environment Setup
1. Run `./startall.sh` to start required services
2. Ensure database connection settings in main.go are correct
3. Install dependencies with `go mod tidy`

### Debugging
1. Use `go run` with appropriate flags for debugging
2. Add logging statements to trace execution
3. Use breakpoints in your IDE when supported

## Common Patterns in This Codebase

### API Handlers
Handlers in `trade/api/api.go` follow a pattern:
1. Extract parameters from context
2. Call business logic functions
3. Handle errors with `apiErr` function
4. Return JSON responses

### Database Models
Models are defined in `trade/model.go` and follow GORM conventions:
1. Use struct tags for database mapping
2. Define relationships with appropriate associations
3. Use pointers for optional fields

### Workers
Worker functions in `trade/worker/` perform background tasks:
1. Use cache manager for efficient data access
2. Implement retry logic for transient failures
3. Log important operations and errors

## Adding New Features

### Protocol Integration
To add support for a new DeFi protocol:
1. Create a new handler in `trade/protocols/`
2. Implement the `DeFiProtocolHandler` interface
3. Add necessary smart contract interfaces
4. Update the worker to process events from the new protocol

### API Endpoints
To add new API endpoints:
1. Add handler functions in `trade/api/api.go`
2. Register routes in `CreateApi` function
3. Ensure proper error handling and validation
4. Update documentation if necessary

This guide should help agentic coding agents effectively contribute to this codebase while maintaining consistency and quality.