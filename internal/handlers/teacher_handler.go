package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TeacherHandler struct {
	teacherService service.TeacherService
	logger         *logrus.Logger
}

func NewTeacherHandler(teacherService service.TeacherService) *TeacherHandler {
	return &TeacherHandler{
		teacherService: teacherService,
		logger:         logger.GetLogger(),
	}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		h.logger.WithError(err).Error("Invalid teacher data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.teacherService.CreateTeacher(&teacher); err != nil {
		h.logger.WithError(err).Error("Failed to create teacher")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, teacher)
}

func (h *TeacherHandler) GetTeacher(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	teacher, err := h.teacherService.GetTeacherByID(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get teacher")
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	teachers, total, err := h.teacherService.GetAllTeachers(page, limit)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get teachers")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teachers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teachers": teachers,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		h.logger.WithError(err).Error("Invalid teacher data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	teacher.ID = uint(id)

	if err := h.teacherService.UpdateTeacher(&teacher); err != nil {
		h.logger.WithError(err).Error("Failed to update teacher")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	if err := h.teacherService.DeleteTeacher(uint(id)); err != nil {
		h.logger.WithError(err).Error("Failed to delete teacher")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}

func (h *TeacherHandler) GetTeacherCourses(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	courses, total, err := h.teacherService.GetTeacherCourses(uint(id), page, limit)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get teacher courses")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
		"total":   total,
		"page":    page,
		"limit":   limit,
	})
}

func (h *TeacherHandler) GetTeachersByDepartment(c *gin.Context) {
	department := c.Query("department")
	if department == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Department parameter is required"})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	teachers, total, err := h.teacherService.GetTeachersByDepartment(department, page, limit)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get teachers by department")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teachers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"teachers": teachers,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}
