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