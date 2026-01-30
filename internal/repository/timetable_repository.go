package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type TimeTableRepository interface {
	Create(timetable *models.TimeTable) error
	FindByID(id uint) (*models.TimeTable, error)
	FindByCourseID(courseID uint) ([]models.TimeTable, error)
	FindByTeacherID(teacherID uint) ([]models.TimeTable, error)
	FindByDayOfWeek(dayOfWeek string) ([]models.TimeTable, error)
	FindAll() ([]models.TimeTable, error)
	Update(timetable *models.TimeTable) error
	Delete(id uint) error
}

type timetableRepository struct {
	db *gorm.DB
}

func NewTimeTableRepository() TimeTableRepository {
	return &timetableRepository{db: database.DB}
}

func (r *timetableRepository) Create(timetable *models.TimeTable) error {
	return r.db.Create(timetable).Error
}

func (r *timetableRepository) FindByID(id uint) (*models.TimeTable, error) {
	var timetable models.TimeTable
	err := r.db.Preload("Course").Preload("Teacher").First(&timetable, id).Error
	return &timetable, err
}

func (r *timetableRepository) FindByCourseID(courseID uint) ([]models.TimeTable, error) {
	var timetables []models.TimeTable
	err := r.db.Where("course_id = ? AND is_active = ?", courseID, true).
		Preload("Course").
		Preload("Teacher").
		Find(&timetables).Error
	return timetables, err
}

func (r *timetableRepository) FindByTeacherID(teacherID uint) ([]models.TimeTable, error) {
	var timetables []models.TimeTable
	err := r.db.Where("teacher_id = ? AND is_active = ?", teacherID, true).
		Preload("Course").
		Preload("Teacher").
		Find(&timetables).Error
	return timetables, err
}

func (r *timetableRepository) FindByDayOfWeek(dayOfWeek string) ([]models.TimeTable, error) {
	var timetables []models.TimeTable
	err := r.db.Where("day_of_week = ? AND is_active = ?", dayOfWeek, true).
		Preload("Course").
		Preload("Teacher").
		Find(&timetables).Error
	return timetables, err
}

func (r *timetableRepository) FindAll() ([]models.TimeTable, error) {
	var timetables []models.TimeTable
	err := r.db.Preload("Course").Preload("Teacher").Find(&timetables).Error
	return timetables, err
}

func (r *timetableRepository) Update(timetable *models.TimeTable) error {
	return r.db.Save(timetable).Error
}

func (r *timetableRepository) Delete(id uint) error {
	return r.db.Delete(&models.TimeTable{}, id).Error
}
