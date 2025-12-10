package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
)

type StudentService interface {
	CreateStudent(student *models.Student) error
	GetStudentByID(id uint) (*models.Student, error)
	GetStudentByUserID(userID uint) (*models.Student, error)
	GetAllStudents(page, limit int) ([]models.Student, int64, error)
	GetStudentsByGradeLevel(gradeLevel string, page, limit int) ([]models.Student, int64, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type studentService struct {
	studentRepo repository.StudentRepository
	logger      *logrus.Logger
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{
		studentRepo: studentRepo,
		logger:      logger.GetLogger(),
	}
}

func (s *studentService) CreateStudent(student *models.Student) error {
	if student.UserID == 0 {
		s.logger.Warn("User ID is required for student")
		return errors.New("user id is required")
	}

	if student.StudentID == "" {
		s.logger.Warn("Student ID is required")
		return errors.New("student id is required")
	}

	if student.EnrollmentDate.IsZero() {
		student.EnrollmentDate = time.Now()
	}

	err := s.studentRepo.Create(student)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", student.StudentID).Error("Failed to create student")
		return errors.New("failed to create student")
	}

	s.logger.WithField("student_id", student.StudentID).WithField("user_id", student.UserID).Info("Student created successfully")
	return nil
}

func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
	student, err := s.studentRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Student not found")
		return nil, errors.New("student not found")
	}

	return student, nil
}

func (s *studentService) GetStudentByUserID(userID uint) (*models.Student, error) {
	student, err := s.studentRepo.FindByUserID(userID)
	if err != nil {
		s.logger.WithError(err).WithField("user_id", userID).Warn("Student not found")
		return nil, errors.New("student not found")
	}

	return student, nil
}

func (s *studentService) GetAllStudents(page, limit int) ([]models.Student, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	students, total, err := s.studentRepo.FindAll(page, limit)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch students")
		return nil, 0, err
	}

	return students, total, nil
}

func (s *studentService) GetStudentsByGradeLevel(gradeLevel string, page, limit int) ([]models.Student, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	students, total, err := s.studentRepo.FindByGradeLevel(gradeLevel, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("grade_level", gradeLevel).Error("Failed to fetch students")
		return nil, 0, err
	}

	return students, total, nil
}

func (s *studentService) UpdateStudent(student *models.Student) error {
	if student.ID == 0 {
		return errors.New("student id is required")
	}

	err := s.studentRepo.Update(student)
	if err != nil {
		s.logger.WithError(err).WithField("id", student.ID).Error("Failed to update student")
		return errors.New("failed to update student")
	}

	s.logger.WithField("id", student.ID).Info("Student updated successfully")
	return nil
}

func (s *studentService) DeleteStudent(id uint) error {
	err := s.studentRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete student")
		return errors.New("failed to delete student")
	}

	s.logger.WithField("id", id).Info("Student deleted successfully")
	return nil
}

// Student service placeholder. Implement student-related service logic here as needed.
