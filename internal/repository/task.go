package repository

import (
	"task-management-system/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	GetAll(query models.PaginationQuery) ([]*models.Task, int64, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetAll(query models.PaginationQuery) ([]*models.Task, int64, error) {
	var tasks []*models.Task
	var totalItems int64

	// Build the query
	db := r.db
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}

	// Count total items
	if err := db.Model(&models.Task{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (query.Page - 1) * query.PageSize
	err := db.Offset(offset).Limit(query.PageSize).Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, totalItems, nil
}

func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
} 