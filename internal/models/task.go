package models

import (
	"time"
)

// TaskStatus represents the current status of a task
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

// Task represents a task in the system
type Task struct {
	ID          uint       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Status      TaskStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	DueDate     time.Time  `gorm:"type:timestamp" json:"due_date"`
	CreatedAt   time.Time  `gorm:"not null;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null;default:current_timestamp" json:"updated_at"`
}

// NewTask creates a new task with default values
func NewTask(title, description string) *Task {
	now := time.Now()
	return &Task{
		Title:       title,
		Description: description,
		Status:      TaskStatusPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateStatus changes the task status and updates the UpdatedAt timestamp
func (t *Task) UpdateStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// UpdateDetails updates the task details and updates the UpdatedAt timestamp
func (t *Task) UpdateDetails(title, description string) {
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
} 