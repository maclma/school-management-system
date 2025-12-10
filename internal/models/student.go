package models

import (
	"time"
)

type Student struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"unique;not null" json:"user_id"`
	StudentID      string     `gorm:"size:50;unique;not null" json:"student_id"`
	GradeLevel     string     `gorm:"size:10" json:"grade_level"`
	EnrollmentDate time.Time  `json:"enrollment_date"`
	GraduationDate *time.Time `json:"graduation_date,omitempty"`
	ParentName     string     `gorm:"size:200" json:"parent_name"`
	ParentPhone    string     `gorm:"size:20" json:"parent_phone"`
	ParentEmail    string     `gorm:"size:100" json:"parent_email"`

	// Relations
	User        User         `gorm:"foreignKey:UserID" json:"user"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
	Attendances []Attendance `json:"attendances,omitempty"`
	Grades      []Grade      `json:"grades,omitempty"`
}
