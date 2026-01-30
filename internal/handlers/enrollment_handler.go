package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	enrollmentService service.EnrollmentService
	studentService    service.StudentService
}

func NewEnrollmentHandler(enrollmentService service.EnrollmentService, studentService service.StudentService) *EnrollmentHandler {
	return &EnrollmentHandler{
		enrollmentService: enrollmentService,
		studentService:    studentService,
	}
}

type EnrollStudentRequest struct {
	StudentID uint   `json:"student_id" binding:"required"`
	CourseID  uint   `json:"course_id" binding:"required"`
	Status    string `json:"status"`
}

func (h *EnrollmentHandler) EnrollStudent(c *gin.Context) {
	var req EnrollStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment := &models.Enrollment{
		StudentID:  req.StudentID,
		CourseID:   req.CourseID,
		EnrolledAt: time.Now(),
		Status:     "active",
	}

	if req.Status != "" {
		enrollment.Status = req.Status
	}

	err := h.enrollmentService.EnrollStudent(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Student enrolled successfully",
		"enrollment": enrollment,
	})
}

func (h *EnrollmentHandler) GetEnrollment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	enrollment, err := h.enrollmentService.GetEnrollmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (h *EnrollmentHandler) GetStudentEnrollments(c *gin.Context) {
	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	enrollments, total, err := h.enrollmentService.GetStudentEnrollments(uint(studentID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  enrollments,
		"total": total,
	})
}

func (h *EnrollmentHandler) GetCourseEnrollments(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	enrollments, total, err := h.enrollmentService.GetCourseEnrollments(uint(courseID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  enrollments,
		"total": total,
	})
}

func (h *EnrollmentHandler) UpdateEnrollmentStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.enrollmentService.UpdateEnrollmentStatus(uint(id), req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment status updated"})
}

func (h *EnrollmentHandler) RemoveEnrollment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	err = h.enrollmentService.RemoveEnrollment(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment removed"})
}

func (h *EnrollmentHandler) GetAllEnrollments(c *gin.Context) {
	page := 1
	limit := 50
	status := c.Query("status")

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	enrollments, total, err := h.enrollmentService.GetAllEnrollments(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	// Filter by status if provided
	if status != "" {
		var filtered []models.Enrollment
		for _, e := range enrollments {
			if e.Status == status {
				filtered = append(filtered, e)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data":  filtered,
			"total": int64(len(filtered)),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  enrollments,
		"total": total,
	})
}

func (h *EnrollmentHandler) ApproveEnrollment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	err = h.enrollmentService.UpdateEnrollmentStatus(uint(id), "approved")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment approved"})
}

func (h *EnrollmentHandler) RejectEnrollment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	err = h.enrollmentService.UpdateEnrollmentStatus(uint(id), "rejected")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment rejected"})
}

func (h *EnrollmentHandler) GetMyEnrollments(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Get student ID from user ID
	student, err := h.studentService.GetStudentByUserID(userIDUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student record not found"})
		return
	}

	page := 1
	limit := 50

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	enrollments, total, err := h.enrollmentService.GetStudentEnrollments(student.ID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  enrollments,
		"total": total,
	})
}
