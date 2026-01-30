package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

/**
 * Database Optimization & Indexing Strategy
 * Creates indexes on frequently queried columns to improve performance
 */

// CreateIndexes adds all necessary database indexes
func CreateIndexes(db *gorm.DB) error {
	indexes := []struct {
		table   string
		columns []string
		unique  bool
		name    string
	}{
		// User indexes
		{table: "users", columns: []string{"email"}, unique: true, name: "idx_users_email"},
		{table: "users", columns: []string{"role"}, unique: false, name: "idx_users_role"},
		{table: "users", columns: []string{"is_active"}, unique: false, name: "idx_users_is_active"},
		{table: "users", columns: []string{"role", "is_active"}, unique: false, name: "idx_users_role_active"},

		// Student indexes
		{table: "students", columns: []string{"user_id"}, unique: false, name: "idx_students_user_id"},
		{table: "students", columns: []string{"student_id"}, unique: true, name: "idx_students_student_id"},
		{table: "students", columns: []string{"grade_level"}, unique: false, name: "idx_students_grade_level"},
		{table: "students", columns: []string{"enrollment_date"}, unique: false, name: "idx_students_enrollment_date"},

		// Teacher indexes
		{table: "teachers", columns: []string{"user_id"}, unique: false, name: "idx_teachers_user_id"},
		{table: "teachers", columns: []string{"department"}, unique: false, name: "idx_teachers_department"},
		{table: "teachers", columns: []string{"teacher_id"}, unique: true, name: "idx_teachers_teacher_id"},

		// Course indexes
		{table: "courses", columns: []string{"department"}, unique: false, name: "idx_courses_department"},
		{table: "courses", columns: []string{"course_code"}, unique: true, name: "idx_courses_code"},
		{table: "courses", columns: []string{"teacher_id"}, unique: false, name: "idx_courses_teacher_id"},
		{table: "courses", columns: []string{"department", "course_code"}, unique: false, name: "idx_courses_dept_code"},

		// Enrollment indexes (critical for filtering and stats)
		{table: "enrollments", columns: []string{"student_id"}, unique: false, name: "idx_enrollments_student_id"},
		{table: "enrollments", columns: []string{"course_id"}, unique: false, name: "idx_enrollments_course_id"},
		{table: "enrollments", columns: []string{"status"}, unique: false, name: "idx_enrollments_status"},
		{table: "enrollments", columns: []string{"student_id", "course_id"}, unique: false, name: "idx_enrollments_student_course"},
		{table: "enrollments", columns: []string{"status"}, unique: false, name: "idx_enrollments_status_date"},
		{table: "enrollments", columns: []string{"enrolled_at"}, unique: false, name: "idx_enrollments_enrolled_at"},

		// Grade indexes (critical for reporting)
		{table: "grades", columns: []string{"student_id"}, unique: false, name: "idx_grades_student_id"},
		{table: "grades", columns: []string{"course_id"}, unique: false, name: "idx_grades_course_id"},
		{table: "grades", columns: []string{"student_id", "course_id"}, unique: false, name: "idx_grades_student_course"},
		{table: "grades", columns: []string{"graded_at"}, unique: false, name: "idx_grades_graded_at"},

		// Attendance indexes
		{table: "attendances", columns: []string{"student_id"}, unique: false, name: "idx_attendance_student_id"},
		{table: "attendances", columns: []string{"course_id"}, unique: false, name: "idx_attendance_course_id"},
		{table: "attendances", columns: []string{"student_id", "course_id"}, unique: false, name: "idx_attendance_student_course"},
		{table: "attendances", columns: []string{"date"}, unique: false, name: "idx_attendance_date"},
		{table: "attendances", columns: []string{"status"}, unique: false, name: "idx_attendance_status"},

		// Assignment indexes
		{table: "assignments", columns: []string{"course_id"}, unique: false, name: "idx_assignments_course_id"},
		{table: "assignments", columns: []string{"created_by"}, unique: false, name: "idx_assignments_created_by"},
		{table: "assignments", columns: []string{"due_date"}, unique: false, name: "idx_assignments_due_date"},

		// Assignment submission indexes
		{table: "assignment_submissions", columns: []string{"assignment_id"}, unique: false, name: "idx_submissions_assignment_id"},
		{table: "assignment_submissions", columns: []string{"student_id"}, unique: false, name: "idx_submissions_student_id"},
		{table: "assignment_submissions", columns: []string{"assignment_id", "student_id"}, unique: false, name: "idx_submissions_assignment_student"},
		{table: "assignment_submissions", columns: []string{"submitted_at"}, unique: false, name: "idx_submissions_submitted_at"},
	}

	for _, idx := range indexes {
		columnStr := ""
		for i, col := range idx.columns {
			if i > 0 {
				columnStr += ", "
			}
			columnStr += col
		}

		var result *gorm.DB
		if idx.unique {
			result = db.Exec(fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS %s ON %s(%s)", idx.name, idx.table, columnStr))
		} else {
			result = db.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s(%s)", idx.name, idx.table, columnStr))
		}

		if result.Error != nil {
			fmt.Printf("Warning: Failed to create index %s: %v\n", idx.name, result.Error)
			// Continue creating other indexes even if one fails
		}
	}

	return nil
}

// OptimizeQueries applies GORM settings for better query performance
func OptimizeQueries(db *gorm.DB) {
	// Enable prepared statement caching
	// Already enabled by default in GORM

	// Set query timeout for safety
	// Can be done per query if needed
}

// QueryStats represents database query statistics
type QueryStats struct {
	TotalQueries  int64
	CachedQueries int64
	SlowQueries   int64
	ErrorQueries  int64
	AvgDuration   time.Duration
}

// AnalyzeSlowQueries logs slow queries (queries taking > threshold)
func AnalyzeSlowQueries(db *gorm.DB, thresholdMs int64) {
	// This is a placeholder for query analysis
	// In production, use a database query log analyzer
}

/**
 * Query optimization tips for common operations:
 *
 * 1. Enrollment listing with stats:
 *    - Use SELECT with specific columns instead of *
 *    - Join with users table when needed for name filtering
 *    - Index on status and created_at for filtering
 *
 * 2. Student grade reports:
 *    - Index on student_id and course_id for lookups
 *    - Use database-level aggregation (AVG, SUM) instead of app-level
 *
 * 3. Attendance tracking:
 *    - Index on student_id, course_id for quick lookups
 *    - Index on attendance_date for date range queries
 *
 * 4. Dashboard statistics:
 *    - Cache count results with expiry (e.g., 5 minutes)
 *    - Use database COUNT with WHERE instead of app-level counting
 *
 * 5. Search/Filter operations:
 *    - Add indexes on frequently filtered columns
 *    - Use LIKE queries sparingly, prefer exact match where possible
 *    - Consider adding full-text search indexes for text fields
 */

// CachedCount provides in-memory caching for count queries
type CachedCount struct {
	cache      map[string]int64
	timestamps map[string]time.Time
	cacheTTL   time.Duration
}

func NewCachedCount(ttl time.Duration) *CachedCount {
	return &CachedCount{
		cache:      make(map[string]int64),
		timestamps: make(map[string]time.Time),
		cacheTTL:   ttl,
	}
}

func (cc *CachedCount) Get(key string, fn func() (int64, error)) (int64, error) {
	// Check if cached value is still valid
	if val, exists := cc.cache[key]; exists {
		if time.Since(cc.timestamps[key]) < cc.cacheTTL {
			return val, nil
		}
	}

	// Cache miss or expired, compute value
	val, err := fn()
	if err != nil {
		return 0, err
	}

	cc.cache[key] = val
	cc.timestamps[key] = time.Now()
	return val, nil
}

func (cc *CachedCount) Invalidate(key string) {
	delete(cc.cache, key)
	delete(cc.timestamps, key)
}

func (cc *CachedCount) Clear() {
	cc.cache = make(map[string]int64)
	cc.timestamps = make(map[string]time.Time)
}
