package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type ImportBatchRepository interface {
	Create(batch *models.ImportBatch) error
	FindByID(id uint) (*models.ImportBatch, error)
	FindAll(page, limit int) ([]models.ImportBatch, int64, error)
	FindByStatus(status string, page, limit int) ([]models.ImportBatch, int64, error)
	Update(batch *models.ImportBatch) error
	Delete(id uint) error
	FindRecentByEntityType(entityType string, limit int) ([]models.ImportBatch, error)
}

type importBatchRepository struct {
	db *gorm.DB
}

func NewImportBatchRepository() ImportBatchRepository {
	return &importBatchRepository{db: database.DB}
}

func (r *importBatchRepository) Create(batch *models.ImportBatch) error {
	return r.db.Create(batch).Error
}

func (r *importBatchRepository) FindByID(id uint) (*models.ImportBatch, error) {
	var batch models.ImportBatch
	err := r.db.Preload("CreatedByUser").First(&batch, id).Error
	return &batch, err
}

func (r *importBatchRepository) FindAll(page, limit int) ([]models.ImportBatch, int64, error) {
	var batches []models.ImportBatch
	var total int64
	offset := (page - 1) * limit
	err := r.db.Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&batches).Error
	return batches, total, err
}

func (r *importBatchRepository) FindByStatus(status string, page, limit int) ([]models.ImportBatch, int64, error) {
	var batches []models.ImportBatch
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("status = ?", status).Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&batches).Error
	return batches, total, err
}

func (r *importBatchRepository) Update(batch *models.ImportBatch) error {
	return r.db.Save(batch).Error
}

func (r *importBatchRepository) Delete(id uint) error {
	return r.db.Delete(&models.ImportBatch{}, id).Error
}

func (r *importBatchRepository) FindRecentByEntityType(entityType string, limit int) ([]models.ImportBatch, error) {
	var batches []models.ImportBatch
	err := r.db.Where("entity_type = ?", entityType).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Find(&batches).Error
	return batches, err
}
