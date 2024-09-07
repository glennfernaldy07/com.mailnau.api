package errors

import (
	"errors"
	"fmt"
)

var (
	// Define common error messages
	ErrNotFound      = errors.New("resource not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternalError = errors.New("internal server error")
)

// ServiceError represents a custom error with contextual information from the service layer.
type ServiceError struct {
	Code    int
	Message string
	Err     error
}

// Implement the error interface
func (e *ServiceError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, OriginalError: %v", e.Code, e.Message, e.Err)
}

// NewServiceErrorWrapper wraps a new error with the provided context.
func NewServiceErrorWrapper(code int, message string, err error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// DeliveryError represents a custom error with contextual information from the service layer.
type DeliveryError struct {
	Code    int
	Message string
	Err     error
}

// Implement the error interface
func (e *DeliveryError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, OriginalError: %v", e.Code, e.Message, e.Err)
}

// NewDeliveryErrorWrapper wraps a new error with the provided context.
func NewDeliveryErrorWrapper(code int, message string, err error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
