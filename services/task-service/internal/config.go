package internal

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "taskdb"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// GetDSN returns the database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ErrorType represents the type of error
type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	InvalidInput ErrorType = "INVALID_INPUT"
	DatabaseError ErrorType = "DATABASE_ERROR"
	InternalError ErrorType = "INTERNAL_ERROR"
)

// AppError represents an application error
type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
		Err:     err,
	}
}

// NewInvalidInputError creates a new invalid input error
func NewInvalidInputError(message string, err error) *AppError {
	return &AppError{
		Type:    InvalidInput,
		Message: message,
		Err:     err,
	}
}

// NewDatabaseError creates a new database error
func NewDatabaseError(message string, err error) *AppError {
	return &AppError{
		Type:    DatabaseError,
		Message: message,
		Err:     err,
	}
}

// NewInternalError creates a new internal error
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: message,
		Err:     err,
	}
} 