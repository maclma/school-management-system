package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitStore tracks request counts per client
type RateLimitStore struct {
	requests map[string][]time.Time
	mu       sync.Mutex
}

// NewRateLimitStore creates a new rate limit store
func NewRateLimitStore() *RateLimitStore {
	store := &RateLimitStore{
		requests: make(map[string][]time.Time),
	}

	// Clean up expired entries every minute
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			store.cleanup()
		}
	}()

	return store
}

// IsAllowed checks if a request is allowed for the given client
func (s *RateLimitStore) IsAllowed(clientID string, limit int, window time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-window)

	// Get or create request list for this client
	requests := s.requests[clientID]

	// Remove old requests outside the window
	var validRequests []time.Time
	for _, req := range requests {
		if req.After(cutoff) {
			validRequests = append(validRequests, req)
		}
	}

	// Check if limit exceeded
	if len(validRequests) >= limit {
		s.requests[clientID] = validRequests
		return false
	}

	// Add current request
	validRequests = append(validRequests, now)
	s.requests[clientID] = validRequests

	return true
}

// cleanup removes entries with no requests
func (s *RateLimitStore) cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	cutoff := time.Now().Add(-10 * time.Minute)

	for clientID, requests := range s.requests {
		var validRequests []time.Time
		for _, req := range requests {
			if req.After(cutoff) {
				validRequests = append(validRequests, req)
			}
		}

		if len(validRequests) == 0 {
			delete(s.requests, clientID)
		} else {
			s.requests[clientID] = validRequests
		}
	}
}

var rateLimitStore = NewRateLimitStore()

// RateLimitMiddleware creates a middleware that rate limits requests per IP
// limit: max requests, window: time duration for limit
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !rateLimitStore.IsAllowed(clientIP, limit, window) {
			c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
			c.Header("X-RateLimit-Window", window.String())
			c.Header("Retry-After", fmt.Sprintf("%d", int(window.Seconds())))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":   "rate limit exceeded",
				"message": fmt.Sprintf("maximum %d requests per %v", limit, window),
			})
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", limit-1)) // Simplified
		c.Next()
	}
}

// APIRateLimit applies rate limit to API routes (higher limit)
func APIRateLimit() gin.HandlerFunc {
	return RateLimitMiddleware(100, 1*time.Minute) // 100 requests per minute
}

// AuthRateLimit applies rate limit to auth routes (lower limit)
func AuthRateLimit() gin.HandlerFunc {
	return RateLimitMiddleware(10, 1*time.Minute) // 10 requests per minute
}

// StrictRateLimit applies strict rate limit (for admin operations)
func StrictRateLimit() gin.HandlerFunc {
	return RateLimitMiddleware(50, 1*time.Minute) // 50 requests per minute
}
