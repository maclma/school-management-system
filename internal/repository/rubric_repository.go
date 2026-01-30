package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"
)

// AssignmentRubricRepository handles rubric data access
type AssignmentRubricRepository struct{}

// NewAssignmentRubricRepository creates a new repository
func NewAssignmentRubricRepository() *AssignmentRubricRepository {
	return &AssignmentRubricRepository{}
}

// Create creates a new rubric
func (r *AssignmentRubricRepository) Create(rubric *models.AssignmentRubric) error {
	db := database.DB
	return db.Create(rubric).Error
}

// GetByID retrieves a rubric by ID
func (r *AssignmentRubricRepository) GetByID(id uint) (*models.AssignmentRubric, error) {
	db := database.DB
	var rubric models.AssignmentRubric
	if err := db.First(&rubric, id).Error; err != nil {
		return nil, err
	}
	return &rubric, nil
}

// GetByAssignmentID retrieves rubrics for an assignment
func (r *AssignmentRubricRepository) GetByAssignmentID(assignmentID uint) ([]models.AssignmentRubric, error) {
	db := database.DB
	var rubrics []models.AssignmentRubric
	if err := db.Where("assignment_id = ?", assignmentID).Find(&rubrics).Error; err != nil {
		return nil, err
	}
	return rubrics, nil
}

// Update updates a rubric
func (r *AssignmentRubricRepository) Update(rubric *models.AssignmentRubric) error {
	db := database.DB
	return db.Save(rubric).Error
}

// Delete deletes a rubric
func (r *AssignmentRubricRepository) Delete(id uint) error {
	db := database.DB
	return db.Delete(&models.AssignmentRubric{}, id).Error
}

// RubricScoreRepository handles rubric score access
type RubricScoreRepository struct{}

// NewRubricScoreRepository creates a new repository
func NewRubricScoreRepository() *RubricScoreRepository {
	return &RubricScoreRepository{}
}

// Create creates a new rubric score
func (r *RubricScoreRepository) Create(score *models.RubricScore) error {
	db := database.DB
	return db.Create(score).Error
}

// GetBySubmissionAndRubric retrieves score for a submission
func (r *RubricScoreRepository) GetBySubmissionAndRubric(submissionID, rubricID uint) (*models.RubricScore, error) {
	db := database.DB
	var score models.RubricScore
	if err := db.Where("submission_id = ? AND rubric_id = ?", submissionID, rubricID).First(&score).Error; err != nil {
		return nil, err
	}
	return &score, nil
}

// Update updates a rubric score
func (r *RubricScoreRepository) Update(score *models.RubricScore) error {
	db := database.DB
	return db.Save(score).Error
}

// GetBySubmissionID retrieves all scores for a submission
func (r *RubricScoreRepository) GetBySubmissionID(submissionID uint) ([]models.RubricScore, error) {
	db := database.DB
	var scores []models.RubricScore
	if err := db.Where("submission_id = ?", submissionID).Find(&scores).Error; err != nil {
		return nil, err
	}
	return scores, nil
}
