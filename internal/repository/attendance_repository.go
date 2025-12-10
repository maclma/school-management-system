package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(attendance *models.Attendance) error
	FindByID(id uint) (*models.Attendance, error)
	FindByStudentAndCourse(studentID, courseID uint, page, limit int) ([]models.Attendance, int64, error)
	FindByStudentID(studentID uint, page, limit int) ([]models.Attendance, int64, error)
	FindByCourseID(courseID uint, page, limit int) ([]models.Attendance, int64, error)
	FindAll(page, limit int) ([]models.Attendance, int64, error)
	Update(attendance *models.Attendance) error
	Delete(id uint) error
	FindByDateRange(startDate, endDate time.Time) ([]models.Attendance, error)
	FindByStudentDateRange(studentID uint, startDate, endDate time.Time) ([]models.Attendance, error)
	CountAttendanceByStudent(studentID, courseID uint) (present, absent, late int64, error error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository() AttendanceRepository {
	return &attendanceRepository{db: database.DB}
}

func (r *attendanceRepository) Create(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *attendanceRepository) FindByID(id uint) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Preload("Student").Preload("Course").First(&attendance, id).Error
	return &attendance, err
}

func (r *attendanceRepository) FindByStudentAndCourse(studentID, courseID uint, page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&attendances).Error

	return attendances, total, err
}

func (r *attendanceRepository) FindByStudentID(studentID uint, page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("student_id = ?", studentID).Count(&total).
		Preload("Course").
		Limit(limit).
		Offset(offset).
		Find(&attendances).Error

	return attendances, total, err
}

func (r *attendanceRepository) FindByCourseID(courseID uint, page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("course_id = ?", courseID).Count(&total).
		Preload("Student").
		Preload("Student.User").
		Limit(limit).
		Offset(offset).
		Find(&attendances).Error

	return attendances, total, err
}

func (r *attendanceRepository) FindAll(page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Attendance{}).Count(&total).
		Preload("Student").
		Preload("Course").
		Limit(limit).
		Offset(offset).
		Find(&attendances).Error

	return attendances, total, err
}

func (r *attendanceRepository) Update(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *attendanceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Attendance{}, id).Error
}

func (r *attendanceRepository) FindByDateRange(startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).
		Preload("Student").
		Preload("Course").
		Find(&attendances).Error

	return attendances, err
}

func (r *attendanceRepository) FindByStudentDateRange(studentID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Where("student_id = ? AND date BETWEEN ? AND ?", studentID, startDate, endDate).
		Preload("Course").
		Find(&attendances).Error

	return attendances, err
}

func (r *attendanceRepository) CountAttendanceByStudent(studentID, courseID uint) (present, absent, late int64, error error) {
	error = r.db.Where("student_id = ? AND course_id = ? AND status = 'present'", studentID, courseID).Count(&present).Error
	if error != nil {
		return
	}

	error = r.db.Where("student_id = ? AND course_id = ? AND status = 'absent'", studentID, courseID).Count(&absent).Error
	if error != nil {
		return
	}

	error = r.db.Where("student_id = ? AND course_id = ? AND status = 'late'", studentID, courseID).Count(&late).Error
	return
}
