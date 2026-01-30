package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GradeHandler struct {
	gradeService   service.GradeService
	studentService service.StudentService
}

func NewGradeHandler(gradeService service.GradeService, studentService service.StudentService) *GradeHandler {
	return &GradeHandler{
		gradeService:   gradeService,
		studentService: studentService,
	}
}

type RecordGradeRequest struct {
	StudentID uint    `json:"student_id" binding:"required"`
	CourseID  uint    `json:"course_id" binding:"required"`
	Grade     string  `json:"grade"`
	Score     float64 `json:"score" binding:"required"`
	MaxScore  float64 `json:"max_score"`
	Remarks   string  `json:"remarks"`
	GradedBy  uint    `json:"graded_by" binding:"required"`
}

func (h *GradeHandler) RecordGrade(c *gin.Context) {
	var req RecordGradeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade := &models.Grade{
		StudentID: req.StudentID,
		CourseID:  req.CourseID,
		Grade:     req.Grade,
		Score:     req.Score,
		MaxScore:  req.MaxScore,
		Remarks:   req.Remarks,
		GradedBy:  req.GradedBy,
		GradedAt:  time.Now(),
	}

	if grade.MaxScore == 0 {
		grade.MaxScore = 100
	}

	err := h.gradeService.RecordGrade(grade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Grade recorded successfully",
		"grade":   grade,
	})
}

func (h *GradeHandler) GetGrade(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	grade, err := h.gradeService.GetGradeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grade)
}

func (h *GradeHandler) GetStudentGrades(c *gin.Context) {
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

	grades, total, err := h.gradeService.GetStudentGrades(uint(studentID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  grades,
		"total": total,
	})
}

func (h *GradeHandler) GetCourseGrades(c *gin.Context) {
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

	grades, total, err := h.gradeService.GetCourseGrades(uint(courseID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  grades,
		"total": total,
	})
}

func (h *GradeHandler) UpdateGrade(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	var req RecordGradeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grade := &models.Grade{
		ID:       uint(id),
		Grade:    req.Grade,
		Score:    req.Score,
		MaxScore: req.MaxScore,
		Remarks:  req.Remarks,
	}

	err = h.gradeService.UpdateGrade(grade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade updated successfully"})
}

func (h *GradeHandler) DeleteGrade(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	err = h.gradeService.DeleteGrade(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted successfully"})
}

func (h *GradeHandler) GetAverageGrade(c *gin.Context) {
	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	average, err := h.gradeService.CalculateAverageGrade(uint(studentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate average"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student_id": studentID,
		"average":    average,
	})
}

func (h *GradeHandler) GetMyGrades(c *gin.Context) {
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

	grades, total, err := h.gradeService.GetStudentGrades(student.ID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch grades"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  grades,
		"total": total,
	})
}
