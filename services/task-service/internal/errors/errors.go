package errors

import "fmt"

type ErrorType string

const (
	NotFound       ErrorType = "NOT_FOUND"
	InvalidInput   ErrorType = "INVALID_INPUT"
	DatabaseError  ErrorType = "DATABASE_ERROR"
	InternalError  ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
	}
}

func NewInvalidInputError(message string) *AppError {
	return &AppError{
		Type:    InvalidInput,
		Message: message,
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Type:    DatabaseError,
		Message: "Database operation failed",
		Err:     err,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: "Internal server error",
		Err:     err,
	}
} 