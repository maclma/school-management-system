package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"
)

type TeacherRepository interface {
	Create(teacher *models.Teacher) error
	GetByID(id uint) (*models.Teacher, error)
	GetByUserID(userID uint) (*models.Teacher, error)
	GetByTeacherID(teacherID string) (*models.Teacher, error)
	GetAll(page, limit int) ([]models.Teacher, int64, error)
	Update(teacher *models.Teacher) error
	Delete(id uint) error
	GetTeacherCourses(teacherID uint, page, limit int) ([]models.Course, int64, error)
	GetByDepartment(department string, page, limit int) ([]models.Teacher, int64, error)
}

type teacherRepository struct{}

func NewTeacherRepository() TeacherRepository {
	return &teacherRepository{}
}

func (r *teacherRepository) Create(teacher *models.Teacher) error {
	db := database.GetDB()
	return db.Create(teacher).Error
}

func (r *teacherRepository) GetByID(id uint) (*models.Teacher, error) {
	db := database.GetDB()
	var teacher models.Teacher
	err := db.Preload("User").Preload("Courses").First(&teacher, id).Error
	return &teacher, err
}

func (r *teacherRepository) GetByUserID(userID uint) (*models.Teacher, error) {
	db := database.GetDB()
	var teacher models.Teacher
	err := db.Preload("User").Preload("Courses").Where("user_id = ?", userID).First(&teacher).Error
	return &teacher, err
}

func (r *teacherRepository) GetByTeacherID(teacherID string) (*models.Teacher, error) {
	db := database.GetDB()
	var teacher models.Teacher
	err := db.Preload("User").Preload("Courses").Where("teacher_id = ?", teacherID).First(&teacher).Error
	return &teacher, err
}

func (r *teacherRepository) GetAll(page, limit int) ([]models.Teacher, int64, error) {
	db := database.GetDB()
	var teachers []models.Teacher
	var total int64

	db.Model(&models.Teacher{}).Count(&total)

	offset := (page - 1) * limit
	err := db.Preload("User").Offset(offset).Limit(limit).Find(&teachers).Error

	return teachers, total, err
}

func (r *teacherRepository) Update(teacher *models.Teacher) error {
	db := database.GetDB()
	return db.Model(&models.Teacher{}).Where("id = ?", teacher.ID).Updates(teacher).Error
}

func (r *teacherRepository) Delete(id uint) error {
	db := database.GetDB()
	return db.Delete(&models.Teacher{}, id).Error
}

func (r *teacherRepository) GetTeacherCourses(teacherID uint, page, limit int) ([]models.Course, int64, error) {
	db := database.GetDB()
	var courses []models.Course
	var total int64

	db.Model(&models.Course{}).Where("teacher_id = ?", teacherID).Count(&total)

	offset := (page - 1) * limit
	err := db.Preload("Teacher").Where("teacher_id = ?", teacherID).Offset(offset).Limit(limit).Find(&courses).Error

	return courses, total, err
}

func (r *teacherRepository) GetByDepartment(department string, page, limit int) ([]models.Teacher, int64, error) {
	db := database.GetDB()
	var teachers []models.Teacher
	var total int64

	db.Model(&models.Teacher{}).Where("department = ?", department).Count(&total)

	offset := (page - 1) * limit
	err := db.Preload("User").Where("department = ?", department).Offset(offset).Limit(limit).Find(&teachers).Error

	return teachers, total, err
}
