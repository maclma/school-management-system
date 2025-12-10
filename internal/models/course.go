package models

type Course struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CourseCode  string `gorm:"size:20;unique;not null" json:"course_code"`
	Name        string `gorm:"size:200;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	CreditHours int    `json:"credit_hours"`
	Department  string `gorm:"size:100" json:"department"`
	TeacherID   uint   `json:"teacher_id"`
	Room        string `gorm:"size:50" json:"room"`
	Schedule    string `gorm:"size:100" json:"schedule"`
	MaxStudents int    `json:"max_students"`

	// Relations
	Teacher     Teacher      `gorm:"foreignKey:TeacherID" json:"teacher"`
	Enrollments []Enrollment `json:"enrollments,omitempty"`
	Grades      []Grade      `json:"grades,omitempty"`
}
