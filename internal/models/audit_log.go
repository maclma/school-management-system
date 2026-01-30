package models

type AuditLog struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id"`
	Action    string `json:"action"`
	Entity    string `json:"entity"`
	EntityID  uint   `json:"entity_id"`
	OldValue  string `json:"old_value"`
	NewValue  string `json:"new_value"`
	IPAddress string `json:"ip_address"`
	Status    string `json:"status"`

	CreatedAt int64 `json:"created_at"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
