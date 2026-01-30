package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type BackupRepository interface {
	Create(backup *models.Backup) error
	FindByID(id uint) (*models.Backup, error)
	FindAll(page, limit int) ([]models.Backup, int64, error)
	FindByStatus(status string) ([]models.Backup, error)
	Update(backup *models.Backup) error
	Delete(id uint) error
	GetLatestCompleted() (*models.Backup, error)
}

type backupRepository struct {
	db *gorm.DB
}

func NewBackupRepository() BackupRepository {
	return &backupRepository{db: database.DB}
}

func (r *backupRepository) Create(backup *models.Backup) error {
	return r.db.Create(backup).Error
}

func (r *backupRepository) FindByID(id uint) (*models.Backup, error) {
	var backup models.Backup
	err := r.db.Preload("CreatedByUser").First(&backup, id).Error
	return &backup, err
}

func (r *backupRepository) FindAll(page, limit int) ([]models.Backup, int64, error) {
	var backups []models.Backup
	var total int64
	offset := (page - 1) * limit
	err := r.db.Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&backups).Error
	return backups, total, err
}

func (r *backupRepository) FindByStatus(status string) ([]models.Backup, error) {
	var backups []models.Backup
	err := r.db.Where("status = ?", status).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Find(&backups).Error
	return backups, err
}

func (r *backupRepository) Update(backup *models.Backup) error {
	return r.db.Save(backup).Error
}

func (r *backupRepository) Delete(id uint) error {
	return r.db.Delete(&models.Backup{}, id).Error
}

func (r *backupRepository) GetLatestCompleted() (*models.Backup, error) {
	var backup models.Backup
	err := r.db.Where("status = ?", "completed").
		Preload("CreatedByUser").
		Order("created_at DESC").
		First(&backup).Error
	return &backup, err
}
