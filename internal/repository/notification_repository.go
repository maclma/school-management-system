package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *models.Notification) error
	FindByID(id uint) (*models.Notification, error)
	FindByUserID(userID uint, page, limit int) ([]models.Notification, int64, error)
	FindUnread(userID uint) ([]models.Notification, error)
	Update(notification *models.Notification) error
	Delete(id uint) error
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{db: database.DB}
}

func (r *notificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) FindByID(id uint) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.Preload("User").First(&notification, id).Error
	return &notification, err
}

func (r *notificationRepository) FindByUserID(userID uint, page, limit int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("user_id = ?", userID).Count(&total).
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error
	return notifications, total, err
}

func (r *notificationRepository) FindUnread(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ? AND is_read = ?", userID, false).
		Preload("User").
		Order("created_at DESC").
		Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) Update(notification *models.Notification) error {
	return r.db.Save(notification).Error
}

func (r *notificationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Notification{}, id).Error
}

func (r *notificationRepository) MarkAsRead(id uint) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

func (r *notificationRepository) MarkAllAsRead(userID uint) error {
	return r.db.Model(&models.Notification{}).Where("user_id = ?", userID).Update("is_read", true).Error
}
