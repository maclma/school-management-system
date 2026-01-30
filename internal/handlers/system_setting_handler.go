package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SystemSettingHandler struct {
	service service.SystemSettingService
}

func NewSystemSettingHandler(service service.SystemSettingService) *SystemSettingHandler {
	return &SystemSettingHandler{service: service}
}

func (h *SystemSettingHandler) GetAll(c *gin.Context) {
	settings, err := h.service.GetAll()
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch system settings"))
		return
	}
	response.Success(c, "System settings fetched", settings)
}

func (h *SystemSettingHandler) GetByKey(c *gin.Context) {
	key := c.Param("key")
	setting, err := h.service.GetByKey(key)
	if err != nil {
		response.Error(c, errors.NotFound("Setting not found"))
		return
	}
	response.Success(c, "Setting fetched", setting)
}

func (h *SystemSettingHandler) Create(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	setting := &models.SystemSetting{Key: req.Key, Value: req.Value}

	if err := h.service.Create(setting); err != nil {
		response.Error(c, errors.InternalError("Failed to create setting"))
		return
	}
	response.Created(c, "Setting created successfully", setting)
}

func (h *SystemSettingHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	setting, err := h.service.GetByKey(req.Key)
	if err != nil {
		response.Error(c, errors.NotFound("Setting not found"))
		return
	}

	setting.Value = req.Value
	if err := h.service.Update(setting); err != nil {
		response.Error(c, errors.InternalError("Failed to update setting"))
		return
	}
	_ = id
	response.Success(c, "Setting updated successfully", setting)
}

func (h *SystemSettingHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, errors.InternalError("Failed to delete setting"))
		return
	}
	response.NoContent(c)
}
