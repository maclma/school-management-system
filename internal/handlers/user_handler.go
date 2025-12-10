package handlers

import (
	"net/http"
	"school-management-system/internal/models"
	"school-management-system/internal/service"
	"school-management-system/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.UserService
	logger      *logrus.Logger
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger.GetLogger(),
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := h.userService.GetUserByID(uint(userID.(float64)))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user profile")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hide sensitive data
	user.Password = ""

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove fields that shouldn't be updated via profile
	delete(updateData, "role")
	delete(updateData, "is_active")
	delete(updateData, "password")

	user, err := h.userService.UpdateUser(uint(userID.(float64)), updateData)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	// Only admin can update user status
	userRole, _ := c.Get("user_role")
	if userRole != string(models.RoleAdmin) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var statusData struct {
		IsActive bool `json:"is_active" binding:"required"`
	}

	if err := c.ShouldBindJSON(&statusData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := map[string]interface{}{
		"is_active": statusData.IsActive,
	}

	user, err := h.userService.UpdateUser(uint(id), updateData)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User status updated successfully",
		"user": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"is_active": user.IsActive,
		},
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	roleStr := c.DefaultQuery("role", "")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	users, total, err := h.userService.GetAllUsers(page, limit, models.UserRole(roleStr))
	if err != nil {
		h.logger.WithError(err).Error("Failed to fetch users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users, "total": total})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(uint(id), updateData)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		h.logger.WithError(err).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
