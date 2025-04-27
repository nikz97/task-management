package repository

import (
	"context"
	"task-service/internal/models"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id uint) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, page, pageSize int, status models.TaskStatus) ([]*models.Task, int64, error)
} 