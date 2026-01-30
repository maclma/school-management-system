package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"time"

	"github.com/sirupsen/logrus"
)

type AssignmentService interface {
	CreateAssignment(assignment *models.Assignment) error
	GetAssignmentByID(id uint) (*models.Assignment, error)
	GetAssignmentsByCourse(courseID uint) ([]models.Assignment, error)
	GetAssignmentsByTeacher(teacherID uint) ([]models.Assignment, error)
	UpdateAssignment(assignment *models.Assignment) error
	DeleteAssignment(id uint) error
}

type AssignmentSubmissionService interface {
	SubmitAssignment(submission *models.AssignmentSubmission) error
	GetSubmissionByID(id uint) (*models.AssignmentSubmission, error)
	GetSubmissionsByAssignment(assignmentID uint) ([]models.AssignmentSubmission, error)
	GetSubmissionsByStudent(studentID uint) ([]models.AssignmentSubmission, error)
	GetSubmissionByAssignmentAndStudent(assignmentID, studentID uint) (*models.AssignmentSubmission, error)
	GradeSubmission(submissionID uint, score float64, feedback string) error
	UpdateSubmission(submission *models.AssignmentSubmission) error
	DeleteSubmission(id uint) error
}

type assignmentService struct {
	assignmentRepo repository.AssignmentRepository
	logger         *logrus.Logger
}

type assignmentSubmissionService struct {
	submissionRepo repository.AssignmentSubmissionRepository
	logger         *logrus.Logger
}

func NewAssignmentService(assignmentRepo repository.AssignmentRepository) AssignmentService {
	return &assignmentService{
		assignmentRepo: assignmentRepo,
		logger:         logger.GetLogger(),
	}
}

func NewAssignmentSubmissionService(submissionRepo repository.AssignmentSubmissionRepository) AssignmentSubmissionService {
	return &assignmentSubmissionService{
		submissionRepo: submissionRepo,
		logger:         logger.GetLogger(),
	}
}

// Assignment methods
func (s *assignmentService) CreateAssignment(assignment *models.Assignment) error {
	if assignment.Title == "" {
		s.logger.Warn("Assignment title is required")
		return errors.New("assignment title is required")
	}

	if assignment.CourseID == 0 {
		s.logger.Warn("Course ID is required")
		return errors.New("course id is required")
	}

	if assignment.CreatedBy == 0 {
		s.logger.Warn("Created by is required")
		return errors.New("created by is required")
	}

	assignment.CreatedAt = time.Now()

	err := s.assignmentRepo.Create(assignment)
	if err != nil {
		s.logger.WithError(err).WithField("title", assignment.Title).Error("Failed to create assignment")
		return errors.New("failed to create assignment")
	}

	s.logger.WithField("title", assignment.Title).WithField("course_id", assignment.CourseID).Info("Assignment created successfully")
	return nil
}

func (s *assignmentService) GetAssignmentByID(id uint) (*models.Assignment, error) {
	assignment, err := s.assignmentRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Assignment not found")
		return nil, errors.New("assignment not found")
	}

	return assignment, nil
}

func (s *assignmentService) GetAssignmentsByCourse(courseID uint) ([]models.Assignment, error) {
	assignments, err := s.assignmentRepo.FindByCourseID(courseID)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", courseID).Error("Failed to fetch assignments")
		return nil, err
	}

	return assignments, nil
}

func (s *assignmentService) GetAssignmentsByTeacher(teacherID uint) ([]models.Assignment, error) {
	assignments, err := s.assignmentRepo.FindAllByTeacherID(teacherID)
	if err != nil {
		s.logger.WithError(err).WithField("teacher_id", teacherID).Error("Failed to fetch teacher assignments")
		return nil, err
	}

	return assignments, nil
}

func (s *assignmentService) UpdateAssignment(assignment *models.Assignment) error {
	if assignment.ID == 0 {
		return errors.New("assignment id is required")
	}

	err := s.assignmentRepo.Update(assignment)
	if err != nil {
		s.logger.WithError(err).WithField("id", assignment.ID).Error("Failed to update assignment")
		return errors.New("failed to update assignment")
	}

	s.logger.WithField("id", assignment.ID).Info("Assignment updated successfully")
	return nil
}

func (s *assignmentService) DeleteAssignment(id uint) error {
	err := s.assignmentRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete assignment")
		return errors.New("failed to delete assignment")
	}

	s.logger.WithField("id", id).Info("Assignment deleted successfully")
	return nil
}

// AssignmentSubmission methods
func (s *assignmentSubmissionService) SubmitAssignment(submission *models.AssignmentSubmission) error {
	if submission.AssignmentID == 0 {
		s.logger.Warn("Assignment ID is required")
		return errors.New("assignment id is required")
	}

	if submission.StudentID == 0 {
		s.logger.Warn("Student ID is required")
		return errors.New("student id is required")
	}

	// Check if submission already exists
	existing, _ := s.submissionRepo.FindByAssignmentAndStudent(submission.AssignmentID, submission.StudentID)
	if existing != nil {
		s.logger.WithFields(logrus.Fields{
			"assignment_id": submission.AssignmentID,
			"student_id":    submission.StudentID,
		}).Warn("Submission already exists")
		return errors.New("submission already exists for this assignment")
	}

	now := time.Now()
	submission.SubmittedAt = &now
	submission.Status = "submitted"

	err := s.submissionRepo.Create(submission)
	if err != nil {
		s.logger.WithError(err).WithFields(logrus.Fields{
			"assignment_id": submission.AssignmentID,
			"student_id":    submission.StudentID,
		}).Error("Failed to submit assignment")
		return errors.New("failed to submit assignment")
	}

	s.logger.WithFields(logrus.Fields{
		"assignment_id": submission.AssignmentID,
		"student_id":    submission.StudentID,
	}).Info("Assignment submitted successfully")
	return nil
}

func (s *assignmentSubmissionService) GetSubmissionByID(id uint) (*models.AssignmentSubmission, error) {
	submission, err := s.submissionRepo.FindByID(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Warn("Submission not found")
		return nil, errors.New("submission not found")
	}

	return submission, nil
}

func (s *assignmentSubmissionService) GetSubmissionsByAssignment(assignmentID uint) ([]models.AssignmentSubmission, error) {
	submissions, err := s.submissionRepo.FindByAssignmentID(assignmentID)
	if err != nil {
		s.logger.WithError(err).WithField("assignment_id", assignmentID).Error("Failed to fetch submissions")
		return nil, err
	}

	return submissions, nil
}

func (s *assignmentSubmissionService) GetSubmissionsByStudent(studentID uint) ([]models.AssignmentSubmission, error) {
	submissions, err := s.submissionRepo.FindByStudentID(studentID)
	if err != nil {
		s.logger.WithError(err).WithField("student_id", studentID).Error("Failed to fetch student submissions")
		return nil, err
	}

	return submissions, nil
}

func (s *assignmentSubmissionService) GetSubmissionByAssignmentAndStudent(assignmentID, studentID uint) (*models.AssignmentSubmission, error) {
	submission, err := s.submissionRepo.FindByAssignmentAndStudent(assignmentID, studentID)
	if err != nil {
		s.logger.WithError(err).WithFields(logrus.Fields{
			"assignment_id": assignmentID,
			"student_id":    studentID,
		}).Warn("Submission not found")
		return nil, errors.New("submission not found")
	}

	return submission, nil
}

func (s *assignmentSubmissionService) GradeSubmission(submissionID uint, score float64, feedback string) error {
	submission, err := s.submissionRepo.FindByID(submissionID)
	if err != nil {
		s.logger.WithError(err).WithField("id", submissionID).Warn("Submission not found")
		return errors.New("submission not found")
	}

	submission.Score = &score
	submission.Feedback = feedback
	submission.Status = "graded"

	err = s.submissionRepo.Update(submission)
	if err != nil {
		s.logger.WithError(err).WithField("id", submissionID).Error("Failed to grade submission")
		return errors.New("failed to grade submission")
	}

	s.logger.WithFields(logrus.Fields{
		"id":    submissionID,
		"score": score,
	}).Info("Submission graded successfully")
	return nil
}

func (s *assignmentSubmissionService) UpdateSubmission(submission *models.AssignmentSubmission) error {
	if submission.ID == 0 {
		return errors.New("submission id is required")
	}

	err := s.submissionRepo.Update(submission)
	if err != nil {
		s.logger.WithError(err).WithField("id", submission.ID).Error("Failed to update submission")
		return errors.New("failed to update submission")
	}

	s.logger.WithField("id", submission.ID).Info("Submission updated successfully")
	return nil
}

func (s *assignmentSubmissionService) DeleteSubmission(id uint) error {
	err := s.submissionRepo.Delete(id)
	if err != nil {
		s.logger.WithError(err).WithField("id", id).Error("Failed to delete submission")
		return errors.New("failed to delete submission")
	}

	s.logger.WithField("id", id).Info("Submission deleted successfully")
	return nil
}
