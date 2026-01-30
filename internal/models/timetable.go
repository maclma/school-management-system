package models

type TimeTable struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	CourseID  uint   `json:"course_id"`
	TeacherID uint   `json:"teacher_id"`
	DayOfWeek string `json:"day_of_week"` // Monday, Tuesday, etc.
	StartTime string `json:"start_time"`  // HH:MM format
	EndTime   string `json:"end_time"`    // HH:MM format
	Classroom string `json:"classroom"`
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`

	Course  Course  `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Teacher Teacher `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
}

func (TimeTable) TableName() string {
	return "timetables"
}
