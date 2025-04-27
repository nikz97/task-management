# Task Management System

A microservice-based Task Management System built with Go, following clean architecture principles.

## Features

- Create, Read, Update, and Delete tasks
- Pagination support for task listing
- Filter tasks by status
- RESTful API design
- PostgreSQL database integration
- Docker containerization
- Environment-based configuration
- CORS support
- Request logging
- Error handling
- Graceful shutdown

## Project Structure

```
services/task-service/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/                       # Private application code
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── errors/
│   │   └── errors.go              # Custom error types
│   └── models/
│       └── task.go                # Domain models
├── pkg/                           # Public packages
│   ├── repository/
│   │   ├── task_repository.go     # Repository interface
│   │   └── postgres/
│   │       └── task_repository.go # PostgreSQL implementation
│   ├── service/
│   │   └── task_service.go        # Business logic
│   └── transport/
│       └── http/
│           ├── handlers/          # HTTP request handlers
│           ├── middleware/        # HTTP middleware
│           └── routes/            # Route definitions
├── Dockerfile                     # Container build
├── docker-compose.yml             # Container orchestration
├── go.mod                         # Go module file
└── schema.sql                     # Database schema
```

## API Endpoints

### Tasks

- `GET /tasks` - List tasks (with pagination and filtering)
  - Query Parameters:
    - `page` (default: 1)
    - `page_size` (default: 10)
    - `status` (optional: "pending", "in_progress", "completed")
- `GET /tasks/:id` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

## Getting Started

1. Clone the repository
2. Start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```
3. The service will be available at `http://localhost:8080`

## Environment Variables

The service can be configured using the following environment variables:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=taskdb
SERVER_PORT=8080
```

## Architecture

The system follows clean architecture principles with clear separation of concerns:

- **Domain Layer** (`internal/models/`): Contains business entities
  - Task model with status and validation

- **Repository Layer** (`pkg/repository/`): Data access layer
  - Interface definition
  - PostgreSQL implementation

- **Service Layer** (`pkg/service/`): Business logic
  - Task management
  - Validation
  - Business rules

- **Transport Layer** (`pkg/transport/`): HTTP handling
  - Request handlers
  - Route definitions
  - Middleware (CORS, Logging, Recovery)

## Testing the APIs

### 1. Create Task
```bash
curl -X POST "http://localhost:8080/tasks" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Task description",
    "status": "pending",
    "due_date": "2024-05-01T00:00:00Z"
  }'
```

### 2. Get Task
```bash
curl -X GET "http://localhost:8080/tasks/1"
```

### 3. Update Task
```bash
curl -X PUT "http://localhost:8080/tasks/1" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "in_progress",
    "title": "Updated Task",
    "description": "Updated description"
  }'
```

### 4. List Tasks
```bash
# Get all tasks
curl -X GET "http://localhost:8080/tasks"

# Get tasks with pagination
curl -X GET "http://localhost:8080/tasks?page=1&page_size=5"

# Get tasks filtered by status
curl -X GET "http://localhost:8080/tasks?status=pending"
```

### 5. Delete Task
```bash
curl -X DELETE "http://localhost:8080/tasks/1"
```

## Response Formats

### Task Object
```json
{
  "id": 1,
  "title": "Task Title",
  "description": "Task Description",
  "status": "pending",
  "due_date": "2024-05-01T00:00:00Z",
  "created_at": "2024-04-27T10:00:00Z",
  "updated_at": "2024-04-27T10:00:00Z"
}
```

### List Response
```json
{
  "data": [...],
  "page": 1,
  "page_size": 10,
  "total_items": 5
}
```

## Error Handling

The service uses custom error types for different scenarios:
- `NotFound`: Resource not found
- `InvalidInput`: Invalid request data
- `DatabaseError`: Database operation failures
- `InternalError`: Unexpected server errors

## Middleware

The service includes the following middleware:
- **CORS**: Handles cross-origin requests
- **Logger**: Logs request details and duration
- **Recovery**: Recovers from panics

## Docker Support

The service is containerized using Docker and can be run using Docker Compose:

```bash
# Build and start services
docker-compose up --build

# Stop services
docker-compose down
```

## Development

To run the service locally without Docker:

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Start PostgreSQL:
   ```bash
   docker-compose up postgres
   ```

3. Run the service:
   ```bash
   go run cmd/api/main.go
   ``` 