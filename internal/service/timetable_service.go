package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type TimeTableService interface {
	Create(timetable *models.TimeTable) error
	GetByID(id uint) (*models.TimeTable, error)
	GetByCourseID(courseID uint) ([]models.TimeTable, error)
	GetByTeacherID(teacherID uint) ([]models.TimeTable, error)
	GetByDayOfWeek(dayOfWeek string) ([]models.TimeTable, error)
	GetAll() ([]models.TimeTable, error)
	Update(timetable *models.TimeTable) error
	Delete(id uint) error
}

type timetableService struct {
	repo repository.TimeTableRepository
}

func NewTimeTableService(repo repository.TimeTableRepository) TimeTableService {
	return &timetableService{repo: repo}
}

func (s *timetableService) Create(timetable *models.TimeTable) error {
	return s.repo.Create(timetable)
}

func (s *timetableService) GetByID(id uint) (*models.TimeTable, error) {
	return s.repo.FindByID(id)
}

func (s *timetableService) GetByCourseID(courseID uint) ([]models.TimeTable, error) {
	return s.repo.FindByCourseID(courseID)
}

func (s *timetableService) GetByTeacherID(teacherID uint) ([]models.TimeTable, error) {
	return s.repo.FindByTeacherID(teacherID)
}

func (s *timetableService) GetByDayOfWeek(dayOfWeek string) ([]models.TimeTable, error) {
	return s.repo.FindByDayOfWeek(dayOfWeek)
}

func (s *timetableService) GetAll() ([]models.TimeTable, error) {
	return s.repo.FindAll()
}

func (s *timetableService) Update(timetable *models.TimeTable) error {
	return s.repo.Update(timetable)
}

func (s *timetableService) Delete(id uint) error {
	return s.repo.Delete(id)
}
