package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type SystemSettingRepository interface {
	Create(setting *models.SystemSetting) error
	FindByKey(key string) (*models.SystemSetting, error)
	FindAll() ([]models.SystemSetting, error)
	Update(setting *models.SystemSetting) error
	Delete(id uint) error
	GetValue(key string, defaultValue string) string
}

type systemSettingRepository struct {
	db *gorm.DB
}

func NewSystemSettingRepository() SystemSettingRepository {
	return &systemSettingRepository{db: database.DB}
}

func (r *systemSettingRepository) Create(setting *models.SystemSetting) error {
	return r.db.Create(setting).Error
}

func (r *systemSettingRepository) FindByKey(key string) (*models.SystemSetting, error) {
	var setting models.SystemSetting
	err := r.db.Where("key = ?", key).First(&setting).Error
	return &setting, err
}

func (r *systemSettingRepository) FindAll() ([]models.SystemSetting, error) {
	var settings []models.SystemSetting
	err := r.db.Find(&settings).Error
	return settings, err
}

func (r *systemSettingRepository) Update(setting *models.SystemSetting) error {
	return r.db.Save(setting).Error
}

func (r *systemSettingRepository) Delete(id uint) error {
	return r.db.Delete(&models.SystemSetting{}, id).Error
}

func (r *systemSettingRepository) GetValue(key string, defaultValue string) string {
	setting, err := r.FindByKey(key)
	if err != nil || setting == nil {
		return defaultValue
	}
	return setting.Value
}
