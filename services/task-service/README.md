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
│       └── main.go           # Application entry point
├── internal/
│   └── config.go            # Configuration and error handling
├── models.go                # Domain models and interfaces
├── repository/
│   ├── task_repository.go   # Repository interface
│   └── postgres/
│       └── task_repository.go
├── service/
│   └── task_service.go      # Business logic
├── pkg/
│   └── http/
│       ├── handlers/
│       ├── middleware/
│       └── routes/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── schema.sql
``` 