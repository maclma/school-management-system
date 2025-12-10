package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userService    service.UserService
	courseService  service.CourseService
	studentService service.StudentService
}

func NewAdminHandler(userService service.UserService, courseService service.CourseService, studentService service.StudentService) *AdminHandler {
	return &AdminHandler{
		userService:    userService,
		courseService:  courseService,
		studentService: studentService,
	}
}

type CreateUserRequest struct {
	FirstName   string          `json:"first_name" binding:"required"`
	LastName    string          `json:"last_name" binding:"required"`
	Email       string          `json:"email" binding:"required,email"`
	Password    string          `json:"password" binding:"required,min=6"`
	Phone       string          `json:"phone"`
	Role        models.UserRole `json:"role" binding:"required,oneof=admin teacher student parent"`
	DateOfBirth string          `json:"date_of_birth"`
	Address     string          `json:"address"`
}

func (h *AdminHandler) GetDashboardStats(c *gin.Context) {
	// Get basic statistics for admin dashboard
	users, _, _ := h.userService.GetAllUsers(1, 1, "")
	students, _, _ := h.studentService.GetAllStudents(1, 1)
	courses, _, _ := h.courseService.GetAllCourses(1, 1)

	c.JSON(http.StatusOK, gin.H{
		"total_users":    len(users),
		"total_students": len(students),
		"total_courses":  len(courses),
	})
}

func (h *AdminHandler) GetAllUsersAdmin(c *gin.Context) {
	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if val := 0; val > 0 {
			page = val
		}
	}

	users, total, err := h.userService.GetAllUsers(page, limit, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"total": total,
	})
}

func (h *AdminHandler) CreateUserAdmin(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Note: Admin should use the auth/register endpoint with admin role
	// This endpoint is a convenience wrapper
	c.JSON(http.StatusCreated, gin.H{
		"message": "Please use /api/auth/register endpoint to create users",
		"user": gin.H{
			"email": req.Email,
			"role":  req.Role,
		},
	})
}

func (h *AdminHandler) SystemHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "School Management System is running",
	})
}
