package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"
	"time"

	"gorm.io/gorm"
)

type GradeRepository interface {
	Create(grade *models.Grade) error
	FindByID(id uint) (*models.Grade, error)
	FindByStudentAndCourse(studentID, courseID uint) (*models.Grade, error)
	FindByStudentID(studentID uint, page, limit int) ([]models.Grade, int64, error)
	FindByCourseID(courseID uint, page, limit int) ([]models.Grade, int64, error)
	FindAll(page, limit int) ([]models.Grade, int64, error)
	Update(grade *models.Grade) error
	Delete(id uint) error
	FindByTeacherID(teacherID uint, page, limit int) ([]models.Grade, int64, error)
	FindGradesInDateRange(startDate, endDate time.Time) ([]models.Grade, error)
}

type gradeRepository struct {
	db *gorm.DB
}

func NewGradeRepository() GradeRepository {
	return &gradeRepository{db: database.DB}
}

func (r *gradeRepository) Create(grade *models.Grade) error {
	return r.db.Create(grade).Error
}

func (r *gradeRepository) FindByID(id uint) (*models.Grade, error) {
	var grade models.Grade
	err := r.db.Preload("Student").Preload("Course").Preload("Teacher").First(&grade, id).Error
	return &grade, err
}

func (r *gradeRepository) FindByStudentAndCourse(studentID, courseID uint) (*models.Grade, error) {
	var grade models.Grade
	err := r.db.Where("student_id = ? AND course_id = ?", studentID, courseID).First(&grade).Error
	return &grade, err
}

func (r *gradeRepository) FindByStudentID(studentID uint, page, limit int) ([]models.Grade, int64, error) {
	var grades []models.Grade
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("student_id = ?", studentID).Count(&total).
		Preload("Course").
		Preload("Teacher").
		Limit(limit).
		Offset(offset).
		Find(&grades).Error

	return grades, total, err
}

func (r *gradeRepository) FindByCourseID(courseID uint, page, limit int) ([]models.Grade, int64, error) {
	var grades []models.Grade
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("course_id = ?", courseID).Count(&total).
		Preload("Student").
		Preload("Student.User").
		Limit(limit).
		Offset(offset).
		Find(&grades).Error

	return grades, total, err
}

func (r *gradeRepository) FindAll(page, limit int) ([]models.Grade, int64, error) {
	var grades []models.Grade
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Grade{}).Count(&total).
		Preload("Student").
		Preload("Course").
		Preload("Teacher").
		Limit(limit).
		Offset(offset).
		Find(&grades).Error

	return grades, total, err
}

func (r *gradeRepository) Update(grade *models.Grade) error {
	return r.db.Save(grade).Error
}

func (r *gradeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Grade{}, id).Error
}

func (r *gradeRepository) FindByTeacherID(teacherID uint, page, limit int) ([]models.Grade, int64, error) {
	var grades []models.Grade
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("graded_by = ?", teacherID).Count(&total).
		Preload("Student").
		Preload("Course").
		Limit(limit).
		Offset(offset).
		Find(&grades).Error

	return grades, total, err
}

func (r *gradeRepository) FindGradesInDateRange(startDate, endDate time.Time) ([]models.Grade, error) {
	var grades []models.Grade
	err := r.db.Where("graded_at BETWEEN ? AND ?", startDate, endDate).
		Preload("Student").
		Preload("Course").
		Find(&grades).Error

	return grades, err
}
