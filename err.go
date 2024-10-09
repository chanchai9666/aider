package aider

import "fmt"

// Define constants for error codes
const (
	ErrNotFound     = 404
	ErrUnauthorized = 401
	ErrInternal     = 500
	ErrBadRequest   = 400
)

// CustomError defines a struct for handling error code and message.
type CustomError struct {
	Code    int
	Message string
}

// Implement the Error method to satisfy the error interface.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// NewError is a constructor function to create a new CustomError.
func NewError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}
