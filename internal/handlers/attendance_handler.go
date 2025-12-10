package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService service.AttendanceService
}

func NewAttendanceHandler(attendanceService service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{attendanceService: attendanceService}
}

type RecordAttendanceRequest struct {
	StudentID uint   `json:"student_id" binding:"required"`
	CourseID  uint   `json:"course_id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=present absent late excused"`
	Remarks   string `json:"remarks"`
	Date      string `json:"date"`
}

func (h *AttendanceHandler) RecordAttendance(c *gin.Context) {
	var req RecordAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance := &models.Attendance{
		StudentID: req.StudentID,
		CourseID:  req.CourseID,
		Status:    req.Status,
		Remarks:   req.Remarks,
		Date:      time.Now(),
	}

	if req.Date != "" {
		if parsedDate, err := time.Parse("2006-01-02", req.Date); err == nil {
			attendance.Date = parsedDate
		}
	}

	err := h.attendanceService.RecordAttendance(attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Attendance recorded successfully",
		"attendance": attendance,
	})
}

func (h *AttendanceHandler) GetAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	attendance, err := h.attendanceService.GetAttendanceByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func (h *AttendanceHandler) GetStudentAttendance(c *gin.Context) {
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

	attendances, total, err := h.attendanceService.GetStudentAttendance(uint(studentID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": total,
	})
}

func (h *AttendanceHandler) GetCourseAttendance(c *gin.Context) {
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

	attendances, total, err := h.attendanceService.GetCourseAttendance(uint(courseID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": total,
	})
}

func (h *AttendanceHandler) GetStudentCourseAttendance(c *gin.Context) {
	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

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

	attendances, total, err := h.attendanceService.GetStudentCourseAttendance(uint(studentID), uint(courseID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": total,
	})
}

func (h *AttendanceHandler) UpdateAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	var req RecordAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance := &models.Attendance{
		ID:      uint(id),
		Status:  req.Status,
		Remarks: req.Remarks,
	}

	err = h.attendanceService.UpdateAttendance(attendance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance updated successfully"})
}

func (h *AttendanceHandler) DeleteAttendance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	err = h.attendanceService.DeleteAttendance(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance deleted successfully"})
}

func (h *AttendanceHandler) GetAttendanceStats(c *gin.Context) {
	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	present, absent, late, err := h.attendanceService.GetStudentAttendanceStats(uint(studentID), uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stats"})
		return
	}

	percentage, _ := h.attendanceService.CalculateAttendancePercentage(uint(studentID), uint(courseID))

	c.JSON(http.StatusOK, gin.H{
		"student_id": studentID,
		"course_id":  courseID,
		"present":    present,
		"absent":     absent,
		"late":       late,
		"total":      present + absent + late,
		"percentage": percentage,
	})
}
