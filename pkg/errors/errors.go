package errors

import (
	"fmt"
	"net/http"
)

// AppError represents a structured application error
type AppError struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	StatusCode int         `json:"-"`
	Details    interface{} `json:"details,omitempty"`
	Timestamp  int64       `json:"timestamp"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code string, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Timestamp:  getTimestamp(),
	}
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details interface{}) *AppError {
	e.Details = details
	return e
}

// Common error constructors
func BadRequest(message string) *AppError {
	return NewAppError("BAD_REQUEST", message, http.StatusBadRequest)
}

func Unauthorized(message string) *AppError {
	return NewAppError("UNAUTHORIZED", message, http.StatusUnauthorized)
}

func Forbidden(message string) *AppError {
	return NewAppError("FORBIDDEN", message, http.StatusForbidden)
}

func NotFound(message string) *AppError {
	return NewAppError("NOT_FOUND", message, http.StatusNotFound)
}

func Conflict(message string) *AppError {
	return NewAppError("CONFLICT", message, http.StatusConflict)
}

func InternalError(message string) *AppError {
	return NewAppError("INTERNAL_ERROR", message, http.StatusInternalServerError)
}

func ValidationError(field string, reason string) *AppError {
	err := BadRequest(fmt.Sprintf("validation error in field '%s': %s", field, reason))
	err.Details = map[string]string{
		"field":  field,
		"reason": reason,
	}
	return err
}

func DatabaseError(operation string, err error) *AppError {
	appErr := InternalError(fmt.Sprintf("database error during %s", operation))
	appErr.WithDetails(map[string]string{
		"operation": operation,
		"error":     err.Error(),
	})
	return appErr
}

func ServiceError(service string, operation string, err error) *AppError {
	appErr := InternalError(fmt.Sprintf("service error in %s.%s", service, operation))
	appErr.WithDetails(map[string]string{
		"service":   service,
		"operation": operation,
		"error":     err.Error(),
	})
	return appErr
}

func DuplicateEntry(entity string, field string, value string) *AppError {
	err := Conflict(fmt.Sprintf("%s with %s '%s' already exists", entity, field, value))
	err.WithDetails(map[string]string{
		"entity": entity,
		"field":  field,
		"value":  value,
	})
	return err
}

func MissingRequiredField(field string) *AppError {
	return ValidationError(field, "required field is missing")
}

func InvalidFieldFormat(field string, format string) *AppError {
	return ValidationError(field, fmt.Sprintf("invalid format, expected %s", format))
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) (*AppError, bool) {
	if err == nil {
		return nil, false
	}
	appErr, ok := err.(*AppError)
	return appErr, ok
}

// getTimestamp returns current Unix timestamp in milliseconds
func getTimestamp() int64 {
	return timeNowFunc()
}

// timeNowFunc can be mocked in tests
var timeNowFunc = func() int64 {
	return 0 // This will be set at runtime
}

func init() {
	timeNowFunc = func() int64 {
		return 0 // Will be overridden at runtime
	}
}
