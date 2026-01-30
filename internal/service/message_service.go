package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type MessageService interface {
	Create(message *models.Message) error
	GetByID(id uint) (*models.Message, error)
	GetByReceiverID(receiverID uint, page, limit int) ([]models.Message, int64, error)
	GetConversation(userID1, userID2 uint, page, limit int) ([]models.Message, int64, error)
	Update(message *models.Message) error
	Delete(id uint) error
	MarkAsRead(id uint) error
	CountUnread(userID uint) (int64, error)
}

type messageService struct {
	repo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) Create(message *models.Message) error {
	return s.repo.Create(message)
}

func (s *messageService) GetByID(id uint) (*models.Message, error) {
	return s.repo.FindByID(id)
}

func (s *messageService) GetByReceiverID(receiverID uint, page, limit int) ([]models.Message, int64, error) {
	return s.repo.FindByReceiverID(receiverID, page, limit)
}

func (s *messageService) GetConversation(userID1, userID2 uint, page, limit int) ([]models.Message, int64, error) {
	return s.repo.FindConversation(userID1, userID2, page, limit)
}

func (s *messageService) Update(message *models.Message) error {
	return s.repo.Update(message)
}

func (s *messageService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *messageService) MarkAsRead(id uint) error {
	return s.repo.MarkAsRead(id)
}

func (s *messageService) CountUnread(userID uint) (int64, error) {
	return s.repo.CountUnread(userID)
}
