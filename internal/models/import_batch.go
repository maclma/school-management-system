package models

type ImportBatch struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	EntityType  string `json:"entity_type"` // student, teacher, course, grades
	FileName    string `json:"file_name"`
	TotalRows   int    `json:"total_rows"`
	SuccessRows int    `json:"success_rows"`
	FailedRows  int    `json:"failed_rows"`
	Status      string `json:"status"` // pending, processing, completed, failed
	Errors      string `json:"errors"`
	CreatedBy   uint   `json:"created_by"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	CreatedByUser User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
}

func (ImportBatch) TableName() string {
	return "import_batches"
}
