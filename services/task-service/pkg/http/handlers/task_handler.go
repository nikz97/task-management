package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-service/pkg/service"
	"task-service/internal/models"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateTaskRequest struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Status      models.TaskStatus `json:"status,omitempty"`
}

type taskResponse struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Status      models.TaskStatus `json:"status"`
	DueDate     string         `json:"due_date,omitempty"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(r.Context(), req.Title, req.Description)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Status != "" {
		err = h.service.UpdateTaskStatus(r.Context(), uint(id), req.Status)
	} else if req.Title != "" || req.Description != "" {
		err = h.service.UpdateTaskDetails(r.Context(), uint(id), req.Title, req.Description)
	}

	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(r.Context(), uint(id)); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if pageSize < 1 {
		pageSize = 10
	}

	status := models.TaskStatus(r.URL.Query().Get("status"))

	tasks, total, err := h.service.ListTasks(r.Context(), page, pageSize, status)
	if err != nil {
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	response := struct {
		Data       []*models.Task `json:"data"`
		Page       int           `json:"page"`
		PageSize   int           `json:"page_size"`
		TotalItems int64         `json:"total_items"`
	}{
		Data:       tasks,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
} 