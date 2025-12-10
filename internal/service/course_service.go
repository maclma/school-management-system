package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"

	"github.com/sirupsen/logrus"
)

type CourseService interface {
	CreateCourse(course *models.Course) error
	GetCourseByID(id uint) (*models.Course, error)
	GetAllCourses(page, limit int) ([]models.Course, int64, error)
	GetCoursesByDepartment(department string, page, limit int) ([]models.Course, int64, error)
	UpdateCourse(course *models.Course) error
	DeleteCourse(id uint) error
	GetCoursesByTeacher(teacherID uint) ([]models.Course, error)
}

type courseService struct {
	courseRepo repository.CourseRepository
	logger     *logrus.Logger
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
		logger:     logger.GetLogger(),
	}
}

func (s *courseService) CreateCourse(course *models.Course) error {
	if course.Name == "" {
		s.logger.Warn("Course name is required")
		return errors.New("course name is required")
	}

	if course.CourseCode == "" {
		s.logger.Warn("Course code is required")
		return errors.New("course code is required")
	}

	// Check if course code already exists
	existing, _ := s.courseRepo.FindByCourseCode(course.CourseCode)
	if existing != nil {
		s.logger.WithField("code", course.CourseCode).Warn("Course code already exists")
		return errors.New("course code already exists")
	}

	err := s.courseRepo.Create(course)
	if err != nil {
		s.logger.WithError(err).WithField("code", course.CourseCode).Error("Failed to create course")
		return errors.New("failed to create course")
	}

	s.logger.WithField("code", course.CourseCode).WithField("name", course.Name).Info("Course created successfully")
	return nil
}

func (s *courseService) GetCourseByID(id uint) (*models.Course, error) {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Course not found")
		return nil, errors.New("course not found")
	}

	return course, nil
}

func (s *courseService) GetAllCourses(page, limit int) ([]models.Course, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	courses, total, err := s.courseRepo.FindAll(page, limit)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch courses")
		return nil, 0, err
	}

	return courses, total, nil
}

func (s *courseService) GetCoursesByDepartment(department string, page, limit int) ([]models.Course, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	courses, total, err := s.courseRepo.FindByDepartment(department, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("department", department).Error("Failed to fetch courses")
		return nil, 0, err
	}

	return courses, total, nil
}

func (s *courseService) UpdateCourse(course *models.Course) error {
	if course.ID == 0 {
		return errors.New("course id is required")
	}

	err := s.courseRepo.Update(course)
	if err != nil {
		s.logger.WithError(err).WithField("id", course.ID).Error("Failed to update course")
		return errors.New("failed to update course")
	}

	s.logger.WithField("id", course.ID).Info("Course updated successfully")
	return nil
}

func (s *courseService) DeleteCourse(id uint) error {
	err := s.courseRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete course")
		return errors.New("failed to delete course")
	}

	s.logger.WithField("id", id).Info("Course deleted successfully")
	return nil
}

func (s *courseService) GetCoursesByTeacher(teacherID uint) ([]models.Course, error) {
	courses, err := s.courseRepo.FindByTeacherID(teacherID)
	if err != nil {
		s.logger.WithError(err).WithField("teacher_id", teacherID).Error("Failed to fetch teacher courses")
		return nil, err
	}

	return courses, nil
}

// Course service placeholder. Implement course-related service logic here as needed.
