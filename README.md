# Go Server API

A well-structured REST API built with Go using clean architecture principles.

## Project Structure

```
go_server/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/                    # Private application code
│   ├── config/                  # Configuration management
│   ├── handlers/                # HTTP handlers (controllers)
│   ├── middleware/              # HTTP middleware
│   ├── models/                  # Data models
│   ├── repository/              # Database layer
│   ├── router/                  # Route definitions
│   └── service/                 # Business logic layer
├── pkg/                         # Public libraries
│   └── utils/                   # Utility functions
├── .env.example                 # Environment variables template
├── .gitignore                   # Git ignore file
├── go.mod                       # Go module file
└── README.md                    # This file
```

## Architecture Layers

1. **Handler Layer** (`internal/handlers/`) - Handles HTTP requests and responses
2. **Service Layer** (`internal/service/`) - Contains business logic
3. **Repository Layer** (`internal/repository/`) - Handles data access

## Getting Started

### Prerequisites

- Go 1.21 or higher
- (Optional) PostgreSQL for database

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Copy the environment file:
```bash
cp .env.example .env
```

4. Update the `.env` file with your configuration

### Running the Server

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Building

```bash
go build -o bin/server cmd/api/main.go
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check API health status

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/{id}` - Get a specific user
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/{id}` - Update a user
- `DELETE /api/v1/users/{id}` - Delete a user

### Protected Routes (requires authentication)
- `GET /api/v1/protected/profile` - Get user profile

## Development

### Adding a New Feature

1. Create models in `internal/models/`
2. Implement repository methods in `internal/repository/`
3. Implement business logic in `internal/service/`
4. Create handlers in `internal/handlers/`
5. Add routes in `internal/router/router.go`

### Testing

```bash
go test ./...
```

## TODO

- [ ] Add database integration (PostgreSQL/MySQL)
- [ ] Implement JWT authentication
- [ ] Add input validation
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add API documentation (Swagger/OpenAPI)
- [ ] Add logging framework
- [ ] Add error handling improvements
- [ ] Add database migrations
- [ ] Add Docker support

## License

MIT
