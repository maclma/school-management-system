package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnnouncementHandler struct {
	service service.AnnouncementService
}

func NewAnnouncementHandler(svc service.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{service: svc}
}

func (h *AnnouncementHandler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req struct {
		Title     string `json:"title" binding:"required"`
		Content   string `json:"content" binding:"required"`
		Audience  string `json:"audience" binding:"required"`
		Priority  string `json:"priority"`
		ExpiresAt int64  `json:"expires_at"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	announcement := &models.Announcement{
		Title:     req.Title,
		Content:   req.Content,
		CreatedBy: userID,
		Audience:  req.Audience,
		Priority:  req.Priority,
		IsActive:  true,
		ExpiresAt: req.ExpiresAt,
	}

	if err := h.service.Create(announcement); err != nil {
		response.Error(c, err)
		return
	}
	response.Created(c, "Announcement created", announcement)
}

func (h *AnnouncementHandler) GetAll(c *gin.Context) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	announcements, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Announcements fetched", announcements, page, limit, int64(total))
}

func (h *AnnouncementHandler) GetActive(c *gin.Context) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	announcements, total, err := h.service.GetActive(page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Active announcements fetched", announcements, page, limit, int64(total))
}

func (h *AnnouncementHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Audience  string `json:"audience"`
		Priority  string `json:"priority"`
		IsActive  bool   `json:"is_active"`
		ExpiresAt int64  `json:"expires_at"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	ann, _ := h.service.GetByID(uint(id))
	ann.Title = req.Title
	ann.Content = req.Content
	ann.Audience = req.Audience
	ann.Priority = req.Priority
	ann.IsActive = req.IsActive
	ann.ExpiresAt = req.ExpiresAt

	if err := h.service.Update(ann); err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, "Announcement updated", ann)
}

func (h *AnnouncementHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, err)
		return
	}
	response.NoContent(c)
}
