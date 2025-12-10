package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
)

type EnrollmentService interface {
	EnrollStudent(enrollment *models.Enrollment) error
	GetEnrollmentByID(id uint) (*models.Enrollment, error)
	GetStudentEnrollments(studentID uint, page, limit int) ([]models.Enrollment, int64, error)
	GetCourseEnrollments(courseID uint, page, limit int) ([]models.Enrollment, int64, error)
	GetAllEnrollments(page, limit int) ([]models.Enrollment, int64, error)
	UpdateEnrollmentStatus(id uint, status string) error
	RemoveEnrollment(id uint) error
	CheckEnrollment(studentID, courseID uint) (bool, error)
	GetCourseEnrollmentCount(courseID uint) (int64, error)
}

type enrollmentService struct {
	enrollmentRepo repository.EnrollmentRepository
	logger         *logrus.Logger
}

func NewEnrollmentService(enrollmentRepo repository.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{
		enrollmentRepo: enrollmentRepo,
		logger:         logger.GetLogger(),
	}
}

func (s *enrollmentService) EnrollStudent(enrollment *models.Enrollment) error {
	if enrollment.StudentID == 0 {
		s.logger.Warn("Student ID is required for enrollment")
		return errors.New("student id is required")
	}

	if enrollment.CourseID == 0 {
		s.logger.Warn("Course ID is required for enrollment")
		return errors.New("course id is required")
	}

	// Check if already enrolled
	existing, _ := s.enrollmentRepo.FindByStudentAndCourse(enrollment.StudentID, enrollment.CourseID)
	if existing != nil {
		s.logger.WithField("student_id", enrollment.StudentID).WithField("course_id", enrollment.CourseID).Warn("Student already enrolled in course")
		return errors.New("student already enrolled in this course")
	}

	if enrollment.EnrolledAt.IsZero() {
		enrollment.EnrolledAt = time.Now()
	}

	if enrollment.Status == "" {
		enrollment.Status = "active"
	}

	err := s.enrollmentRepo.Create(enrollment)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", enrollment.StudentID).WithField("course_id", enrollment.CourseID).Error("Failed to enroll student")
		return errors.New("failed to enroll student")
	}

	s.logger.WithField("student_id", enrollment.StudentID).WithField("course_id", enrollment.CourseID).Info("Student enrolled successfully")
	return nil
}

func (s *enrollmentService) GetEnrollmentByID(id uint) (*models.Enrollment, error) {
	enrollment, err := s.enrollmentRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Enrollment not found")
		return nil, errors.New("enrollment not found")
	}

	return enrollment, nil
}

func (s *enrollmentService) GetStudentEnrollments(studentID uint, page, limit int) ([]models.Enrollment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	enrollments, total, err := s.enrollmentRepo.FindByStudentID(studentID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).Error("Failed to fetch enrollments")
		return nil, 0, err
	}

	return enrollments, total, nil
}

func (s *enrollmentService) GetCourseEnrollments(courseID uint, page, limit int) ([]models.Enrollment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	enrollments, total, err := s.enrollmentRepo.FindByCourseID(courseID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", courseID).Error("Failed to fetch enrollments")
		return nil, 0, err
	}

	return enrollments, total, nil
}

func (s *enrollmentService) GetAllEnrollments(page, limit int) ([]models.Enrollment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	enrollments, total, err := s.enrollmentRepo.FindAll(page, limit)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch enrollments")
		return nil, 0, err
	}

	return enrollments, total, nil
}

func (s *enrollmentService) UpdateEnrollmentStatus(id uint, status string) error {
	enrollment, err := s.enrollmentRepo.FindByID(id)
	if err != nil {
		return errors.New("enrollment not found")
	}

	enrollment.Status = status
	err = s.enrollmentRepo.Update(enrollment)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to update enrollment")
		return errors.New("failed to update enrollment")
	}

	s.logger.WithField("id", id).WithField("status", status).Info("Enrollment status updated")
	return nil
}

func (s *enrollmentService) RemoveEnrollment(id uint) error {
	err := s.enrollmentRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to remove enrollment")
		return errors.New("failed to remove enrollment")
	}

	s.logger.WithField("id", id).Info("Enrollment removed")
	return nil
}

func (s *enrollmentService) CheckEnrollment(studentID, courseID uint) (bool, error) {
	enrollment, err := s.enrollmentRepo.FindByStudentAndCourse(studentID, courseID)
	if err != nil {
		return false, nil
	}

	return enrollment != nil, nil
}

func (s *enrollmentService) GetCourseEnrollmentCount(courseID uint) (int64, error) {
	return s.enrollmentRepo.CountByCourseID(courseID)
}
