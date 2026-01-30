package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type BackupService interface {
	Create(backup *models.Backup) error
	GetByID(id uint) (*models.Backup, error)
	GetAll(page, limit int) ([]models.Backup, int64, error)
	GetByStatus(status string) ([]models.Backup, error)
	Update(backup *models.Backup) error
	Delete(id uint) error
	GetLatestCompleted() (*models.Backup, error)
}

type backupService struct {
	repo repository.BackupRepository
}

func NewBackupService(repo repository.BackupRepository) BackupService {
	return &backupService{repo: repo}
}

func (s *backupService) Create(backup *models.Backup) error {
	return s.repo.Create(backup)
}

func (s *backupService) GetByID(id uint) (*models.Backup, error) {
	return s.repo.FindByID(id)
}

func (s *backupService) GetAll(page, limit int) ([]models.Backup, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *backupService) GetByStatus(status string) ([]models.Backup, error) {
	return s.repo.FindByStatus(status)
}

func (s *backupService) Update(backup *models.Backup) error {
	return s.repo.Update(backup)
}

func (s *backupService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *backupService) GetLatestCompleted() (*models.Backup, error) {
	return s.repo.GetLatestCompleted()
}
