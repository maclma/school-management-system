package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type GradeTranscriptService interface {
	Create(transcript *models.GradeTranscript) error
	GetByID(id uint) (*models.GradeTranscript, error)
	GetByStudentID(studentID uint, page, limit int) ([]models.GradeTranscript, int64, error)
	GetByStudentAndYear(studentID uint, year int) (*models.GradeTranscript, error)
	Update(transcript *models.GradeTranscript) error
	Delete(id uint) error
	GetLatestByStudent(studentID uint) (*models.GradeTranscript, error)
	CalculateGPA(studentID uint) (float64, error)
}

type gradeTranscriptService struct {
	repo repository.GradeTranscriptRepository
}

func NewGradeTranscriptService(repo repository.GradeTranscriptRepository) GradeTranscriptService {
	return &gradeTranscriptService{repo: repo}
}

func (s *gradeTranscriptService) Create(transcript *models.GradeTranscript) error {
	return s.repo.Create(transcript)
}

func (s *gradeTranscriptService) GetByID(id uint) (*models.GradeTranscript, error) {
	return s.repo.FindByID(id)
}

func (s *gradeTranscriptService) GetByStudentID(studentID uint, page, limit int) ([]models.GradeTranscript, int64, error) {
	return s.repo.FindByStudentID(studentID, page, limit)
}

func (s *gradeTranscriptService) GetByStudentAndYear(studentID uint, year int) (*models.GradeTranscript, error) {
	return s.repo.FindByStudentAndYear(studentID, year)
}

func (s *gradeTranscriptService) Update(transcript *models.GradeTranscript) error {
	return s.repo.Update(transcript)
}

func (s *gradeTranscriptService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *gradeTranscriptService) GetLatestByStudent(studentID uint) (*models.GradeTranscript, error) {
	return s.repo.FindLatestByStudent(studentID)
}

func (s *gradeTranscriptService) CalculateGPA(studentID uint) (float64, error) {
	transcript, err := s.repo.FindLatestByStudent(studentID)
	if err != nil {
		return 0, err
	}
	if transcript.TotalCredits == 0 {
		return 0, nil
	}
	return transcript.GradePointsSum / transcript.TotalCredits, nil
}
