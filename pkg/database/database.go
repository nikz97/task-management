package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetEnvOrDefault returns the value of the environment variable if it exists,
// otherwise returns the default value
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetConfigFromEnv returns database configuration from environment variables
func GetConfigFromEnv() *Config {
	return &Config{
		Host:     GetEnvOrDefault("DB_HOST", "localhost"),
		Port:     GetEnvOrDefault("DB_PORT", "5432"),
		User:     GetEnvOrDefault("DB_USER", "postgres"),
		Password: GetEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:   GetEnvOrDefault("DB_NAME", "taskdb"),
	}
}

// NewPostgresConnection creates a new PostgreSQL database connection using GORM
func NewPostgresConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Get the underlying SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting SQL DB: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
} 