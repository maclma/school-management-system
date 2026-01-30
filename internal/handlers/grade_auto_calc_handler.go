package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GradeAutoCalcHandler struct {
	service *service.GradeAutoCalculationService
}

func NewGradeAutoCalcHandler(svc *service.GradeAutoCalculationService) *GradeAutoCalcHandler {
	return &GradeAutoCalcHandler{service: svc}
}

// RecordGradeWithAutoCalc records a grade and auto-calculates letter grade
func (h *GradeAutoCalcHandler) RecordGradeWithAutoCalc(c *gin.Context) {
	var req struct {
		StudentID uint    `json:"student_id" binding:"required"`
		CourseID  uint    `json:"course_id" binding:"required"`
		Score     float64 `json:"score" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	grade := &models.Grade{
		StudentID: req.StudentID,
		CourseID:  req.CourseID,
		Score:     req.Score,
	}

	if err := h.service.RecordGradeAndAutoCalculate(grade); err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Grade recorded and auto-calculated", grade)
}

// GetCourseAverage returns average grade for a course
func (h *GradeAutoCalcHandler) GetCourseAverage(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)

	avg, err := h.service.CalculateCourseAverage(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Course average retrieved", map[string]interface{}{
		"course_id": courseID,
		"average":   avg,
	})
}

// GetGradeDistribution returns grade distribution for a class
func (h *GradeAutoCalcHandler) GetGradeDistribution(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)

	distribution, err := h.service.CalculateClassGradeDistribution(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Grade distribution retrieved", distribution)
}

// GetStudentGradeStats returns comprehensive grade stats for a student
func (h *GradeAutoCalcHandler) GetStudentGradeStats(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)

	stats, err := h.service.GetStudentGradeStats(uint(studentID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Student grade statistics retrieved", stats)
}
