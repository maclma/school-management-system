package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	FindByID(id uint) (*models.Student, error)
	FindByUserID(userID uint) (*models.Student, error)
	FindByStudentID(studentID string) (*models.Student, error)
	FindAll(page, limit int) ([]models.Student, int64, error)
	FindByGradeLevel(gradeLevel string, page, limit int) ([]models.Student, int64, error)
	Update(student *models.Student) error
	Delete(id uint) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository() StudentRepository {
	return &studentRepository{db: database.DB}
}

func (r *studentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.Preload("User").Preload("Enrollments").Preload("Attendances").First(&student, id).Error
	return &student, err
}

func (r *studentRepository) FindByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	err := r.db.Where("user_id = ?", userID).Preload("User").First(&student).Error
	return &student, err
}

func (r *studentRepository) FindByStudentID(studentID string) (*models.Student, error) {
	var student models.Student
	err := r.db.Where("student_id = ?", studentID).Preload("User").First(&student).Error
	return &student, err
}

func (r *studentRepository) FindAll(page, limit int) ([]models.Student, int64, error) {
	var students []models.Student
	var total int64

	offset := (page - 1) * limit
	err := r.db.Model(&models.Student{}).Count(&total).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&students).Error

	return students, total, err
}

func (r *studentRepository) FindByGradeLevel(gradeLevel string, page, limit int) ([]models.Student, int64, error) {
	var students []models.Student
	var total int64

	offset := (page - 1) * limit
	err := r.db.Where("grade_level = ?", gradeLevel).Count(&total).
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&students).Error

	return students, total, err
}

func (r *studentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

// Student repository placeholder. Implement student repository methods here as needed.
