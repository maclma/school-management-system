package service

import (
	"time"

	"school-management-system/internal/models"
	"school-management-system/pkg/database"
)

// AttendanceAutomationService provides automated attendance tracking
type AttendanceAutomationService struct {
	emailService *EmailService
}

// NewAttendanceAutomationService creates a new service
func NewAttendanceAutomationService(emailService *EmailService) *AttendanceAutomationService {
	return &AttendanceAutomationService{
		emailService: emailService,
	}
}

// CalculateAttendancePercentage calculates student's attendance percentage
func (aas *AttendanceAutomationService) CalculateAttendancePercentage(studentID uint, courseID uint) (float64, error) {
	db := database.DB

	var present, total int64

	// Count present days
	if err := db.Model(&models.Attendance{}).
		Where("student_id = ? AND course_id = ? AND present = ?", studentID, courseID, true).
		Count(&present).Error; err != nil {
		return 0, err
	}

	// Count total days
	if err := db.Model(&models.Attendance{}).
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		Count(&total).Error; err != nil {
		return 0, err
	}

	if total == 0 {
		return 0, nil
	}

	percentage := (float64(present) / float64(total)) * 100
	return percentage, nil
}

// CheckLowAttendance checks if attendance is below threshold and sends alert
func (aas *AttendanceAutomationService) CheckLowAttendance(studentID uint, courseID uint, threshold float64) (bool, error) {
	percentage, err := aas.CalculateAttendancePercentage(studentID, courseID)
	if err != nil {
		return false, err
	}

	if percentage < threshold {
		// Get student info
		db := database.DB
		var student models.Student
		if err := db.First(&student, studentID).Error; err == nil {
			// Get course info
			var course models.Course
			if err := db.First(&course, courseID).Error; err == nil {
				// Send email alert
				aas.emailService.SendAttendanceAlert(
					student.User.Email,
					student.User.FirstName+" "+student.User.LastName,
					course.Name,
					percentage,
				)
			}
		}
		return true, nil
	}

	return false, nil
}

// RecordAttendanceAndCheck records attendance and checks for low attendance
func (aas *AttendanceAutomationService) RecordAttendanceAndCheck(attendance *models.Attendance, attendanceThreshold float64) error {
	db := database.DB

	// Create attendance record
	if err := db.Create(attendance).Error; err != nil {
		return err
	}

	// Check attendance percentage
	aas.CheckLowAttendance(attendance.StudentID, attendance.CourseID, attendanceThreshold)

	return nil
}

// GetAttendanceStats returns attendance statistics for a course
func (aas *AttendanceAutomationService) GetAttendanceStats(courseID uint) (map[string]interface{}, error) {
	db := database.DB

	var totalSessions, totalStudents int64
	var avgAttendance float64

	// Count total sessions
	if err := db.Model(&models.Attendance{}).
		Where("course_id = ?", courseID).
		Distinct("DATE(date)").
		Count(&totalSessions).Error; err != nil {
		return nil, err
	}

	// Count students in course
	if err := db.Model(&models.Enrollment{}).
		Where("course_id = ?", courseID).
		Count(&totalStudents).Error; err != nil {
		return nil, err
	}

	// Calculate average attendance
	var result struct {
		AvgAttendance float64
	}
	if err := db.Raw(`
		SELECT AVG(attendance_pct) as avg_attendance FROM (
			SELECT 
				student_id,
				COUNT(CASE WHEN present = true THEN 1 END) * 100.0 / COUNT(*) as attendance_pct
			FROM attendance
			WHERE course_id = ?
			GROUP BY student_id
		)
	`, courseID).Scan(&result).Error; err == nil {
		avgAttendance = result.AvgAttendance
	}

	return map[string]interface{}{
		"total_sessions":     totalSessions,
		"total_students":     totalStudents,
		"average_attendance": avgAttendance,
		"last_updated_at":    time.Now(),
	}, nil
}

// GetStudentAttendanceStatusByThreshold returns students below attendance threshold
func (aas *AttendanceAutomationService) GetStudentAttendanceStatusByThreshold(courseID uint, threshold float64) ([]map[string]interface{}, error) {
	db := database.DB

	var results []map[string]interface{}

	rows, err := db.Raw(`
		SELECT 
			s.id,
			u.first_name,
			u.last_name,
			u.email,
			COUNT(CASE WHEN a.status = 'present' THEN 1 END) * 100.0 / COUNT(*) as attendance_percentage
		FROM students s
		JOIN users u ON s.user_id = u.id
		JOIN enrollments e ON s.id = e.student_id
		LEFT JOIN attendance a ON s.id = a.student_id AND e.course_id = a.course_id
		WHERE e.course_id = ?
		GROUP BY s.id, u.first_name, u.last_name, u.email
		HAVING attendance_percentage < ?
	`, courseID, threshold).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint
		var firstName, lastName, email string
		var attendancePercentage float64

		if err := rows.Scan(&id, &firstName, &lastName, &email, &attendancePercentage); err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"student_id":            id,
			"name":                  firstName + " " + lastName,
			"email":                 email,
			"attendance_percentage": attendancePercentage,
		})
	}

	return results, nil
}

// GenerateAttendanceReport generates a detailed attendance report
func (aas *AttendanceAutomationService) GenerateAttendanceReport(courseID uint) (map[string]interface{}, error) {
	db := database.DB

	stats, err := aas.GetAttendanceStats(courseID)
	if err != nil {
		return nil, err
	}

	lowAttendance, err := aas.GetStudentAttendanceStatusByThreshold(courseID, 80)
	if err != nil {
		return nil, err
	}

	// Get course info
	var course models.Course
	db.First(&course, courseID)

	return map[string]interface{}{
		"course_id":            courseID,
		"course_name":          course.Name,
		"stats":                stats,
		"students_below_80pct": lowAttendance,
		"report_generated_at":  time.Now(),
	}, nil
}
