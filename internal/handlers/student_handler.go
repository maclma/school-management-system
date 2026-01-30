package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	appErrors "school-management-system/pkg/errors"
	"school-management-system/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

type CreateStudentRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	StudentID   string `json:"student_id" binding:"required,min=3"`
	GradeLevel  string `json:"grade_level" binding:"required"`
	ParentName  string `json:"parent_name"`
	ParentPhone string `json:"parent_phone" binding:"omitempty,min=10"`
	ParentEmail string `json:"parent_email" binding:"omitempty,email"`
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, appErrors.BadRequest("invalid request body: "+err.Error()))
		return
	}

	// Validate student ID is not empty
	if req.StudentID == "" {
		response.Error(c, appErrors.MissingRequiredField("student_id"))
		return
	}

	student := &models.Student{
		UserID:         req.UserID,
		StudentID:      req.StudentID,
		GradeLevel:     req.GradeLevel,
		ParentName:     req.ParentName,
		ParentPhone:    req.ParentPhone,
		ParentEmail:    req.ParentEmail,
		EnrollmentDate: time.Now(),
	}

	err := h.studentService.CreateStudent(student)
	if err != nil {
		response.Error(c, appErrors.ServiceError("StudentService", "CreateStudent", err))
		return
	}

	response.Created(c, "Student created successfully", student)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, appErrors.InvalidFieldFormat("id", "unsigned integer"))
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		response.Error(c, appErrors.NotFound("Student not found"))
		return
	}

	response.Success(c, "Student retrieved successfully", student)
}

func (h *StudentHandler) GetAllStudents(c *gin.Context) {
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

	students, total, err := h.studentService.GetAllStudents(page, limit)
	if err != nil {
		response.Error(c, appErrors.ServiceError("StudentService", "GetAllStudents", err))
		return
	}

	response.Paginated(c, "Students retrieved successfully", students, page, limit, total)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, appErrors.InvalidFieldFormat("id", "unsigned integer"))
		return
	}

	var req CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, appErrors.BadRequest("invalid request body: "+err.Error()))
		return
	}

	student := &models.Student{
		ID:          uint(id),
		UserID:      req.UserID,
		StudentID:   req.StudentID,
		GradeLevel:  req.GradeLevel,
		ParentName:  req.ParentName,
		ParentPhone: req.ParentPhone,
		ParentEmail: req.ParentEmail,
	}

	err = h.studentService.UpdateStudent(student)
	if err != nil {
		response.Error(c, appErrors.ServiceError("StudentService", "UpdateStudent", err))
		return
	}

	response.Success(c, "Student updated successfully", nil)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, appErrors.InvalidFieldFormat("id", "unsigned integer"))
		return
	}

	err = h.studentService.DeleteStudent(uint(id))
	if err != nil {
		response.Error(c, appErrors.ServiceError("StudentService", "DeleteStudent", err))
		return
	}

	response.NoContent(c)
}
