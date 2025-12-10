package models

import (
	"time"
)

type Teacher struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"unique;not null" json:"user_id"`
	TeacherID     string    `gorm:"size:50;unique;not null" json:"teacher_id"`
	Department    string    `gorm:"size:100" json:"department"`
	Qualification string    `gorm:"type:text" json:"qualification"`
	HireDate      time.Time `json:"hire_date"`
	Salary        float64   `json:"salary"`

	// Relations
	User    User     `gorm:"foreignKey:UserID" json:"user"`
	Courses []Course `json:"courses,omitempty"`
}
