package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	Create(enrollment *models.Enrollment) error
	FindByID(id uint) (*models.Enrollment, error)
	FindByStudentAndCourse(studentID, courseID uint) (*models.Enrollment, error)
	FindByStudentID(studentID uint, page, limit int) ([]models.Enrollment, int64, error)
	FindByCourseID(courseID uint, page, limit int) ([]models.Enrollment, int64, error)
	FindAll(page, limit int) ([]models.Enrollment, int64, error)
	Update(enrollment *models.Enrollment) error
	Delete(id uint) error
	CountByCourseID(courseID uint) (int64, error)
}

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository() EnrollmentRepository {
	return &enrollmentRepository{db: database.DB}
}

func (r *enrollmentRepository) Create(enrollment *models.Enrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *enrollmentRepository) FindByID(id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.Preload("Student").Preload("Course").First(&enrollment, id).Error
	return &enrollment, err
}

func (r *enrollmentRepository) FindByStudentAndCourse(studentID, courseID uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).First(&enrollment).Error
	return &enrollment, err
}

func (r *enrollmentRepository) FindByStudentID(studentID uint, page, limit int) ([]models.Enrollment, int64, error) {
	var enrollments []models.Enrollment
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("student_id = ?", studentID).Count(&total).
		Preload("Course").
		Preload("Course.Teacher").
		Limit(limit).
		Offset(offset).
		Find(&enrollments).Error

	return enrollments, total, err
}

func (r *enrollmentRepository) FindByCourseID(courseID uint, page, limit int) ([]models.Enrollment, int64, error) {
	var enrollments []models.Enrollment
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("course_id = ?", courseID).Count(&total).
		Preload("Student").
		Preload("Student.User").
		Limit(limit).
		Offset(offset).
		Find(&enrollments).Error

	return enrollments, total, err
}

func (r *enrollmentRepository) FindAll(page, limit int) ([]models.Enrollment, int64, error) {
	var enrollments []models.Enrollment
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Enrollment{}).Count(&total).
		Preload("Student").
		Preload("Course").
		Limit(limit).
		Offset(offset).
		Find(&enrollments).Error

	return enrollments, total, err
}

func (r *enrollmentRepository) Update(enrollment *models.Enrollment) error {
	return r.db.Save(enrollment).Error
}

func (r *enrollmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Enrollment{}, id).Error
}

func (r *enrollmentRepository) CountByCourseID(courseID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Enrollment{}).Where("course_id = ? AND status = 'active'", courseID).Count(&count).Error
	return count, err
}
