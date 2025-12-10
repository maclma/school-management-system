package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"school-management-system/internal/config"
	"school-management-system/internal/handlers"
	"school-management-system/internal/middleware"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/internal/service"
	"school-management-system/pkg/database"
)

var (
	testRouter *gin.Engine
	testDB     *gorm.DB
	authToken  string
)

func uniqueEmail(prefix string) string {
	return fmt.Sprintf("%s_%d@example.com", prefix, time.Now().UnixNano())
}

func clearDB() {
	// Attempt to clear users table between tests to avoid unique constraint conflicts
	if testDB != nil {
		testDB.Exec("DELETE FROM users")
		// vacuum for sqlite to release pages (no-op for others)
		testDB.Exec("VACUUM")
	}
}

func init() {
	// Suppress Gin debug output during tests
	gin.SetMode(gin.TestMode)

	// Set up SQLite in-memory database for testing
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_PATH", ":memory:")

	cfg, _ := config.LoadConfig()

	// Connect to test DB
	var err error
	testDB, err = database.ConnectDB(cfg)
	if err != nil {
		fmt.Printf("Failed to connect to test DB: %v\n", err)
		return
	}

	// Auto migrate models
	testDB.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Course{},
		&models.Enrollment{},
		&models.Grade{},
		&models.Attendance{},
	)

	// Setup router
	testRouter = gin.New()
	testRouter.Use(gin.Logger())
	testRouter.Use(gin.Recovery())

	// Initialize repositories and services
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiry)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	testRouter.POST("/api/auth/register", authHandler.Register)
	testRouter.POST("/api/auth/login", authHandler.Login)

	api := testRouter.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		api.GET("/profile", userHandler.GetProfile)
		api.PUT("/profile", userHandler.UpdateProfile)
		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:id", userHandler.GetUser)
	}
}

func TestRegister(t *testing.T) {
	clearDB()
	payload := map[string]interface{}{
		"first_name": "Test",
		"last_name":  "User",
		"email":      uniqueEmail("test"),
		"password":   "password123",
		"role":       "student",
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusCreated {
		t.Errorf("Expected status %d or %d, got %d", http.StatusOK, http.StatusCreated, w.Code)
	}
}

func TestLogin(t *testing.T) {
	// First register a user
	clearDB()
	email := uniqueEmail("login")
	registerPayload := map[string]interface{}{
		"first_name": "Login",
		"last_name":  "Test",
		"email":      email,
		"password":   "password123",
		"role":       "student",
	}
	registerBody, _ := json.Marshal(registerPayload)
	registerReq := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerW := httptest.NewRecorder()
	testRouter.ServeHTTP(registerW, registerReq)

	// Now login
	loginPayload := map[string]interface{}{
		"email":    email,
		"password": "password123",
	}
	loginBody, _ := json.Marshal(loginPayload)
	loginReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	testRouter.ServeHTTP(loginW, loginReq)

	if loginW.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, loginW.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(loginW.Body.Bytes(), &response)
	authToken, _ = response["token"].(string)

	if authToken == "" {
		t.Errorf("Expected token in response, got empty")
	}
}

func TestGetProfile(t *testing.T) {
	// Register and login first
	clearDB()
	email := uniqueEmail("profile")
	registerPayload := map[string]interface{}{
		"first_name": "Profile",
		"last_name":  "Test",
		"email":      email,
		"password":   "password123",
		"role":       "student",
	}
	registerBody, _ := json.Marshal(registerPayload)
	registerReq := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerW := httptest.NewRecorder()
	testRouter.ServeHTTP(registerW, registerReq)

	// Login
	loginPayload := map[string]interface{}{
		"email":    email,
		"password": "password123",
	}
	loginBody, _ := json.Marshal(loginPayload)
	loginReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	testRouter.ServeHTTP(loginW, loginReq)

	var loginResp map[string]interface{}
	json.Unmarshal(loginW.Body.Bytes(), &loginResp)
	token, _ := loginResp["token"].(string)

	// Get profile
	profileReq := httptest.NewRequest("GET", "/api/profile", nil)
	profileReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	profileW := httptest.NewRecorder()
	testRouter.ServeHTTP(profileW, profileReq)

	if profileW.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, profileW.Code)
	}

	var profile map[string]interface{}
	json.Unmarshal(profileW.Body.Bytes(), &profile)

	if gotEmail, ok := profile["email"].(string); !ok || gotEmail != email {
		t.Errorf("Expected email '%v', got %v", email, profile["email"])
	}
}

func TestUpdateProfile(t *testing.T) {
	// Register and login
	clearDB()
	email := uniqueEmail("update")
	registerPayload := map[string]interface{}{
		"first_name": "Update",
		"last_name":  "Test",
		"email":      email,
		"password":   "password123",
		"role":       "student",
	}
	registerBody, _ := json.Marshal(registerPayload)
	registerReq := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerW := httptest.NewRecorder()
	testRouter.ServeHTTP(registerW, registerReq)

	// Login
	loginPayload := map[string]interface{}{
		"email":    email,
		"password": "password123",
	}
	loginBody, _ := json.Marshal(loginPayload)
	loginReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	testRouter.ServeHTTP(loginW, loginReq)

	var loginResp map[string]interface{}
	json.Unmarshal(loginW.Body.Bytes(), &loginResp)
	token, _ := loginResp["token"].(string)

	// Update profile
	updatePayload := map[string]interface{}{
		"first_name": "UpdatedFirstName",
		"phone":      "1234567890",
	}
	updateBody, _ := json.Marshal(updatePayload)
	updateReq := httptest.NewRequest("PUT", "/api/profile", bytes.NewReader(updateBody))
	updateReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	updateReq.Header.Set("Content-Type", "application/json")
	updateW := httptest.NewRecorder()
	testRouter.ServeHTTP(updateW, updateReq)

	if updateW.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, updateW.Code)
	}

	var updated map[string]interface{}
	json.Unmarshal(updateW.Body.Bytes(), &updated)

	if firstName, ok := updated["first_name"].(string); !ok || firstName != "UpdatedFirstName" {
		t.Errorf("Expected first_name 'UpdatedFirstName', got %v", updated["first_name"])
	}
}

func TestUnauthorizedAccess(t *testing.T) {
	// Try to access protected endpoint without token
	req := httptest.NewRequest("GET", "/api/profile", nil)
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}
