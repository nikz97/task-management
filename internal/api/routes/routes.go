package routes

import (
	"task-management-system/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskHandler *handlers.TaskHandler) *gin.Engine {
	router := gin.Default()

	// Task routes
	tasks := router.Group("/api/tasks")
	{
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTask)
		tasks.GET("/", taskHandler.GetAllTasks)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}

	return router
} 