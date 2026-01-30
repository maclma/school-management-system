package models

type GradeTranscript struct {
	ID                 uint    `gorm:"primaryKey" json:"id"`
	StudentID          uint    `json:"student_id"`
	GPA                float64 `json:"gpa"`
	TotalCredits       float64 `json:"total_credits"`
	EarnedCredits      float64 `json:"earned_credits"`
	GradePointsSum     float64 `json:"grade_points_sum"`
	TranscriptSemester string  `json:"transcript_semester"`
	Year               int     `json:"year"`
	IsOfficial         bool    `json:"is_official"`
	GeneratedAt        int64   `json:"generated_at"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	Student Student `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}

func (GradeTranscript) TableName() string {
	return "grade_transcripts"
}
