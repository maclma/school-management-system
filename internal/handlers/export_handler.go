package handlers

import (
	"fmt"
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ExportHandler struct {
	exportService *service.ExportService
}

func NewExportHandler(svc *service.ExportService) *ExportHandler {
	return &ExportHandler{exportService: svc}
}

// ExportPaymentsCSV exports payments as CSV
func (h *ExportHandler) ExportPaymentsCSV(c *gin.Context) {
	studentID := uint(0)
	if s := c.Query("student_id"); s != "" {
		if parsed, err := strconv.ParseUint(s, 10, 32); err == nil {
			studentID = uint(parsed)
		}
	}
	status := c.Query("status")

	data, err := h.exportService.ExportPaymentsCSV(studentID, status)
	if err != nil {
		response.Error(c, err)
		return
	}

	filename := fmt.Sprintf("payments_%s.csv", time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}

// ExportGradesCSV exports grades as CSV
func (h *ExportHandler) ExportGradesCSV(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Query("course_id"), 10, 32)

	data, err := h.exportService.ExportGradesCSV(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	filename := fmt.Sprintf("grades_course_%d_%s.csv", courseID, time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}

// ExportAttendanceCSV exports attendance records as CSV
func (h *ExportHandler) ExportAttendanceCSV(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Query("course_id"), 10, 32)

	data, err := h.exportService.ExportAttendanceCSV(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	filename := fmt.Sprintf("attendance_course_%d_%s.csv", courseID, time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}

// ExportStudentTranscript exports student transcript as CSV
func (h *ExportHandler) ExportStudentTranscript(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Query("student_id"), 10, 32)

	data, err := h.exportService.ExportStudentTranscriptCSV(uint(studentID))
	if err != nil {
		response.Error(c, err)
		return
	}

	filename := fmt.Sprintf("transcript_student_%d_%s.csv", studentID, time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}

// ExportEnrollments exports course enrollments as CSV
func (h *ExportHandler) ExportEnrollments(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Query("course_id"), 10, 32)

	data, err := h.exportService.ExportEnrollmentsCSV(uint(courseID))
	if err != nil {
		response.Error(c, err)
		return
	}

	filename := fmt.Sprintf("enrollments_course_%d_%s.csv", courseID, time.Now().Format("2006-01-02"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "text/csv")
	c.Data(200, "text/csv", data)
}
