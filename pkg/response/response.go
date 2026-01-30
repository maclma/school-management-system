package response

import (
	"net/http"
	"school-management-system/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	RequestID string      `json:"request_id,omitempty"`
}

// Success sends a success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC(),
		RequestID: getRequestID(c),
	})
}

// Created sends a 201 Created response
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC(),
		RequestID: getRequestID(c),
	})
}

// Error sends an error response
func Error(c *gin.Context, err error) {
	appErr, isAppError := errors.IsAppError(err)

	if !isAppError {
		// Wrap non-app errors
		appErr = errors.InternalError(err.Error())
	}

	c.JSON(appErr.StatusCode, Response{
		Success:   false,
		Message:   appErr.Message,
		Error:     appErr,
		Timestamp: time.Now().UTC(),
		RequestID: getRequestID(c),
	})
}

// Paginated sends a paginated response
func Paginated(c *gin.Context, message string, data interface{}, page int, limit int, total int64) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	paginationData := map[string]interface{}{
		"items":       data,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	}

	c.JSON(http.StatusOK, Response{
		Success:   true,
		Message:   message,
		Data:      paginationData,
		Timestamp: time.Now().UTC(),
		RequestID: getRequestID(c),
	})
}

// NoContent sends a 204 No Content response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// BadRequest sends a 400 Bad Request error
func BadRequest(c *gin.Context, message string) {
	Error(c, errors.BadRequest(message))
}

// Unauthorized sends a 401 Unauthorized error
func Unauthorized(c *gin.Context, message string) {
	Error(c, errors.Unauthorized(message))
}

// Forbidden sends a 403 Forbidden error
func Forbidden(c *gin.Context, message string) {
	Error(c, errors.Forbidden(message))
}

// NotFound sends a 404 Not Found error
func NotFound(c *gin.Context, message string) {
	Error(c, errors.NotFound(message))
}

// Conflict sends a 409 Conflict error
func Conflict(c *gin.Context, message string) {
	Error(c, errors.Conflict(message))
}

// InternalError sends a 500 Internal Server Error
func InternalError(c *gin.Context, message string) {
	Error(c, errors.InternalError(message))
}

// getRequestID retrieves request ID from context
func getRequestID(c *gin.Context) string {
	if id, exists := c.Get("request_id"); exists {
		return id.(string)
	}
	return ""
}
