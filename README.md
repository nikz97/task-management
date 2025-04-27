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

## Project Structure

```
.
├── api/
│   ├── config/         # Configuration management
│   ├── database/       # Database connection and configuration
│   ├── domain/         # Core business logic
│   │   ├── models/     # Domain models
│   │   └── repository/ # Data access interfaces and implementations
│   ├── handlers/       # HTTP request handlers
│   ├── routes/         # Route definitions
│   └── service/        # Business logic layer
├── tests/              # Test files
│   ├── handlers/       # Handler tests
│   ├── service/        # Service tests
│   ├── repository/     # Repository tests
│   └── database/       # Database tests
├── .env                # Environment variables
├── .env.example        # Example environment variables
├── docker-compose.yml  # Docker compose configuration
├── Dockerfile         # Docker build configuration
├── go.mod             # Go module file
├── go.sum             # Go module checksum
├── main.go            # Application entry point
├── README.md          # Project documentation
└── schema.sql         # Database schema
```

## API Endpoints

- `GET /tasks` - List tasks (with pagination and filtering)
- `GET /tasks/:id` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and update the configuration:
   ```bash
   cp .env.example .env
   ```
3. Start the PostgreSQL database using Docker:
   ```bash
   docker-compose up -d db
   ```
4. Run `go mod download` to install dependencies
5. Run `go run main.go` to start the server

## Database Setup

The application uses PostgreSQL. The database configuration can be set through environment variables in the `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=taskdb
```

## Architecture

The system follows clean architecture principles with clear separation of concerns:

- **Domain Layer** (`api/domain/`): Contains business entities and interfaces
  - Models: Core business entities
  - Repository: Data access interfaces

- **Service Layer** (`api/service/`): Implements business logic
  - Task management
  - Validation
  - Business rules

- **Handler Layer** (`api/handlers/`): Manages HTTP requests and responses
  - Request parsing
  - Response formatting
  - Error handling

- **Database Layer** (`api/database/`): Database connection and configuration
  - Connection management
  - Configuration
  - Environment variables

## Testing

The project includes comprehensive test coverage:

- Unit tests for all layers
- Integration tests for API endpoints
- Database tests
- Mock implementations for testing

Run tests with:
```bash
go test ./tests/...
```

## Docker Support

The application can be run using Docker:

1. Build the image:
   ```bash
   docker build -t task-management .
   ```

2. Run with Docker Compose:
   ```bash
   docker-compose up
   ```

## Testing the APIs

You can test the APIs using the following curl commands:

### 1. List Tasks
```bash
# Get all tasks (default pagination)
curl -X GET "http://localhost:8080/tasks" | json_pp

# Get tasks with custom pagination
curl -X GET "http://localhost:8080/tasks?page=1&page_size=5" | json_pp

# Get tasks filtered by status
curl -X GET "http://localhost:8080/tasks?status=pending" | json_pp
```

### 2. Get Single Task
```bash
# Get task by ID
curl -X GET "http://localhost:8080/tasks/1" | json_pp
```

### 3. Create Task
```bash
# Create task with all fields
curl -X POST "http://localhost:8080/tasks" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Task description",
    "status": "pending",
    "due_date": "2025-04-30T00:00:00Z"
  }' | json_pp
```

### 4. Update Task
```bash
# Update task status
curl -X PUT "http://localhost:8080/tasks/1" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }' | json_pp
```

### 5. Delete Task
```bash
# Delete task by ID
curl -X DELETE "http://localhost:8080/tasks/1" -v
```

### Available Task Status Values
- `pending`
- `in_progress`
- `completed`

### Response Formats

1. **List Response (Paginated):**
```json
{
  "data": [...],
  "page": 1,
  "page_size": 10,
  "total_items": 6,
  "total_pages": 1
}
```

2. **Single Task Response:**
```json
{
  "id": 1,
  "title": "Task Title",
  "description": "Task Description",
  "status": "pending",
  "due_date": "2025-04-30T00:00:00Z",
  "created_at": "2025-04-27T10:56:43.125769Z",
  "updated_at": "2025-04-27T10:56:43.125769Z"
}
```

### HTTP Status Codes
- 200: Success (GET, PUT)
- 201: Created (POST)
- 204: No Content (DELETE)
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error 