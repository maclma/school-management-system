package handlers

import (
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BackupHandler struct {
	service service.BackupService
}

func NewBackupHandler(svc service.BackupService) *BackupHandler {
	return &BackupHandler{service: svc}
}

func (h *BackupHandler) GetAll(c *gin.Context) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	backups, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Backups fetched", backups, page, limit, int64(total))
}

func (h *BackupHandler) GetLatest(c *gin.Context) {
	backup, err := h.service.GetLatestCompleted()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Latest backup fetched", backup)
}

func (h *BackupHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	backup, err := h.service.GetByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Backup fetched", backup)
}

func (h *BackupHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.NoContent(c)
}
