package models

import (
	"time"
)

type Grade struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `json:"student_id"`
	CourseID  uint      `json:"course_id"`
	Grade     string    `gorm:"size:5" json:"grade"`
	Score     float64   `json:"score"`
	MaxScore  float64   `gorm:"default:100" json:"max_score"`
	Remarks   string    `gorm:"type:text" json:"remarks"`
	GradedBy  uint      `json:"graded_by"`
	GradedAt  time.Time `json:"graded_at"`

	// Relations
	Student Student `gorm:"foreignKey:StudentID" json:"student"`
	Course  Course  `gorm:"foreignKey:CourseID" json:"course"`
	Teacher Teacher `gorm:"foreignKey:GradedBy" json:"teacher"`
}
