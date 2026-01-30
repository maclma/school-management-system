package handlers

import (
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttendanceAutomationHandler struct {
	service *service.AttendanceAutomationService
}

func NewAttendanceAutomationHandler(svc *service.AttendanceAutomationService) *AttendanceAutomationHandler {
	return &AttendanceAutomationHandler{service: svc}
}

// GetAttendanceStats returns attendance statistics for a course
func (h *AttendanceAutomationHandler) GetAttendanceStats(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)

	stats, err := h.service.GetAttendanceStats(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Attendance statistics retrieved", stats)
}

// GetStudentAttendancePercentage returns attendance percentage for a student
func (h *AttendanceAutomationHandler) GetStudentAttendancePercentage(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)

	percentage, err := h.service.CalculateAttendancePercentage(uint(studentID), uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Attendance percentage retrieved", map[string]interface{}{
		"student_id":            studentID,
		"course_id":             courseID,
		"attendance_percentage": percentage,
	})
}

// CheckLowAttendance checks if student has low attendance
func (h *AttendanceAutomationHandler) CheckLowAttendance(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)
	threshold := 80.0
	if t := c.Query("threshold"); t != "" {
		if parsed, err := strconv.ParseFloat(t, 64); err == nil {
			threshold = parsed
		}
	}

	isLow, err := h.service.CheckLowAttendance(uint(studentID), uint(courseID), threshold)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Low attendance check completed", map[string]interface{}{
		"student_id":         studentID,
		"course_id":          courseID,
		"threshold":          threshold,
		"is_below_threshold": isLow,
	})
}

// GetStudentsWithLowAttendance returns all students below threshold
func (h *AttendanceAutomationHandler) GetStudentsWithLowAttendance(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)
	threshold := 80.0
	if t := c.Query("threshold"); t != "" {
		if parsed, err := strconv.ParseFloat(t, 64); err == nil {
			threshold = parsed
		}
	}

	students, err := h.service.GetStudentAttendanceStatusByThreshold(uint(courseID), threshold)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Students with low attendance retrieved", students)
}

// GetAttendanceReport generates a detailed attendance report
func (h *AttendanceAutomationHandler) GetAttendanceReport(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 32)

	report, err := h.service.GenerateAttendanceReport(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Attendance report generated", report)
}
