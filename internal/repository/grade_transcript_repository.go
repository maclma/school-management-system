package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type GradeTranscriptRepository interface {
	Create(transcript *models.GradeTranscript) error
	FindByID(id uint) (*models.GradeTranscript, error)
	FindByStudentID(studentID uint, page, limit int) ([]models.GradeTranscript, int64, error)
	FindByStudentAndYear(studentID uint, year int) (*models.GradeTranscript, error)
	Update(transcript *models.GradeTranscript) error
	Delete(id uint) error
	FindLatestByStudent(studentID uint) (*models.GradeTranscript, error)
}

type gradeTranscriptRepository struct {
	db *gorm.DB
}

func NewGradeTranscriptRepository() GradeTranscriptRepository {
	return &gradeTranscriptRepository{db: database.DB}
}

func (r *gradeTranscriptRepository) Create(transcript *models.GradeTranscript) error {
	return r.db.Create(transcript).Error
}

func (r *gradeTranscriptRepository) FindByID(id uint) (*models.GradeTranscript, error) {
	var transcript models.GradeTranscript
	err := r.db.Preload("Student").First(&transcript, id).Error
	return &transcript, err
}

func (r *gradeTranscriptRepository) FindByStudentID(studentID uint, page, limit int) ([]models.GradeTranscript, int64, error) {
	var transcripts []models.GradeTranscript
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("student_id = ?", studentID).Count(&total).
		Preload("Student").
		Order("year DESC, transcript_semester DESC").
		Limit(limit).
		Offset(offset).
		Find(&transcripts).Error
	return transcripts, total, err
}

func (r *gradeTranscriptRepository) FindByStudentAndYear(studentID uint, year int) (*models.GradeTranscript, error) {
	var transcript models.GradeTranscript
	err := r.db.Where("student_id = ? AND year = ?", studentID, year).
		Preload("Student").
		First(&transcript).Error
	return &transcript, err
}

func (r *gradeTranscriptRepository) Update(transcript *models.GradeTranscript) error {
	return r.db.Save(transcript).Error
}

func (r *gradeTranscriptRepository) Delete(id uint) error {
	return r.db.Delete(&models.GradeTranscript{}, id).Error
}

func (r *gradeTranscriptRepository) FindLatestByStudent(studentID uint) (*models.GradeTranscript, error) {
	var transcript models.GradeTranscript
	err := r.db.Where("student_id = ?", studentID).
		Preload("Student").
		Order("year DESC, transcript_semester DESC").
		First(&transcript).Error
	return &transcript, err
}
