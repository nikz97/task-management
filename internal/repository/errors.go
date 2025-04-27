package repository

import "errors"

var (
	// ErrTaskNotFound is returned when a requested task does not exist
	ErrTaskNotFound = errors.New("task not found")
) 