package handlers

import (
	"school-management-system/internal/service"
	"school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GradeTranscriptHandler struct {
	service service.GradeTranscriptService
}

func NewGradeTranscriptHandler(svc service.GradeTranscriptService) *GradeTranscriptHandler {
	return &GradeTranscriptHandler{service: svc}
}

func (h *GradeTranscriptHandler) GetByStudentID(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	transcripts, total, err := h.service.GetByStudentID(uint(studentID), page, limit)
	if err != nil {
		response.Error(c, errors.InternalError("Failed to fetch transcripts"))
		return
	}

	response.Paginated(c, "Transcripts fetched", transcripts, page, limit, total)
}

func (h *GradeTranscriptHandler) GetLatest(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)

	transcript, err := h.service.GetLatestByStudent(uint(studentID))
	if err != nil {
		response.Error(c, errors.NotFound("Transcript not found"))
		return
	}

	response.Success(c, "Latest transcript fetched", transcript)
}

func (h *GradeTranscriptHandler) GetGPA(c *gin.Context) {
	studentID, _ := strconv.ParseUint(c.Param("student_id"), 10, 32)

	gpa, err := h.service.CalculateGPA(uint(studentID))
	if err != nil {
		response.Error(c, errors.InternalError("Failed to calculate GPA"))
		return
	}

	response.Success(c, "GPA calculated", gin.H{"student_id": studentID, "gpa": gpa})
}
