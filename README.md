# Gin Template

A Golang web application template based on the Gin framework with a clean architecture.

## Project Structure

```
├── cmd/
│   ├── main.go          # Application entry point
│   ├── wire_gen.go      # Generated dependency injection
│   └── wire.go          # Dependency injection configuration
├── configs/
│   └── config.yaml      # Application configuration
└── internal/
    ├── base/            # Base components
    │   ├── conf/        # Configuration handling
    │   ├── data/        # Data layer components
    │   └── server/      # Server implementation
    ├── controller/      # HTTP controllers
    ├── entity/          # Domain models
    ├── repo/            # Repositories
    ├── router/          # HTTP routing
    └── service/         # Business logic
```

## Features

- Clean architecture design
- Dependency injection using [Wire](https://github.com/google/wire)
- Configuration management
- HTTP server using [Gin framework](https://github.com/gin-gonic/gin)

## Prerequisites

- Go 1.18 or higher
- Git

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/gin-template.git
cd gin-template

# Install dependencies
go mod download
```

## Configuration

Application configuration is stored in `configs/config.yaml`. You can modify it according to your requirements:

```yaml
# Example configuration, update with your actual configuration
server:
  http:
    addr: :8080
```

## Running the Application

```bash
# Build and run
go run ./cmd/main.go

# Or build executable
go build -o app ./cmd
./app
```

## Development

### Adding new endpoints

1. Define entities in `internal/entity`
2. Implement data access in `internal/repo`
3. Implement business logic in `internal/service`
4. Create controllers in `internal/controller`
5. Add routes in `internal/router`
6. Register routes in `internal/base/server/http`, just like

```go
func NewHTTPServer(debug bool, apiRouter *router.APIRouter) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })

	rootGroup := r.Group("")
	apiRouter.RegisterRouter(rootGroup)

	return r
}
```

### Dependency Injection

This project uses [Wire](https://github.com/google/wire) for dependency injection:

```bash
# Generate wire_gen.go after adding or modifying providers
wire ./cmd
```

## License

[MIT License](LICENSE)