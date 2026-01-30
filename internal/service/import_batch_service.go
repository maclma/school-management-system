package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type ImportBatchService interface {
	Create(batch *models.ImportBatch) error
	GetByID(id uint) (*models.ImportBatch, error)
	GetAll(page, limit int) ([]models.ImportBatch, int64, error)
	GetByStatus(status string, page, limit int) ([]models.ImportBatch, int64, error)
	Update(batch *models.ImportBatch) error
	Delete(id uint) error
	GetRecentByEntityType(entityType string, limit int) ([]models.ImportBatch, error)
}

type importBatchService struct {
	repo repository.ImportBatchRepository
}

func NewImportBatchService(repo repository.ImportBatchRepository) ImportBatchService {
	return &importBatchService{repo: repo}
}

func (s *importBatchService) Create(batch *models.ImportBatch) error {
	return s.repo.Create(batch)
}

func (s *importBatchService) GetByID(id uint) (*models.ImportBatch, error) {
	return s.repo.FindByID(id)
}

func (s *importBatchService) GetAll(page, limit int) ([]models.ImportBatch, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *importBatchService) GetByStatus(status string, page, limit int) ([]models.ImportBatch, int64, error) {
	return s.repo.FindByStatus(status, page, limit)
}

func (s *importBatchService) Update(batch *models.ImportBatch) error {
	return s.repo.Update(batch)
}

func (s *importBatchService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *importBatchService) GetRecentByEntityType(entityType string, limit int) ([]models.ImportBatch, error) {
	return s.repo.FindRecentByEntityType(entityType, limit)
}
