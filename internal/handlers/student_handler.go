package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

type CreateStudentRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	StudentID   string `json:"student_id" binding:"required"`
	GradeLevel  string `json:"grade_level"`
	ParentName  string `json:"parent_name"`
	ParentPhone string `json:"parent_phone"`
	ParentEmail string `json:"parent_email"`
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := &models.Student{
		UserID:         req.UserID,
		StudentID:      req.StudentID,
		GradeLevel:     req.GradeLevel,
		ParentName:     req.ParentName,
		ParentPhone:    req.ParentPhone,
		ParentEmail:    req.ParentEmail,
		EnrollmentDate: time.Now(),
	}

	err := h.studentService.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"student": student,
	})
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 && val <= 100 {
			limit = val
		}
	}

	students, total, err := h.studentService.GetAllStudents(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  students,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := &models.Student{
		ID:          uint(id),
		UserID:      req.UserID,
		StudentID:   req.StudentID,
		GradeLevel:  req.GradeLevel,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		ParentEmail: req.ParentEmail,
	}

	err = h.studentService.UpdateStudent(student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	err = h.studentService.DeleteStudent(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

// Student handler placeholder. Implement HTTP handlers for students here as needed.
