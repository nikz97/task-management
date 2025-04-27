package service

import (
	"errors"
	"task-management-system/internal/models"
	"task-management-system/internal/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTask(id uint) (*models.Task, error)
	GetAllTasks(query models.PaginationQuery) (models.PaginatedResponse, error)
	UpdateTask(updates map[string]interface{}) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(task)
}

func (s *taskService) GetTask(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) GetAllTasks(query models.PaginationQuery) (models.PaginatedResponse, error) {
	// Validate pagination parameters
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}
	if query.PageSize > 100 {
		query.PageSize = 100 // Maximum page size
	}

	// Get paginated tasks
	tasks, totalItems, err := s.repo.GetAll(query)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	// Create paginated response
	return models.NewPaginatedResponse(tasks, query, totalItems), nil
}

func (s *taskService) UpdateTask(updates map[string]interface{}) error {
	id, ok := updates["id"].(uint)
	if !ok {
		return errors.New("task id is required")
	}

	// Verify task exists
	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Update(existingTask.ID, updates)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
} 