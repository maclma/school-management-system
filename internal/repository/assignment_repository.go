package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type AssignmentRepository interface {
	Create(assignment *models.Assignment) error
	FindByID(id uint) (*models.Assignment, error)
	FindByCourseID(courseID uint) ([]models.Assignment, error)
	Update(assignment *models.Assignment) error
	Delete(id uint) error
	FindAllByTeacherID(teacherID uint) ([]models.Assignment, error)
}

type AssignmentSubmissionRepository interface {
	Create(submission *models.AssignmentSubmission) error
	FindByID(id uint) (*models.AssignmentSubmission, error)
	FindByAssignmentID(assignmentID uint) ([]models.AssignmentSubmission, error)
	FindByStudentID(studentID uint) ([]models.AssignmentSubmission, error)
	FindByAssignmentAndStudent(assignmentID, studentID uint) (*models.AssignmentSubmission, error)
	Update(submission *models.AssignmentSubmission) error
	Delete(id uint) error
}

type assignmentRepository struct {
	db *gorm.DB
}

type assignmentSubmissionRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository() AssignmentRepository {
	return &assignmentRepository{db: database.DB}
}

func NewAssignmentSubmissionRepository() AssignmentSubmissionRepository {
	return &assignmentSubmissionRepository{db: database.DB}
}

// Assignment methods
func (r *assignmentRepository) Create(assignment *models.Assignment) error {
	return r.db.Create(assignment).Error
}

func (r *assignmentRepository) FindByID(id uint) (*models.Assignment, error) {
	var assignment models.Assignment
	err := r.db.Preload("Course").Preload("Teacher").Preload("Submissions").Preload("Submissions.Student").First(&assignment, id).Error
	return &assignment, err
}

func (r *assignmentRepository) FindByCourseID(courseID uint) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.db.Where("course_id = ?", courseID).Preload("Course").Preload("Teacher").Order("due_date ASC").Find(&assignments).Error
	return assignments, err
}

func (r *assignmentRepository) Update(assignment *models.Assignment) error {
	return r.db.Save(assignment).Error
}

func (r *assignmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Assignment{}, id).Error
}

func (r *assignmentRepository) FindAllByTeacherID(teacherID uint) ([]models.Assignment, error) {
	var assignments []models.Assignment
	err := r.db.Where("created_by = ?", teacherID).Preload("Course").Preload("Teacher").Order("due_date ASC").Find(&assignments).Error
	return assignments, err
}

// AssignmentSubmission methods
func (r *assignmentSubmissionRepository) Create(submission *models.AssignmentSubmission) error {
	return r.db.Create(submission).Error
}

func (r *assignmentSubmissionRepository) FindByID(id uint) (*models.AssignmentSubmission, error) {
	var submission models.AssignmentSubmission
	err := r.db.Preload("Assignment").Preload("Student").First(&submission, id).Error
	return &submission, err
}

func (r *assignmentSubmissionRepository) FindByAssignmentID(assignmentID uint) ([]models.AssignmentSubmission, error) {
	var submissions []models.AssignmentSubmission
	err := r.db.Where("assignment_id = ?", assignmentID).Preload("Student").Order("submitted_at DESC").Find(&submissions).Error
	return submissions, err
}

func (r *assignmentSubmissionRepository) FindByStudentID(studentID uint) ([]models.AssignmentSubmission, error) {
	var submissions []models.AssignmentSubmission
	err := r.db.Where("student_id = ?", studentID).Preload("Assignment").Preload("Assignment.Course").Order("submitted_at DESC").Find(&submissions).Error
	return submissions, err
}

func (r *assignmentSubmissionRepository) FindByAssignmentAndStudent(assignmentID, studentID uint) (*models.AssignmentSubmission, error) {
	var submission models.AssignmentSubmission
	err := r.db.Where("assignment_id = ? AND student_id = ?", assignmentID, studentID).Preload("Assignment").Preload("Student").First(&submission).Error
	return &submission, err
}

func (r *assignmentSubmissionRepository) Update(submission *models.AssignmentSubmission) error {
	return r.db.Save(submission).Error
}

func (r *assignmentSubmissionRepository) Delete(id uint) error {
	return r.db.Delete(&models.AssignmentSubmission{}, id).Error
}
