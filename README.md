# Task Management System

A microservice-based Task Management System built with Go, following clean architecture principles.

## Architecture

The system consists of three microservices:
- **Task Service**: Manages task CRUD operations
- **User Service**: Handles user authentication and management
- **Notification Service**: Manages task notifications and alerts

## Features

- Create, Read, Update, and Delete tasks
- Pagination support for task listing
- Filter tasks by status
- RESTful API design
- PostgreSQL database integration
- Docker containerization with horizontal scaling
- Nginx load balancing
- Environment-based configuration
- CORS support
- Request logging
- Error handling
- Graceful shutdown

## Project Structure

```
.
├── services/                    # Microservices
│   ├── task-service/           # Task Management Service
│   │   ├── cmd/
│   │   │   └── api/
│   │   │       └── main.go    # Application entry point
│   │   ├── internal/          # Private application code
│   │   │   ├── config/
│   │   │   │   └── config.go # Configuration management
│   │   │   ├── errors/
│   │   │   │   └── errors.go # Custom error types
│   │   │   └── models/
│   │   │       └── task.go   # Domain models
│   │   ├── pkg/              # Public packages
│   │   │   ├── repository/
│   │   │   │   ├── task_repository.go
│   │   │   │   └── postgres/
│   │   │   │       └── task_repository.go
│   │   │   ├── service/
│   │   │   │   └── task_service.go
│   │   │   └── transport/
│   │   │       └── http/
│   │   │           ├── handlers/
│   │   │           ├── middleware/
│   │   │           └── routes/
│   │   ├── Dockerfile
│   │   ├── docker-compose.yml
│   │   ├── nginx.conf
│   │   ├── go.mod
│   │   └── schema.sql
│   ├── user-service/          # User Management Service
│   └── notification-service/  # Notification Service
├── .env                       # Environment variables
├── .env.example              # Example environment variables
├── docker-compose.yml        # Root docker-compose for all services
├── go.mod                    # Root go module
└── go.sum                    # Go module checksums
```

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Start the services using Docker Compose:
   ```bash
   # Start all services
   docker-compose up --build

   # Or start individual services with scaling
   cd services/task-service
   docker-compose up --build --scale task-service=3
   ```
4. The services will be available at:
   - Task Service: `http://localhost:8080`
   - User Service: `http://localhost:8081`
   - Notification Service: `http://localhost:8082`

## Environment Variables

Each service can be configured using environment variables. See `.env.example` for the complete list:

```env
# Database Configuration
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=taskdb

# Service Ports
TASK_SERVICE_PORT=8080
USER_SERVICE_PORT=8081
NOTIFICATION_SERVICE_PORT=8082
```

## Architecture

The system follows clean architecture principles with clear separation of concerns:

### Task Service
- **Domain Layer** (`internal/models/`): Task entities and validation
- **Repository Layer** (`pkg/repository/`): Data access and PostgreSQL implementation
- **Service Layer** (`pkg/service/`): Business logic and task management
- **Transport Layer** (`pkg/transport/`): HTTP handlers, middleware, and routes


## Horizontal Scaling

The services are designed for horizontal scaling:
- Multiple instances of each service can run simultaneously
- Nginx load balancers distribute traffic across instances
- Shared PostgreSQL databases ensure data consistency
- Stateless design allows for easy scaling

To scale a service:
```bash
cd services/<service-name>
docker-compose up --build --scale <service-name>=3
```

## API Documentation

Each service has its own API documentation. See the respective service's README.md for detailed API endpoints.

### Task Service Endpoints
- `GET /tasks` - List tasks (with pagination and filtering)
- `GET /tasks/:id` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

## Error Handling

Each service implements custom error types for different scenarios:
- `NotFound`: Resource not found
- `InvalidInput`: Invalid request data
- `DatabaseError`: Database operation failures
- `InternalError`: Unexpected server errors

## Middleware

Common middleware across services:
- **CORS**: Handles cross-origin requests
- **Logger**: Logs request details and duration
- **Recovery**: Recovers from panics
- **Auth**: JWT token validation (User Service)

## Docker Support

The services are containerized using Docker and can be run using Docker Compose:

```bash
# Build and start all services
docker-compose up --build

# Stop all services
docker-compose down
```

## Development

To run services locally without Docker:

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Start PostgreSQL:
   ```bash
   docker-compose up postgres
   ```

3. Run a service:
   ```bash
   cd services/<service-name>
   go run cmd/api/main.go
   ``` 