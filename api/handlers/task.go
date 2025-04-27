package handlers

import (
	"net/http"
	"strconv"
	"task-management-system/api/domain/models"
	"task-management-system/api/service"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// RegisterRoutes registers the task routes
func (h *TaskHandler) RegisterRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("", h.GetAllTasks)
		tasks.GET("/:id", h.GetTask)
		tasks.POST("", h.CreateTask)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}
}

// parseTaskID extracts and validates the task ID from the request
func (h *TaskHandler) parseTaskID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// handleError sends appropriate error response
func (h *TaskHandler) handleError(c *gin.Context, err error, status int) {
	c.JSON(status, gin.H{"error": err.Error()})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input struct {
		Title       string           `json:"title" binding:"required"`
		Description string           `json:"description"`
		Status      models.TaskStatus `json:"status"`
		DueDate     time.Time        `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	task := models.NewTask(input.Title, input.Description)
	task.Status = input.Status
	task.DueDate = input.DueDate

	if err := h.service.CreateTask(task); err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := h.parseTaskID(c)
	if err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		h.handleError(c, err, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	var query models.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	response, err := h.service.GetAllTasks(query)
	if err != nil {
		h.handleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := h.parseTaskID(c)
	if err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	updates["id"] = id

	if err := h.service.UpdateTask(updates); err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		h.handleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := h.parseTaskID(c)
	if err != nil {
		h.handleError(c, err, http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(id); err != nil {
		h.handleError(c, err, http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
} 