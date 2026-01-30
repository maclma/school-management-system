package handlers

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RubricHandler struct {
	rubricRepo *repository.AssignmentRubricRepository
	scoreRepo  *repository.RubricScoreRepository
}

func NewRubricHandler(rubricRepo *repository.AssignmentRubricRepository, scoreRepo *repository.RubricScoreRepository) *RubricHandler {
	return &RubricHandler{
		rubricRepo: rubricRepo,
		scoreRepo:  scoreRepo,
	}
}

// CreateRubric creates a new assignment rubric
func (h *RubricHandler) CreateRubric(c *gin.Context) {
	var req struct {
		AssignmentID uint   `json:"assignment_id" binding:"required"`
		Name         string `json:"name" binding:"required"`
		Description  string `json:"description"`
		TotalPoints  float64 `json:"total_points" binding:"required"`
		Criteria     []models.RubricCriterion `json:"criteria" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	rubric := &models.AssignmentRubric{
		AssignmentID: req.AssignmentID,
		Name:         req.Name,
		Description:  req.Description,
		TotalPoints:  req.TotalPoints,
		IsActive:     true,
	}

	// Set criteria
	if err := rubric.SetCriteria(req.Criteria); err != nil {
		response.Error(c, err)
		return
	}

	if err := h.rubricRepo.Create(rubric); err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Rubric created", rubric)
}

// GetRubric retrieves a rubric by ID
func (h *RubricHandler) GetRubric(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	rubric, err := h.rubricRepo.GetByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Rubric retrieved", rubric)
}

// GetRubricsByAssignment retrieves all rubrics for an assignment
func (h *RubricHandler) GetRubricsByAssignment(c *gin.Context) {
	assignmentID, _ := strconv.ParseUint(c.Param("assignment_id"), 10, 32)

	rubrics, err := h.rubricRepo.GetByAssignmentID(uint(assignmentID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Assignment rubrics retrieved", rubrics)
}

// UpdateRubric updates a rubric
func (h *RubricHandler) UpdateRubric(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	rubric, err := h.rubricRepo.GetByID(uint(id))
	if err != nil {
		response.Error(c, err)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		TotalPoints float64 `json:"total_points"`
		Criteria    []models.RubricCriterion `json:"criteria"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	if req.Name != "" {
		rubric.Name = req.Name
	}
	if req.Description != "" {
		rubric.Description = req.Description
	}
	if req.TotalPoints > 0 {
		rubric.TotalPoints = req.TotalPoints
	}
	if len(req.Criteria) > 0 {
		if err := rubric.SetCriteria(req.Criteria); err != nil {
			response.Error(c, err)
			return
		}
	}

	if err := h.rubricRepo.Update(rubric); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Rubric updated", rubric)
}

// DeleteRubric deletes a rubric
func (h *RubricHandler) DeleteRubric(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := h.rubricRepo.Delete(uint(id)); err != nil {
		response.Error(c, err)
		return
	}

	response.NoContent(c)
}

// ScoreSubmission scores a submission using a rubric
func (h *RubricHandler) ScoreSubmission(c *gin.Context) {
	submissionID, _ := strconv.ParseUint(c.Param("submission_id"), 10, 32)

	var req struct {
		RubricID      uint    `json:"rubric_id" binding:"required"`
		CriterionScores map[uint]float64 `json:"criterion_scores" binding:"required"`
		TotalScore    float64 `json:"total_score" binding:"required"`
		Comments      string  `json:"comments"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	score := &models.RubricScore{
		SubmissionID:      uint(submissionID),
		RubricID:          req.RubricID,
		TotalScore:        req.TotalScore,
		FeedbackComments:  req.Comments,
		ScoredByTeacherID: c.GetUint("user_id"),
	}

	// Set criterion scores
	// (In a real implementation, marshal the map to JSON)

	if err := h.scoreRepo.Create(score); err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Submission scored", score)
}

// GetSubmissionScore retrieves the rubric score for a submission
func (h *RubricHandler) GetSubmissionScore(c *gin.Context) {
	submissionID, _ := strconv.ParseUint(c.Param("submission_id"), 10, 32)
	rubricID, _ := strconv.ParseUint(c.Param("rubric_id"), 10, 32)

	score, err := h.scoreRepo.GetBySubmissionAndRubric(uint(submissionID), uint(rubricID))
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, "Submission score retrieved", score)
}
