package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type AuditLogService interface {
	Log(log *models.AuditLog) error
	GetByID(id uint) (*models.AuditLog, error)
	GetByUserID(userID uint, page, limit int) ([]models.AuditLog, int64, error)
	GetByEntity(entity string, entityID uint) ([]models.AuditLog, error)
	GetAll(page, limit int) ([]models.AuditLog, int64, error)
	CleanupOldLogs(days int) error
}

type auditLogService struct {
	repo repository.AuditLogRepository
}

func NewAuditLogService(repo repository.AuditLogRepository) AuditLogService {
	return &auditLogService{repo: repo}
}

func (s *auditLogService) Log(log *models.AuditLog) error {
	return s.repo.Create(log)
}

func (s *auditLogService) GetByID(id uint) (*models.AuditLog, error) {
	return s.repo.FindByID(id)
}

func (s *auditLogService) GetByUserID(userID uint, page, limit int) ([]models.AuditLog, int64, error) {
	return s.repo.FindByUserID(userID, page, limit)
}

func (s *auditLogService) GetByEntity(entity string, entityID uint) ([]models.AuditLog, error) {
	return s.repo.FindByEntity(entity, entityID)
}

func (s *auditLogService) GetAll(page, limit int) ([]models.AuditLog, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *auditLogService) CleanupOldLogs(days int) error {
	return s.repo.DeleteOlderThan(days)
}
