package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
)

type AttendanceService interface {
	RecordAttendance(attendance *models.Attendance) error
	GetAttendanceByID(id uint) (*models.Attendance, error)
	GetStudentAttendance(studentID uint, page, limit int) ([]models.Attendance, int64, error)
	GetCourseAttendance(courseID uint, page, limit int) ([]models.Attendance, int64, error)
	GetStudentCourseAttendance(studentID, courseID uint, page, limit int) ([]models.Attendance, int64, error)
	GetAllAttendance(page, limit int) ([]models.Attendance, int64, error)
	UpdateAttendance(attendance *models.Attendance) error
	DeleteAttendance(id uint) error
	GetAttendanceInRange(startDate, endDate time.Time) ([]models.Attendance, error)
	GetStudentAttendanceStats(studentID, courseID uint) (present, absent, late int64, error error)
	CalculateAttendancePercentage(studentID, courseID uint) (float64, error)
}

type attendanceService struct {
	attendanceRepo repository.AttendanceRepository
	logger         *logrus.Logger
}

func NewAttendanceService(attendanceRepo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{
		attendanceRepo: attendanceRepo,
		logger:         logger.GetLogger(),
	}
}

func (s *attendanceService) RecordAttendance(attendance *models.Attendance) error {
	if attendance.StudentID == 0 {
		s.logger.Warn("Student ID is required for attendance")
		return errors.New("student id is required")
	}

	if attendance.CourseID == 0 {
		s.logger.Warn("Course ID is required for attendance")
		return errors.New("course id is required")
	}

	validStatuses := map[string]bool{"present": true, "absent": true, "late": true, "excused": true}
	if !validStatuses[attendance.Status] {
		s.logger.WithField("status", attendance.Status).Warn("Invalid attendance status")
		return errors.New("invalid attendance status")
	}

	if attendance.Date.IsZero() {
		attendance.Date = time.Now()
	}

	err := s.attendanceRepo.Create(attendance)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", attendance.StudentID).WithField("course_id", attendance.CourseID).Error("Failed to record attendance")
		return errors.New("failed to record attendance")
	}

	s.logger.WithField("student_id", attendance.StudentID).WithField("course_id", attendance.CourseID).WithField("status", attendance.Status).Info("Attendance recorded")
	return nil
}

func (s *attendanceService) GetAttendanceByID(id uint) (*models.Attendance, error) {
	attendance, err := s.attendanceRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Attendance record not found")
		return nil, errors.New("attendance not found")
	}

	return attendance, nil
}

func (s *attendanceService) GetStudentAttendance(studentID uint, page, limit int) ([]models.Attendance, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	attendances, total, err := s.attendanceRepo.FindByStudentID(studentID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).Error("Failed to fetch attendance")
		return nil, 0, err
	}

	return attendances, total, nil
}

func (s *attendanceService) GetCourseAttendance(courseID uint, page, limit int) ([]models.Attendance, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	attendances, total, err := s.attendanceRepo.FindByCourseID(courseID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", courseID).Error("Failed to fetch attendance")
		return nil, 0, err
	}

	return attendances, total, nil
}

func (s *attendanceService) GetStudentCourseAttendance(studentID, courseID uint, page, limit int) ([]models.Attendance, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	attendances, total, err := s.attendanceRepo.FindByStudentAndCourse(studentID, courseID, page, limit)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).WithField("course_id", courseID).Error("Failed to fetch attendance")
		return nil, 0, err
	}

	return attendances, total, nil
}

func (s *attendanceService) GetAllAttendance(page, limit int) ([]models.Attendance, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	attendances, total, err := s.attendanceRepo.FindAll(page, limit)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch attendance")
		return nil, 0, err
	}

	return attendances, total, nil
}

func (s *attendanceService) UpdateAttendance(attendance *models.Attendance) error {
	if attendance.ID == 0 {
		return errors.New("attendance id is required")
	}

	validStatuses := map[string]bool{"present": true, "absent": true, "late": true, "excused": true}
	if !validStatuses[attendance.Status] {
		return errors.New("invalid attendance status")
	}

	err := s.attendanceRepo.Update(attendance)
	if err != nil {
		s.logger.WithError(err).WithField("id", attendance.ID).Error("Failed to update attendance")
		return errors.New("failed to update attendance")
	}

	s.logger.WithField("id", attendance.ID).Info("Attendance updated")
	return nil
}

func (s *attendanceService) DeleteAttendance(id uint) error {
	err := s.attendanceRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete attendance")
		return errors.New("failed to delete attendance")
	}

	s.logger.WithField("id", id).Info("Attendance deleted")
	return nil
}

func (s *attendanceService) GetAttendanceInRange(startDate, endDate time.Time) ([]models.Attendance, error) {
	attendances, err := s.attendanceRepo.FindByDateRange(startDate, endDate)
	if err != nil {
		s.logger.WithError(err).Error("Failed to fetch attendance in range")
		return nil, err
	}

	return attendances, nil
}

func (s *attendanceService) GetStudentAttendanceStats(studentID, courseID uint) (present, absent, late int64, error error) {
	return s.attendanceRepo.CountAttendanceByStudent(studentID, courseID)
}

func (s *attendanceService) CalculateAttendancePercentage(studentID, courseID uint) (float64, error) {
	present, absent, late, err := s.attendanceRepo.CountAttendanceByStudent(studentID, courseID)
	if err != nil {
		s.logger.WithError(err).Error("Failed to calculate attendance percentage")
		return 0, err
	}

	total := present + absent + late
	if total == 0 {
		return 0, nil
	}

	return (float64(present) / float64(total)) * 100, nil
}
