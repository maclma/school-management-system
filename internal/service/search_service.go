package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/database"
	"time"
)

// SearchService provides advanced search capabilities
type SearchService struct {
	announcementRepo repository.AnnouncementRepository
	paymentRepo      repository.PaymentRepository
	studentRepo      repository.StudentRepository
}

// NewSearchService creates a new search service
func NewSearchService(
	announcementRepo repository.AnnouncementRepository,
	paymentRepo repository.PaymentRepository,
	studentRepo repository.StudentRepository,
) *SearchService {
	return &SearchService{
		announcementRepo: announcementRepo,
		paymentRepo:      paymentRepo,
		studentRepo:      studentRepo,
	}
}

// SearchAnnouncementsAdvanced searches announcements with filters
func (s *SearchService) SearchAnnouncementsAdvanced(query string, audience string, priority string, page int, limit int) ([]models.Announcement, int64, error) {
	db := database.DB
	var announcements []models.Announcement
	var total int64

	q := db.Where("is_active = ?", true)

	if query != "" {
		q = q.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if audience != "" {
		q = q.Where("audience = ?", audience)
	}
	if priority != "" {
		q = q.Where("priority = ?", priority)
	}

	q.Model(&models.Announcement{}).Count(&total)

	offset := (page - 1) * limit
	if err := q.Offset(offset).Limit(limit).Find(&announcements).Error; err != nil {
		return nil, 0, err
	}

	return announcements, total, nil
}

// SearchPayments searches payments with filters
func (s *SearchService) SearchPayments(studentID uint, status string, page int, limit int) ([]models.Payment, int64, error) {
	db := database.DB
	var payments []models.Payment
	var total int64

	q := db.Where("1=1")

	if studentID > 0 {
		q = q.Where("student_id = ?", studentID)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}

	q.Model(&models.Payment{}).Count(&total)

	offset := (page - 1) * limit
	if err := q.Offset(offset).Limit(limit).Order("created_at DESC").Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}

// SearchStudents searches students by name, email, ID
func (s *SearchService) SearchStudents(query string, page int, limit int) ([]models.Student, int64, error) {
	db := database.DB
	var students []models.Student
	var total int64

	like := "%" + query + "%"
	q := db.Model(&models.Student{}).
		Joins("JOIN users ON users.id = students.user_id").
		Where("users.first_name LIKE ? OR users.last_name LIKE ? OR users.email LIKE ? OR students.student_id LIKE ?", like, like, like, like)

	q.Count(&total)

	offset := (page - 1) * limit
	if err := q.Offset(offset).Limit(limit).Preload("User").Find(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, total, nil
}

// SearchGradesByRange searches grades in a point range
func (s *SearchService) SearchGradesByRange(courseID uint, minScore float64, maxScore float64, page int, limit int) ([]models.Grade, int64, error) {
	db := database.DB
	var grades []models.Grade
	var total int64

	q := db.Where("course_id = ? AND score >= ? AND score <= ?", courseID, minScore, maxScore)

	q.Model(&models.Grade{}).Count(&total)

	offset := (page - 1) * limit
	if err := q.Offset(offset).Limit(limit).Find(&grades).Error; err != nil {
		return nil, 0, err
	}

	return grades, total, nil
}

// SearchOverduePayments finds all overdue payments
func (s *SearchService) SearchOverduePayments() ([]models.Payment, error) {
	db := database.DB
	var payments []models.Payment

	now := time.Now().Unix()
	if err := db.Where("status = ? AND due_date < ?", "pending", now).Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

// SearchLowAttendanceStudents finds students with attendance below threshold
func (s *SearchService) SearchLowAttendanceStudents(courseID uint, attendanceThreshold float64) ([]models.Student, error) {
	db := database.DB
	var students []models.Student

	// Subquery to calculate attendance percentage
	err := db.Raw(`
		SELECT DISTINCT s.* FROM students s
		JOIN enrollments e ON s.id = e.student_id
		WHERE e.course_id = ?
		AND (SELECT COUNT(*) FROM attendance WHERE student_id = s.id AND course_id = ? AND present = false)
			/ 
		(SELECT COUNT(*) FROM attendance WHERE student_id = s.id AND course_id = ?) * 100 > ?
	`, courseID, courseID, courseID, attendanceThreshold).Scan(&students).Error

	if err != nil {
		return nil, err
	}

	return students, nil
}
