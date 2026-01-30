package models

type SystemSetting struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Key   string `gorm:"uniqueIndex;not null" json:"key"`
	Value string `json:"value"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (SystemSetting) TableName() string {
	return "system_settings"
}
