package main

import (
	"log"
	"task-management-system/api/handlers"
	"task-management-system/api/domain/models"
	"task-management-system/api/domain/repository"
	"task-management-system/api/service"
	"task-management-system/api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get database configuration from environment variables
	dbConfig := database.GetConfigFromEnv()

	// Initialize database connection
	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize repository
	taskRepo := repository.NewTaskRepository(db)

	// Initialize service
	taskService := service.NewTaskService(taskRepo)

	// Initialize handler
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup router
	router := gin.Default()
	taskHandler.RegisterRoutes(router)

	// Get port from environment variable or use default
	port := database.GetEnvOrDefault("PORT", "8080")

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
} 