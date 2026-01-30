package models

type Announcement struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy uint   `json:"created_by"`
	Audience  string `json:"audience"` // all, students, teachers, specific_class
	Priority  string `json:"priority"` // low, normal, high
	IsActive  bool   `json:"is_active"`
	ExpiresAt int64  `json:"expires_at"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	CreatedByUser User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
}

func (Announcement) TableName() string {
	return "announcements"
}
