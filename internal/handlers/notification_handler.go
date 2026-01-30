package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service service.NotificationService
}

func NewNotificationHandler(svc service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: svc}
}

func (h *NotificationHandler) GetMyNotifications(c *gin.Context) {
	userID := c.GetUint("user_id")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil {
			limit = parsed
		}
	}

	notifs, total, err := h.service.GetByUserID(userID, page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch notifications"))
		return
	}

	response.Paginated(c, "Notifications fetched", notifs, page, limit, total)
}

func (h *NotificationHandler) GetUnread(c *gin.Context) {
	userID := c.GetUint("user_id")
	notifs, err := h.service.GetUnread(userID)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch unread notifications"))
		return
	}
	response.Success(c, "Unread notifications fetched", notifs)
}

func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.MarkAsRead(uint(id)); err != nil {
		response.Error(c, errors.InternalError("Failed to mark notification as read"))
		return
	}
	response.NoContent(c)
}

func (h *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	userID := c.GetUint("user_id")
	if err := h.service.MarkAllAsRead(userID); err != nil {
		response.Error(c, errors.InternalError("Failed to mark all as read"))
		return
	}
	response.NoContent(c)
}

func (h *NotificationHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, errors.InternalError("Failed to delete notification"))
		return
	}
	response.NoContent(c)
}

func (h *NotificationHandler) Create(c *gin.Context) {
	var req struct {
		UserID  uint   `json:"user_id" binding:"required"`
		Title   string `json:"title" binding:"required"`
		Message string `json:"message" binding:"required"`
		Type    string `json:"type" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	notif := &models.Notification{
		UserID:  req.UserID,
		Title:   req.Title,
		Message: req.Message,
		Type:    req.Type,
	}
	if err := h.service.Create(notif); err != nil {
		response.Error(c, errors.InternalError("Failed to create notification"))
		return
	}
	response.Created(c, "Notification created", notif)
}
