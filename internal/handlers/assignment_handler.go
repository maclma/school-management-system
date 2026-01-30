package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	assignmentService service.AssignmentService
	submissionService service.AssignmentSubmissionService
	studentService    service.StudentService
}

func NewAssignmentHandler(assignmentService service.AssignmentService, submissionService service.AssignmentSubmissionService, studentService service.StudentService) *AssignmentHandler {
	return &AssignmentHandler{
		assignmentService: assignmentService,
		submissionService: submissionService,
		studentService:    studentService,
	}
}

type CreateAssignmentRequest struct {
	CourseID    uint      `json:"course_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	MaxScore    float64   `json:"max_score"`
}

type SubmitAssignmentRequest struct {
	AssignmentID uint   `json:"assignment_id" binding:"required"`
	FileURL      string `json:"file_url"`
}

type GradeSubmissionRequest struct {
	Score    float64 `json:"score" binding:"required"`
	Feedback string  `json:"feedback"`
}

// Assignment CRUD operations
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var req CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get teacher ID from context (set by auth middleware)
	teacherID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	assignment := &models.Assignment{
		CourseID:    req.CourseID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		MaxScore:    req.MaxScore,
		CreatedBy:   teacherID.(uint),
	}

	err := h.assignmentService.CreateAssignment(assignment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Assignment created successfully",
		"assignment": assignment,
	})
}

func (h *AssignmentHandler) GetAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	assignment, err := h.assignmentService.GetAssignmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func (h *AssignmentHandler) GetAssignmentsByCourse(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	assignments, err := h.assignmentService.GetAssignmentsByCourse(uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"assignments": assignments,
		"count":       len(assignments),
	})
}

func (h *AssignmentHandler) GetAssignmentsByTeacher(c *gin.Context) {
	teacherID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	assignments, err := h.assignmentService.GetAssignmentsByTeacher(teacherID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"assignments": assignments,
		"count":       len(assignments),
	})
}

func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	var req CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment, err := h.assignmentService.GetAssignmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Check if user is the creator
	userID, exists := c.Get("user_id")
	if !exists || assignment.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own assignments"})
		return
	}

	assignment.CourseID = req.CourseID
	assignment.Title = req.Title
	assignment.Description = req.Description
	assignment.DueDate = req.DueDate
	assignment.MaxScore = req.MaxScore

	err = h.assignmentService.UpdateAssignment(assignment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Assignment updated successfully",
		"assignment": assignment,
	})
}

func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	assignment, err := h.assignmentService.GetAssignmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Check if user is the creator
	userID, exists := c.Get("user_id")
	if !exists || assignment.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own assignments"})
		return
	}

	err = h.assignmentService.DeleteAssignment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}

// Assignment Submission operations
func (h *AssignmentHandler) SubmitAssignment(c *gin.Context) {
	var req SubmitAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	submission := &models.AssignmentSubmission{
		AssignmentID: req.AssignmentID,
		StudentID:    studentID.(uint),
		FileURL:      req.FileURL,
		Status:       "submitted",
	}

	err := h.submissionService.SubmitAssignment(submission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Assignment submitted successfully",
		"submission": submission,
	})
}

func (h *AssignmentHandler) GetSubmissionsByAssignment(c *gin.Context) {
	assignmentID, err := strconv.ParseUint(c.Param("assignment_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}

	// Check if user is the teacher of this assignment
	assignment, err := h.assignmentService.GetAssignmentByID(uint(assignmentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists || assignment.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only view submissions for your own assignments"})
		return
	}

	submissions, err := h.submissionService.GetSubmissionsByAssignment(uint(assignmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch submissions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"submissions": submissions,
		"count":       len(submissions),
	})
}

func (h *AssignmentHandler) GetSubmissionsByStudent(c *gin.Context) {
	studentID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	submissions, err := h.submissionService.GetSubmissionsByStudent(studentID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch submissions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"submissions": submissions,
		"count":       len(submissions),
	})
}

func (h *AssignmentHandler) GradeSubmission(c *gin.Context) {
	submissionID, err := strconv.ParseUint(c.Param("submission_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid submission ID"})
		return
	}

	var req GradeSubmissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user is the teacher of this assignment
	submission, err := h.submissionService.GetSubmissionByID(uint(submissionID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	assignment, err := h.assignmentService.GetAssignmentByID(submission.AssignmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists || assignment.CreatedBy != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only grade submissions for your own assignments"})
		return
	}

	err = h.submissionService.GradeSubmission(uint(submissionID), req.Score, req.Feedback)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission graded successfully"})
}
