package service

import (
	"context"
	"task-service/internal/models"
	"task-service/pkg/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string) (*models.Task, error) {
	task := models.NewTask(title, description)
	err := s.repo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id uint) (*models.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, id uint, status models.TaskStatus) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.UpdateStatus(status)
	return s.repo.Update(ctx, task)
}

func (s *TaskService) UpdateTaskDetails(ctx context.Context, id uint, title, description string) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.UpdateDetails(title, description)
	return s.repo.Update(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context, page, pageSize int, status models.TaskStatus) ([]*models.Task, int64, error) {
	return s.repo.List(ctx, page, pageSize, status)
} 