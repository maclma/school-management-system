package models

import (
	"time"
)

type Attendance struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `json:"student_id"`
	CourseID  uint      `json:"course_id"`
	Date      time.Time `json:"date"`
	Status    string    `gorm:"size:20;not null" json:"status"` // present, absent, late, excused
	Remarks   string    `gorm:"type:text" json:"remarks"`

	// Relations
	Student Student `gorm:"foreignKey:StudentID" json:"student"`
	Course  Course  `gorm:"foreignKey:CourseID" json:"course"`
}
