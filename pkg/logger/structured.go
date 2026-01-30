package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

// StructuredLogger wraps logrus with structured logging methods
type StructuredLogger struct {
	*logrus.Logger
}

// WithRequestID adds request ID to log entry
func (sl *StructuredLogger) WithRequestID(requestID string) *logrus.Entry {
	return sl.WithField("request_id", requestID)
}

// LogRequest logs incoming HTTP request
func (sl *StructuredLogger) LogRequest(method, path, ip string, requestID string) {
	sl.WithFields(logrus.Fields{
		"request_id": requestID,
		"method":     method,
		"path":       path,
		"ip":         ip,
		"timestamp":  time.Now().Unix(),
	}).Infof("Request: %s %s from %s", method, path, ip)
}

// LogResponse logs HTTP response
func (sl *StructuredLogger) LogResponse(statusCode int, duration time.Duration, requestID string) {
	sl.WithFields(logrus.Fields{
		"request_id":  requestID,
		"status":      statusCode,
		"duration_ms": duration.Milliseconds(),
		"timestamp":   time.Now().Unix(),
	}).Infof("Response: %d in %v", statusCode, duration)
}

// LogError logs error with context
func (sl *StructuredLogger) LogError(err error, context string, requestID string, fields ...interface{}) {
	entry := sl.WithFields(logrus.Fields{
		"error":      err.Error(),
		"context":    context,
		"request_id": requestID,
		"timestamp":  time.Now().Unix(),
	})

	// Add any additional fields
	if len(fields) > 0 {
		for i := 0; i < len(fields)-1; i += 2 {
			entry = entry.WithField(fields[i].(string), fields[i+1])
		}
	}

	entry.Errorf("Error in %s: %v", context, err)
}

// LogDatabaseOperation logs database operations
func (sl *StructuredLogger) LogDatabaseOperation(operation string, table string, duration time.Duration, rowsAffected int64, requestID string) {
	sl.WithFields(logrus.Fields{
		"request_id":    requestID,
		"operation":     operation,
		"table":         table,
		"duration_ms":   duration.Milliseconds(),
		"rows_affected": rowsAffected,
		"timestamp":     time.Now().Unix(),
	}).Debugf("DB Operation: %s on %s in %v", operation, table, duration)
}

// LogServiceCall logs service method calls
func (sl *StructuredLogger) LogServiceCall(service string, method string, duration time.Duration, requestID string) {
	sl.WithFields(logrus.Fields{
		"request_id":  requestID,
		"service":     service,
		"method":      method,
		"duration_ms": duration.Milliseconds(),
		"timestamp":   time.Now().Unix(),
	}).Debugf("Service: %s.%s executed in %v", service, method, duration)
}

// LogAuthenticationEvent logs authentication events
func (sl *StructuredLogger) LogAuthenticationEvent(eventType string, userID uint, email string, success bool, requestID string, reason string) {
	level := logrus.InfoLevel
	if !success {
		level = logrus.WarnLevel
	}

	entry := sl.WithFields(logrus.Fields{
		"request_id": requestID,
		"event_type": eventType,
		"user_id":    userID,
		"email":      email,
		"success":    success,
		"timestamp":  time.Now().Unix(),
	})

	if reason != "" {
		entry = entry.WithField("reason", reason)
	}

	entry.Logf(level, "Auth Event: %s for %s - %v", eventType, email, success)
}

// LogAuthorizationEvent logs authorization/permission checks
func (sl *StructuredLogger) LogAuthorizationEvent(userID uint, resource string, action string, allowed bool, requestID string) {
	level := logrus.InfoLevel
	if !allowed {
		level = logrus.WarnLevel
	}

	sl.WithFields(logrus.Fields{
		"request_id": requestID,
		"user_id":    userID,
		"resource":   resource,
		"action":     action,
		"allowed":    allowed,
		"timestamp":  time.Now().Unix(),
	}).Logf(level, "Authorization: User %d attempting %s on %s - %v", userID, action, resource, allowed)
}

// LogDataModification logs data create/update/delete operations
func (sl *StructuredLogger) LogDataModification(operation string, entity string, entityID uint, userID uint, details interface{}, requestID string) {
	sl.WithFields(logrus.Fields{
		"request_id": requestID,
		"operation":  operation,
		"entity":     entity,
		"entity_id":  entityID,
		"user_id":    userID,
		"details":    details,
		"timestamp":  time.Now().Unix(),
	}).Infof("Data Modification: %s %s(ID:%d) by User(ID:%d)", operation, entity, entityID, userID)
}

// ToStructured converts standard logger to StructuredLogger
func ToStructured(log *logrus.Logger) *StructuredLogger {
	return &StructuredLogger{log}
}
