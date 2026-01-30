package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Create(announcement *models.Announcement) error
	FindByID(id uint) (*models.Announcement, error)
	FindAll(page, limit int) ([]models.Announcement, int64, error)
	FindActive(page, limit int) ([]models.Announcement, int64, error)
	FindByAudience(audience string, page, limit int) ([]models.Announcement, int64, error)
	Update(announcement *models.Announcement) error
	Delete(id uint) error
}

type announcementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository() AnnouncementRepository {
	return &announcementRepository{db: database.DB}
}

func (r *announcementRepository) Create(announcement *models.Announcement) error {
	return r.db.Create(announcement).Error
}

func (r *announcementRepository) FindByID(id uint) (*models.Announcement, error) {
	var announcement models.Announcement
	err := r.db.Preload("CreatedByUser").First(&announcement, id).Error
	return &announcement, err
}

func (r *announcementRepository) FindAll(page, limit int) ([]models.Announcement, int64, error) {
	var announcements []models.Announcement
	var total int64
	offset := (page - 1) * limit
	err := r.db.Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&announcements).Error
	return announcements, total, err
}

func (r *announcementRepository) FindActive(page, limit int) ([]models.Announcement, int64, error) {
	var announcements []models.Announcement
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("is_active = ? AND (expires_at IS NULL OR expires_at > ?)", true, 0).
		Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&announcements).Error
	return announcements, total, err
}

func (r *announcementRepository) FindByAudience(audience string, page, limit int) ([]models.Announcement, int64, error) {
	var announcements []models.Announcement
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("audience = ?", audience).Count(&total).
		Preload("CreatedByUser").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&announcements).Error
	return announcements, total, err
}

func (r *announcementRepository) Update(announcement *models.Announcement) error {
	return r.db.Save(announcement).Error
}

func (r *announcementRepository) Delete(id uint) error {
	return r.db.Delete(&models.Announcement{}, id).Error
}
