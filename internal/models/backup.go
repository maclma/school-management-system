package models

type Backup struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	BackupName  string `json:"backup_name"`
	Description string `json:"description"`
	Size        int64  `json:"size"` // in bytes
	Location    string `json:"location"`
	Status      string `json:"status"` // pending, completed, failed
	CreatedBy   uint   `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`

	CreatedByUser User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
}

func (Backup) TableName() string {
	return "backups"
}
