package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type NotificationService interface {
	Create(notification *models.Notification) error
	GetByID(id uint) (*models.Notification, error)
	GetByUserID(userID uint, page, limit int) ([]models.Notification, int64, error)
	GetUnread(userID uint) ([]models.Notification, error)
	Update(notification *models.Notification) error
	Delete(id uint) error
	MarkAsRead(id uint) error
	MarkAllAsRead(userID uint) error
}

type notificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) Create(notification *models.Notification) error {
	return s.repo.Create(notification)
}

func (s *notificationService) GetByID(id uint) (*models.Notification, error) {
	return s.repo.FindByID(id)
}

func (s *notificationService) GetByUserID(userID uint, page, limit int) ([]models.Notification, int64, error) {
	return s.repo.FindByUserID(userID, page, limit)
}

func (s *notificationService) GetUnread(userID uint) ([]models.Notification, error) {
	return s.repo.FindUnread(userID)
}

func (s *notificationService) Update(notification *models.Notification) error {
	return s.repo.Update(notification)
}

func (s *notificationService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *notificationService) MarkAsRead(id uint) error {
	return s.repo.MarkAsRead(id)
}

func (s *notificationService) MarkAllAsRead(userID uint) error {
	return s.repo.MarkAllAsRead(userID)
}
