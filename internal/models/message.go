package models

type Message struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	SenderID   uint   `json:"sender_id"`
	ReceiverID uint   `json:"receiver_id"`
	Content    string `json:"content"`
	IsRead     bool   `json:"is_read"`
	ReadAt     int64  `json:"read_at"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`

	Sender   User `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
}

func (Message) TableName() string {
	return "messages"
}
