package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: svc}
}

func (h *PaymentHandler) Create(c *gin.Context) {
	var req struct {
		StudentID     uint    `json:"student_id" binding:"required"`
		Amount        float64 `json:"amount" binding:"required"`
		Description   string  `json:"description"`
		DueDate       int64   `json:"due_date"`
		PaymentMethod string  `json:"payment_method"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	payment := &models.Payment{
		StudentID:     req.StudentID,
		Amount:        req.Amount,
		Description:   req.Description,
		Status:        "pending",
		DueDate:       req.DueDate,
		PaymentMethod: req.PaymentMethod,
	}

	if err := h.service.Create(payment); err != nil {
		response.Error(c, errors.InternalError("Failed to create payment"))
		return
	}
	response.Created(c, "Payment created", payment)
}

func (h *PaymentHandler) GetByStudent(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	payments, total, err := h.service.GetByStudentID(uint(studentID), page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch payments"))
		return
	}

	response.Paginated(c, "Payments fetched", payments, page, limit, total)
}

func (h *PaymentHandler) GetAll(c *gin.Context) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	payments, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch payments"))
		return
	}

	response.Paginated(c, "Payments fetched", payments, page, limit, total)
}

func (h *PaymentHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var req struct {
		Status        string  `json:"status"`
		Amount        float64 `json:"amount"`
		PaymentMethod string  `json:"payment_method"`
		TransactionID string  `json:"transaction_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	payment, _ := h.service.GetByID(uint(id))
	if req.Status != "" {
		payment.Status = req.Status
	}
	if req.Amount > 0 {
		payment.Amount = req.Amount
	}
	if req.PaymentMethod != "" {
		payment.PaymentMethod = req.PaymentMethod
	}
	if req.TransactionID != "" {
		payment.TransactionID = req.TransactionID
	}

	if err := h.service.Update(payment); err != nil {
		response.Error(c, errors.InternalError("Failed to update payment"))
		return
	}
	response.Success(c, "Payment updated", payment)
}

func (h *PaymentHandler) GetStudentBalance(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)
	balance, err := h.service.GetStudentBalance(uint(studentID))
	if err != nil {
		response.Error(c, errors.InternalError("Failed to get balance"))
		return
	}
	response.Success(c, "Balance fetched", gin.H{"student_id": studentID, "balance": balance})
}
