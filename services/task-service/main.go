package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"task-service/internal/config"
	"task-service/pkg/http/handlers"
	"task-service/pkg/http/routes"
	"task-service/pkg/http/middleware"
	"task-service/pkg/service"
	taskRepo "task-service/pkg/repository/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Database connection
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repository
	taskRepository := taskRepo.NewTaskRepository(db)

	// Initialize service
	taskService := service.NewTaskService(taskRepository)

	// Initialize handler
	taskHandler := handlers.NewTaskHandler(taskService)

	// Setup routes with middleware
	router := routes.SetupRoutes(taskHandler)
	handler := middleware.Recovery(middleware.Logger(middleware.CORS(router)))

	// Create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: handler,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
} 