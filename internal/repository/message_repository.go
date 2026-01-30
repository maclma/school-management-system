package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(message *models.Message) error
	FindByID(id uint) (*models.Message, error)
	FindByReceiverID(receiverID uint, page, limit int) ([]models.Message, int64, error)
	FindConversation(userID1, userID2 uint, page, limit int) ([]models.Message, int64, error)
	Update(message *models.Message) error
	Delete(id uint) error
	MarkAsRead(id uint) error
	CountUnread(userID uint) (int64, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{db: database.DB}
}

func (r *messageRepository) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) FindByID(id uint) (*models.Message, error) {
	var message models.Message
	err := r.db.Preload("Sender").Preload("Receiver").First(&message, id).Error
	return &message, err
}

func (r *messageRepository) FindByReceiverID(receiverID uint, page, limit int) ([]models.Message, int64, error) {
	var messages []models.Message
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("receiver_id = ?", receiverID).Count(&total).
		Preload("Sender").
		Preload("Receiver").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, total, err
}

func (r *messageRepository) FindConversation(userID1, userID2 uint, page, limit int) ([]models.Message, int64, error) {
	var messages []models.Message
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", userID1, userID2, userID2, userID1).
		Count(&total).
		Preload("Sender").
		Preload("Receiver").
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, total, err
}

func (r *messageRepository) Update(message *models.Message) error {
	return r.db.Save(message).Error
}

func (r *messageRepository) Delete(id uint) error {
	return r.db.Delete(&models.Message{}, id).Error
}

func (r *messageRepository) MarkAsRead(id uint) error {
	return r.db.Model(&models.Message{}).Where("id = ?", id).Updates(map[string]interface{}{"is_read": true, "read_at": 0}).Error
}

func (r *messageRepository) CountUnread(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Message{}).Where("receiver_id = ? AND is_read = ?", userID, false).Count(&count).Error
	return count, err
}
