package models

import (
	"time"
)

type Assignment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CourseID    uint      `json:"course_id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	DueDate     time.Time `json:"due_date"`
	MaxScore    float64   `gorm:"default:100" json:"max_score"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`

	// Relations
	Course      Course                 `gorm:"foreignKey:CourseID" json:"course"`
	Teacher     Teacher                `gorm:"foreignKey:CreatedBy" json:"teacher"`
	Submissions []AssignmentSubmission `json:"submissions,omitempty"`
}

type AssignmentSubmission struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	AssignmentID uint       `json:"assignment_id"`
	StudentID    uint       `json:"student_id"`
	SubmittedAt  *time.Time `json:"submitted_at"`
	Score        *float64   `json:"score"`
	Feedback     string     `gorm:"type:text" json:"feedback"`
	FileURL      string     `gorm:"size:500" json:"file_url"`
	Status       string     `gorm:"size:20;default:'pending'" json:"status"` // pending, submitted, graded

	// Relations
	Assignment Assignment `gorm:"foreignKey:AssignmentID" json:"assignment"`
	Student    Student    `gorm:"foreignKey:StudentID" json:"student"`
}
