package repository

import (
	"task-management-system/api/domain/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	GetAll(query models.PaginationQuery) ([]*models.Task, int64, error)
	Update(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// buildQuery constructs the base query with filters
func (r *taskRepository) buildQuery(query models.PaginationQuery) *gorm.DB {
	db := r.db
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	return db
}

// applyPagination applies pagination parameters to the query
func (r *taskRepository) applyPagination(db *gorm.DB, query models.PaginationQuery) *gorm.DB {
	offset := (query.Page - 1) * query.PageSize
	return db.Offset(offset).Limit(query.PageSize)
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

	// Build base query with filters
	db := r.buildQuery(query)

	// Count total items
	if err := db.Model(&models.Task{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and get results
	err := r.applyPagination(db, query).Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, totalItems, nil
}

func (r *taskRepository) Update(id uint, updates map[string]interface{}) error {
	// Remove id from updates as it's not a field to update
	delete(updates, "id")
	
	// Use GORM's Updates method for partial updates
	return r.db.Model(&models.Task{}).Where("id = ?", id).Updates(updates).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
} 