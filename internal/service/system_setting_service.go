package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type SystemSettingService interface {
	Create(setting *models.SystemSetting) error
	GetByKey(key string) (*models.SystemSetting, error)
	GetAll() ([]models.SystemSetting, error)
	Update(setting *models.SystemSetting) error
	Delete(id uint) error
	GetValue(key string, defaultValue string) string
}

type systemSettingService struct {
	repo repository.SystemSettingRepository
}

func NewSystemSettingService(repo repository.SystemSettingRepository) SystemSettingService {
	return &systemSettingService{repo: repo}
}

func (s *systemSettingService) Create(setting *models.SystemSetting) error {
	return s.repo.Create(setting)
}

func (s *systemSettingService) GetByKey(key string) (*models.SystemSetting, error) {
	return s.repo.FindByKey(key)
}

func (s *systemSettingService) GetAll() ([]models.SystemSetting, error) {
	return s.repo.FindAll()
}

func (s *systemSettingService) Update(setting *models.SystemSetting) error {
	return s.repo.Update(setting)
}

func (s *systemSettingService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *systemSettingService) GetValue(key string, defaultValue string) string {
	return s.repo.GetValue(key, defaultValue)
}
