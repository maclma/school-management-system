package models

import (
	"time"
)

type Enrollment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	StudentID  uint      `json:"student_id"`
	CourseID   uint      `json:"course_id"`
	EnrolledAt time.Time `json:"enrolled_at"`
	Status     string    `gorm:"size:20;default:'active'" json:"status"`

	// Relations
	Student Student `gorm:"foreignKey:StudentID" json:"student"`
	Course  Course  `gorm:"foreignKey:CourseID" json:"course"`
}
