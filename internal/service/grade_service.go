package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
)

type GradeService interface {
	RecordGrade(grade *models.Grade) error
	GetGradeByID(id uint) (*models.Grade, error)
	GetStudentGrades(studentID uint, page, limit int) ([]models.Grade, int64, error)
	GetCourseGrades(courseID uint, page, limit int) ([]models.Grade, int64, error)
	GetStudentCourseGrade(studentID, courseID uint) (*models.Grade, error)
	GetAllGrades(page, limit int) ([]models.Grade, int64, error)
	UpdateGrade(grade *models.Grade) error
	DeleteGrade(id uint) error
	GetTeacherGrades(teacherID uint, page, limit int) ([]models.Grade, int64, error)
	GetGradesInRange(startDate, endDate time.Time) ([]models.Grade, error)
	CalculateAverageGrade(studentID uint) (float64, error)
}

type gradeService struct {
	gradeRepo repository.GradeRepository
	logger    *logrus.Logger
}

func NewGradeService(gradeRepo repository.GradeRepository) GradeService {
	return &gradeService{
		gradeRepo: gradeRepo,
		logger:    logger.GetLogger(),
	}
}

func (s *gradeService) RecordGrade(grade *models.Grade) error {
	if grade.StudentID == 0 {
		s.logger.Warn("Student ID is required for grade")
		return errors.New("student id is required")
	}

	if grade.CourseID == 0 {
		s.logger.Warn("Course ID is required for grade")
		return errors.New("course id is required")
	}

	if grade.GradedBy == 0 {
		s.logger.Warn("Graded by (teacher ID) is required")
		return errors.New("teacher id is required")
	}

	if grade.Score < 0 || grade.Score > grade.MaxScore {
		s.logger.WithField("score", grade.Score).WithField("max_score", grade.MaxScore).Warn("Invalid score")
		return errors.New("score must be between 0 and max_score")
	}

	if grade.GradedAt.IsZero() {
		grade.GradedAt = time.Now()
	}

	err := s.gradeRepo.Create(grade)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", grade.StudentID).WithField("course_id", grade.CourseID).Error("Failed to record grade")
		return errors.New("failed to record grade")
	}

	s.logger.WithField("student_id", grade.StudentID).WithField("course_id", grade.CourseID).WithField("score", grade.Score).Info("Grade recorded")
	return nil
}

func (s *gradeService) GetGradeByID(id uint) (*models.Grade, error) {
	grade, err := s.gradeRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Grade not found")
		return nil, errors.New("grade not found")
	}

	return grade, nil
}

func (s *gradeService) GetStudentGrades(studentID uint, page, limit int) ([]models.Grade, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	grades, total, err := s.gradeRepo.FindByStudentID(studentID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).Error("Failed to fetch grades")
		return nil, 0, err
	}

	return grades, total, nil
}

func (s *gradeService) GetCourseGrades(courseID uint, page, limit int) ([]models.Grade, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	grades, total, err := s.gradeRepo.FindByCourseID(courseID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", courseID).Error("Failed to fetch grades")
		return nil, 0, err
	}

	return grades, total, nil
}

func (s *gradeService) GetStudentCourseGrade(studentID, courseID uint) (*models.Grade, error) {
	grade, err := s.gradeRepo.FindByStudentAndCourse(studentID, courseID)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).WithField("course_id", courseID).Warn("Grade not found")
		return nil, errors.New("grade not found")
	}

	return grade, nil
}

func (s *gradeService) GetAllGrades(page, limit int) ([]models.Grade, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	grades, total, err := s.gradeRepo.FindAll(page, limit)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch grades")
		return nil, 0, err
	}

	return grades, total, nil
}

func (s *gradeService) UpdateGrade(grade *models.Grade) error {
	if grade.ID == 0 {
		return errors.New("grade id is required")
	}

	if grade.Score < 0 || grade.Score > grade.MaxScore {
		s.logger.WithField("score", grade.Score).Warn("Invalid score")
		return errors.New("score must be between 0 and max_score")
	}

	grade.GradedAt = time.Now()

	err := s.gradeRepo.Update(grade)
	if err != nil {
		s.logger.WithError(err).WithField("id", grade.ID).Error("Failed to update grade")
		return errors.New("failed to update grade")
	}

	s.logger.WithField("id", grade.ID).Info("Grade updated")
	return nil
}

func (s *gradeService) DeleteGrade(id uint) error {
	err := s.gradeRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete grade")
		return errors.New("failed to delete grade")
	}

	s.logger.WithField("id", id).Info("Grade deleted")
	return nil
}

func (s *gradeService) GetTeacherGrades(teacherID uint, page, limit int) ([]models.Grade, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	grades, total, err := s.gradeRepo.FindByTeacherID(teacherID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("teacher_id", teacherID).Error("Failed to fetch grades")
		return nil, 0, err
	}

	return grades, total, nil
}

func (s *gradeService) GetGradesInRange(startDate, endDate time.Time) ([]models.Grade, error) {
	grades, err := s.gradeRepo.FindGradesInDateRange(startDate, endDate)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch grades in range")
		return nil, err
	}

	return grades, nil
}

func (s *gradeService) CalculateAverageGrade(studentID uint) (float64, error) {
	grades, _, err := s.gradeRepo.FindByStudentID(studentID, 1, 1000)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).Error("Failed to calculate average")
		return 0, err
	}

	if len(grades) == 0 {
		return 0, nil
	}

	var total float64
	for _, g := range grades {
		total += g.Score
	}

	return total / float64(len(grades)), nil
}
