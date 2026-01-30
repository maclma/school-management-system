package middleware

import (
	"net/http"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"

	"github.com/gin-gonic/gin"
)

// ValidationMiddleware validates request content type
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
			contentType := c.ContentType()
			if contentType != "application/json" && contentType != "" {
				response.Error(c, errors.BadRequest("invalid content-type, expected application/json"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// MaxBodySizeMiddleware limits request body size
func MaxBodySizeMiddleware(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}

// BindJSONWithError binds JSON and returns proper error response
func BindJSONWithError(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return false
	}
	return true
}

// BindQueryWithError binds query parameters and returns proper error response
func BindQueryWithError(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindQuery(obj); err != nil {
		response.BadRequest(c, "invalid query parameters: "+err.Error())
		return false
	}
	return true
}

// BindURIWithError binds URI parameters and returns proper error response
func BindURIWithError(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindUri(obj); err != nil {
		response.BadRequest(c, "invalid URI parameters: "+err.Error())
		return false
	}
	return true
}
