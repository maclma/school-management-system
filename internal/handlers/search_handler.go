package handlers

import (
	"school-management-system/internal/service"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchService *service.SearchService
}

func NewSearchHandler(svc *service.SearchService) *SearchHandler {
	return &SearchHandler{searchService: svc}
}

// SearchAnnouncements searches announcements with filters
func (h *SearchHandler) SearchAnnouncements(c *gin.Context) {
	query := c.Query("q")
	audience := c.Query("audience")
	priority := c.Query("priority")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	announcements, total, err := h.searchService.SearchAnnouncementsAdvanced(query, audience, priority, page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Search results", announcements, page, limit, total)
}

// SearchPayments searches payments with filters
func (h *SearchHandler) SearchPayments(c *gin.Context) {
	studentID := uint(0)
	if s := c.Query("student_id"); s != "" {
		if parsed, err := strconv.ParseUint(s, 10, 32); err == nil {
			studentID = uint(parsed)
		}
	}
	status := c.Query("status")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	payments, total, err := h.searchService.SearchPayments(studentID, status, page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Payment search results", payments, page, limit, total)
}

// SearchStudents searches for students
func (h *SearchHandler) SearchStudents(c *gin.Context) {
	query := c.Query("q")
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	students, total, err := h.searchService.SearchStudents(query, page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Student search results", students, page, limit, total)
}

// SearchGradesByRange searches grades in a point range
func (h *SearchHandler) SearchGradesByRange(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Query("course_id"), 10, 32)
	minScore, _ := strconv.ParseFloat(c.Query("min_score"), 64)
	maxScore, _ := strconv.ParseFloat(c.Query("max_score"), 64)

	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}
	limit := 20

	grades, total, err := h.searchService.SearchGradesByRange(uint(courseID), minScore, maxScore, page, limit)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, "Grade search results", grades, page, limit, total)
}

// SearchOverduePayments returns all overdue payments
func (h *SearchHandler) SearchOverduePayments(c *gin.Context) {
	payments, err := h.searchService.SearchOverduePayments()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Overdue payments retrieved", payments)
}
