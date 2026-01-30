package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(log *models.AuditLog) error
	FindByID(id uint) (*models.AuditLog, error)
	FindByUserID(userID uint, page, limit int) ([]models.AuditLog, int64, error)
	FindByEntity(entity string, entityID uint) ([]models.AuditLog, error)
	FindAll(page, limit int) ([]models.AuditLog, int64, error)
	DeleteOlderThan(days int) error
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository() AuditLogRepository {
	return &auditLogRepository{db: database.DB}
}

func (r *auditLogRepository) Create(log *models.AuditLog) error {
	return r.db.Create(log).Error
}

func (r *auditLogRepository) FindByID(id uint) (*models.AuditLog, error) {
	var log models.AuditLog
	err := r.db.First(&log, id).Error
	return &log, err
}

func (r *auditLogRepository) FindByUserID(userID uint, page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("user_id = ?", userID).Count(&total).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error
	return logs, total, err
}

func (r *auditLogRepository) FindByEntity(entity string, entityID uint) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Where("entity = ? AND entity_id = ?", entity, entityID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}

func (r *auditLogRepository) FindAll(page, limit int) ([]models.AuditLog, int64, error) {
	var logs []models.AuditLog
	var total int64
	offset := (page - 1) * limit
	err := r.db.Count(&total).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error
	return logs, total, err
}

func (r *auditLogRepository) DeleteOlderThan(days int) error {
	return r.db.Where("created_at < datetime('now', '-' || ? || ' days')", days).Delete(&models.AuditLog{}).Error
}
