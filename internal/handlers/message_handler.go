package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service service.MessageService
}

func NewMessageHandler(svc service.MessageService) *MessageHandler {
	return &MessageHandler{service: svc}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	senderID := c.GetUint("user_id")
	var req struct {
		ReceiverID uint   `json:"receiver_id" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	msg := &models.Message{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
		IsRead:     false,
	}
	if err := h.service.Create(msg); err != nil {
		response.Error(c, errors.InternalError("Failed to send message"))
		return
	}
	response.Created(c, "Message sent", msg)
}

func (h *MessageHandler) GetInbox(c *gin.Context) {
	userID := c.GetUint("user_id")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	msgs, total, err := h.service.GetByReceiverID(userID, page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch messages"))
		return
	}

	response.Paginated(c, "Messages fetched", msgs, page, limit, total)
}

func (h *MessageHandler) GetConversation(c *gin.Context) {
	userID := c.GetUint("user_id")
	otherID, _ := strconv.ParseUint(c.Param("user_id"), 10, 32)
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	msgs, total, err := h.service.GetConversation(userID, uint(otherID), page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch conversation"))
		return
	}

	response.Paginated(c, "Conversation fetched", msgs, page, limit, total)
}

func (h *MessageHandler) CountUnread(c *gin.Context) {
	userID := c.GetUint("user_id")
	count, err := h.service.CountUnread(userID)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to count unread"))
		return
	}
	response.Success(c, "Unread count", gin.H{"count": count})
}

func (h *MessageHandler) MarkAsRead(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.MarkAsRead(uint(id)); err != nil {
		response.Error(c, errors.InternalError("Failed to mark as read"))
		return
	}
	response.NoContent(c)
}
