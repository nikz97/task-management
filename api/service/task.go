package service

import (
	"errors"
	"task-management-system/api/domain/models"
	"task-management-system/api/domain/repository"
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

// validateTask validates task creation requirements
func (s *taskService) validateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

// validatePagination normalizes pagination parameters
func (s *taskService) validatePagination(query *models.PaginationQuery) {
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}
	if query.PageSize > 100 {
		query.PageSize = 100 // Maximum page size
	}
}

func (s *taskService) CreateTask(task *models.Task) error {
	if err := s.validateTask(task); err != nil {
		return err
	}
	return s.repo.Create(task)
}

func (s *taskService) GetTask(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) GetAllTasks(query models.PaginationQuery) (models.PaginatedResponse, error) {
	s.validatePagination(&query)

	tasks, totalItems, err := s.repo.GetAll(query)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	return models.NewPaginatedResponse(tasks, query, totalItems), nil
}

func (s *taskService) UpdateTask(updates map[string]interface{}) error {
	id, ok := updates["id"].(uint)
	if !ok {
		return errors.New("task id is required")
	}

	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Update(existingTask.ID, updates)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
} 