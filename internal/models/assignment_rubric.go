package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// RubricCriterion represents a grading criterion in a rubric
type RubricCriterion struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`       // e.g., "Clarity", "Organization"
	Weight      float64 `json:"weight"`     // Percentage weight (0-100)
	MaxPoints   float64 `json:"max_points"` // Maximum points for this criterion
	Description string  `json:"description"`
}

// AssignmentRubric defines grading criteria for assignments
type AssignmentRubric struct {
	ID           uint            `gorm:"primaryKey" json:"id"`
	AssignmentID uint            `json:"assignment_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	TotalPoints  float64         `json:"total_points"`
	Criteria     json.RawMessage `gorm:"type:json" json:"criteria"` // Array of RubricCriterion
	IsActive     bool            `json:"is_active" gorm:"default:true"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

// TableName specifies the table name for this model
func (AssignmentRubric) TableName() string {
	return "assignment_rubrics"
}

// GetCriteria returns parsed criteria
func (ar *AssignmentRubric) GetCriteria() ([]RubricCriterion, error) {
	var criteria []RubricCriterion
	if len(ar.Criteria) > 0 {
		err := json.Unmarshal(ar.Criteria, &criteria)
		if err != nil {
			return nil, err
		}
	}
	return criteria, nil
}

// SetCriteria sets criteria from a slice
func (ar *AssignmentRubric) SetCriteria(criteria []RubricCriterion) error {
	data, err := json.Marshal(criteria)
	if err != nil {
		return err
	}
	ar.Criteria = data
	return nil
}

// Value implements the driver.Valuer interface
func (ar AssignmentRubric) Value() (driver.Value, error) {
	return json.Marshal(ar)
}

// Scan implements the sql.Scanner interface
func (ar *AssignmentRubric) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion failed")
	}
	return json.Unmarshal(bytes, &ar)
}

// RubricScore represents a student's score on a rubric
type RubricScore struct {
	ID                uint            `gorm:"primaryKey" json:"id"`
	SubmissionID      uint            `json:"submission_id"`
	RubricID          uint            `json:"rubric_id"`
	CriterionScores   json.RawMessage `gorm:"type:json" json:"criterion_scores"` // Map of criterion -> score
	TotalScore        float64         `json:"total_score"`
	FeedbackComments  string          `json:"feedback_comments"`
	ScoredByTeacherID uint            `json:"scored_by_teacher_id"`
	ScoredAt          string          `json:"scored_at"`
	CreatedAt         string          `json:"created_at"`
	UpdatedAt         string          `json:"updated_at"`
}

// TableName specifies the table name
func (RubricScore) TableName() string {
	return "rubric_scores"
}
