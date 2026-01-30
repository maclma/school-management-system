package models

type Notification struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Type      string `json:"type"` // email, sms, in-app
	Subject   string `json:"subject"`
	IsRead    bool   `json:"is_read"`
	SentAt    int64  `json:"sent_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}
