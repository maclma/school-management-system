package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *models.Course) error
	FindByID(id uint) (*models.Course, error)
	FindByCourseCode(code string) (*models.Course, error)
	FindAll(page, limit int) ([]models.Course, int64, error)
	FindByDepartment(department string, page, limit int) ([]models.Course, int64, error)
	Update(course *models.Course) error
	Delete(id uint) error
	FindByTeacherID(teacherID uint) ([]models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{db: database.DB}
}

func (r *courseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) FindByID(id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("Teacher").Preload("Enrollments").First(&course, id).Error
	return &course, err
}

func (r *courseRepository) FindByCourseCode(code string) (*models.Course, error) {
	var course models.Course
	err := r.db.Where("course_code = ?", code).First(&course).Error
	return &course, err
}

func (r *courseRepository) FindAll(page, limit int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Course{}).Count(&total).
		Preload("Teacher").
		Limit(limit).
		Offset(offset).
		Find(&courses).Error

	return courses, total, err
}

func (r *courseRepository) FindByDepartment(department string, page, limit int) ([]models.Course, int64, error) {
	var courses []models.Course
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("department = ?", department).Count(&total).
		Preload("Teacher").
		Limit(limit).
		Offset(offset).
		Find(&courses).Error

	return courses, total, err
}

func (r *courseRepository) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id uint) error {
	return r.db.Delete(&models.Course{}, id).Error
}

func (r *courseRepository) FindByTeacherID(teacherID uint) ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Where("teacher_id = ?", teacherID).Preload("Enrollments").Find(&courses).Error
	return courses, err
}

// Course repository placeholder. Implement repository methods here as needed.
