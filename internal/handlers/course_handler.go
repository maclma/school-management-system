package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{courseService: courseService}
}

type CreateCourseRequest struct {
	CourseCode  string `json:"course_code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CreditHours int    `json:"credit_hours"`
	Department  string `json:"department"`
	TeacherID   uint   `json:"teacher_id"`
	Room        string `json:"room"`
	Schedule    string `json:"schedule"`
	MaxStudents int    `json:"max_students"`
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := &models.Course{
		CourseCode:  req.CourseCode,
		Name:        req.Name,
		Description: req.Description,
		CreditHours: req.CreditHours,
		Department:  req.Department,
		TeacherID:   req.TeacherID,
		Room:        req.Room,
		Schedule:    req.Schedule,
		MaxStudents: req.MaxStudents,
	}

	err := h.courseService.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})
}

func (h *CourseHandler) GetCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.courseService.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
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

	courses, total, err := h.courseService.GetAllCourses(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  courses,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *CourseHandler) GetCoursesByDepartment(c *gin.Context) {
	department := c.Query("department")
	if department == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Department is required"})
		return
	}

	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	courses, total, err := h.courseService.GetCoursesByDepartment(department, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  courses,
		"total": total,
	})
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var req CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := &models.Course{
		ID:          uint(id),
		CourseCode:  req.CourseCode,
		Name:        req.Name,
		Description: req.Description,
		CreditHours: req.CreditHours,
		Department:  req.Department,
		TeacherID:   req.TeacherID,
		Room:        req.Room,
		Schedule:    req.Schedule,
		MaxStudents: req.MaxStudents,
	}

	err = h.courseService.UpdateCourse(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = h.courseService.DeleteCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
