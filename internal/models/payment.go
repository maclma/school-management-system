package models

type Payment struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	StudentID     uint    `json:"student_id"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
	Status        string  `json:"status"` // pending, paid, overdue, cancelled
	DueDate       int64   `json:"due_date"`
	PaidDate      int64   `json:"paid_date"`
	PaymentMethod string  `json:"payment_method"` // cash, check, online
	TransactionID string  `json:"transaction_id"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	Student Student `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}
