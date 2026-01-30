package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TeacherService interface {
	CreateTeacher(teacher *models.Teacher) error
	GetTeacherByID(id uint) (*models.Teacher, error)
	GetTeacherByUserID(userID uint) (*models.Teacher, error)
	GetTeacherByTeacherID(teacherID string) (*models.Teacher, error)
	GetAllTeachers(page, limit int) ([]models.Teacher, int64, error)
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(id uint) error
	GetTeacherCourses(teacherID uint, page, limit int) ([]models.Course, int64, error)
	GetTeachersByDepartment(department string, page, limit int) ([]models.Teacher, int64, error)
}

type teacherService struct {
	teacherRepo repository.TeacherRepository
	logger      *logrus.Logger
}

func NewTeacherService(teacherRepo repository.TeacherRepository) TeacherService {
	return &teacherService{
		teacherRepo: teacherRepo,
		logger:      logger.GetLogger(),
	}
}

func (s *teacherService) CreateTeacher(teacher *models.Teacher) error {
	s.logger.WithField("teacher_id", teacher.TeacherID).Info("Creating teacher")

	// Check if teacher with this teacher_id already exists
	_, err := s.teacherRepo.GetByTeacherID(teacher.TeacherID)
	if err == nil {
		return errors.New("teacher with this teacher ID already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return s.teacherRepo.Create(teacher)
}

func (s *teacherService) GetTeacherByID(id uint) (*models.Teacher, error) {
	s.logger.WithField("id", id).Info("Getting teacher by ID")
	return s.teacherRepo.GetByID(id)
}

func (s *teacherService) GetTeacherByUserID(userID uint) (*models.Teacher, error) {
	s.logger.WithField("user_id", userID).Info("Getting teacher by user ID")
	return s.teacherRepo.GetByUserID(userID)
}

func (s *teacherService) GetTeacherByTeacherID(teacherID string) (*models.Teacher, error) {
	s.logger.WithField("teacher_id", teacherID).Info("Getting teacher by teacher ID")
	return s.teacherRepo.GetByTeacherID(teacherID)
}

func (s *teacherService) GetAllTeachers(page, limit int) ([]models.Teacher, int64, error) {
	s.logger.WithFields(logrus.Fields{"page": page, "limit": limit}).Info("Getting all teachers")
	return s.teacherRepo.GetAll(page, limit)
}

func (s *teacherService) UpdateTeacher(teacher *models.Teacher) error {
	s.logger.WithField("id", teacher.ID).Info("Updating teacher")
	return s.teacherRepo.Update(teacher)
}

func (s *teacherService) DeleteTeacher(id uint) error {
	s.logger.WithField("id", id).Info("Deleting teacher")
	return s.teacherRepo.Delete(id)
}

func (s *teacherService) GetTeacherCourses(teacherID uint, page, limit int) ([]models.Course, int64, error) {
	s.logger.WithFields(logrus.Fields{"teacher_id": teacherID, "page": page, "limit": limit}).Info("Getting teacher courses")
	return s.teacherRepo.GetTeacherCourses(teacherID, page, limit)
}

func (s *teacherService) GetTeachersByDepartment(department string, page, limit int) ([]models.Teacher, int64, error) {
	s.logger.WithFields(logrus.Fields{"department": department, "page": page, "limit": limit}).Info("Getting teachers by department")
	return s.teacherRepo.GetByDepartment(department, page, limit)
}
