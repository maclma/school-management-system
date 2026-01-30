package handlers

import (
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImportBatchHandler struct {
	service service.ImportBatchService
}

func NewImportBatchHandler(svc service.ImportBatchService) *ImportBatchHandler {
	return &ImportBatchHandler{service: svc}
}

func (h *ImportBatchHandler) GetAll(c *gin.Context) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	batches, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Batches fetched", batches, page, limit, int64(total))
}

func (h *ImportBatchHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	batch, err := h.service.GetByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Batch fetched", batch)
}

func (h *ImportBatchHandler) GetByStatus(c *gin.Context) {
	status := c.Param("status")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	batches, total, err := h.service.GetByStatus(status, page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Batches fetched", batches, page, limit, int64(total))
}

func (h *ImportBatchHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.NoContent(c)
}
