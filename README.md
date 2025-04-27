# Task Management System

A microservice-based Task Management System built with Go, following clean architecture principles.

## Features

- Create, Read, Update, and Delete tasks
- Pagination support for task listing
- Filter tasks by status
- RESTful API design
- MySQL database integration

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   └── task.go
│   ├── repository/
│   │   └── task_repository.go
│   ├── service/
│   │   └── task_service.go
│   └── handler/
│       └── task_handler.go
├── pkg/
│   └── database/
│       └── mysql.go
├── .env
├── go.mod
└── README.md
```

## API Endpoints

- `GET /tasks` - List tasks (with pagination and filtering)
- `GET /tasks/:id` - Get a specific task
- `POST /tasks` - Create a new task
- `PUT /tasks/:id` - Update a task
- `DELETE /tasks/:id` - Delete a task

## Getting Started

1. Clone the repository
2. Copy `.env.example` to `.env` and update the database configuration
3. Run `go mod download` to install dependencies
4. Run `go run cmd/api/main.go` to start the server

## Database Setup

The application uses MySQL. Create a database and update the `.env` file with your database credentials.

## Architecture

The system follows clean architecture principles with clear separation of concerns:

- **Domain Layer**: Contains business entities and interfaces
- **Repository Layer**: Handles data persistence
- **Service Layer**: Implements business logic
- **Handler Layer**: Manages HTTP requests and responses

## Scalability

The service is designed to be horizontally scalable:
- Stateless design allows multiple instances
- Database connection pooling
- Configurable through environment variables

## Inter-Service Communication

For future microservice integration:
- REST APIs for synchronous communication
- Message queues (e.g., RabbitMQ) for asynchronous communication
- gRPC for efficient service-to-service communication 

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

# Get tasks with both pagination and status filter
curl -X GET "http://localhost:8080/tasks?page=1&page_size=5&status=in_progress" | json_pp
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

# Create task with minimal fields
curl -X POST "http://localhost:8080/tasks" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Task description"
  }' | json_pp
```

### 4. Update Task
```bash
# Update task status only
curl -X PUT "http://localhost:8080/tasks/1" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }' | json_pp

# Update multiple fields
curl -X PUT "http://localhost:8080/tasks/1" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "in_progress",
    "description": "Updated description",
    "due_date": "2025-05-01T00:00:00Z"
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