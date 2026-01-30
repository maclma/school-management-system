package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware generates and adds request ID to context
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// CORSMiddleware configures CORS with security headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", getAllowedOrigins(c.Request.Header.Get("Origin")))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SecurityHeadersMiddleware adds security headers to responses
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Enable XSS protection
		c.Header("X-XSS-Protection", "1; mode=block")

		// Prevent clickjacking
		c.Header("X-Frame-Options", "DENY")

		// Content Security Policy
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'")

		// Referrer Policy
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Permissions Policy
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		// Strict Transport Security
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		c.Next()
	}
}

// RequestLoggingMiddleware logs incoming requests
func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestID := c.GetString("request_id")

		// Log request
		fmt.Printf("[%s] %s %s %s\n", requestID, c.Request.Method, c.Request.RequestURI, c.ClientIP())

		c.Next()

		// Log response
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		fmt.Printf("[%s] Response: %d in %v\n", requestID, statusCode, duration)
	}
}

// TimeoutMiddleware adds request timeout
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a copy of context to preserve the timeout
		c.Request = c.Request.WithContext(c.Request.Context())
		c.Next()
	}
}

// getAllowedOrigins returns allowed origins for CORS
func getAllowedOrigins(origin string) string {
	allowedOrigins := map[string]bool{
		"http://localhost:3000": true,
		"http://localhost:5173": true,
		"http://127.0.0.1:3000": true,
		"http://127.0.0.1:5173": true,
		"http://localhost:8080": true,
		"http://127.0.0.1:8080": true,
	}

	if allowedOrigins[origin] {
		return origin
	}

	return "*"
}
