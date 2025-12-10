package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"school-management-system/internal/config"
	"school-management-system/internal/handlers"
	"school-management-system/internal/middleware"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/internal/service"
	"school-management-system/pkg/database"
	"school-management-system/pkg/logger"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env if present (local development)
	_ = godotenv.Load()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize logger
	appLogger := logger.InitLogger(cfg)
	appLogger.Info("Starting School Management System")

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Initialize database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		appLogger.Fatal("Failed to connect to database:", err)
	}
	defer database.CloseDB()

	// Auto migrate models
	appLogger.Info("Running database migrations...")
	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Course{},
		&models.Enrollment{},
		&models.Grade{},
		&models.Attendance{},
	)
	if err != nil {
		appLogger.Fatal("Failed to migrate database:", err)
	}
	appLogger.Info("Database migrations completed")

	// Create admin user if not exists (optional)
	createAdminUser(db, cfg, appLogger)

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	courseRepo := repository.NewCourseRepository()
	studentRepo := repository.NewStudentRepository()
	enrollmentRepo := repository.NewEnrollmentRepository()
	gradeRepo := repository.NewGradeRepository()
	attendanceRepo := repository.NewAttendanceRepository()

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiry)
	userService := service.NewUserService(userRepo)
	courseService := service.NewCourseService(courseRepo)
	studentService := service.NewStudentService(studentRepo)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	gradeService := service.NewGradeService(gradeRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	courseHandler := handlers.NewCourseHandler(courseService)
	studentHandler := handlers.NewStudentHandler(studentService)
	enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentService)
	gradeHandler := handlers.NewGradeHandler(gradeService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService)
	adminHandler := handlers.NewAdminHandler(userService, courseService, studentService)

	// Setup router with CORS
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health check endpoint (register before wildcard/static routes)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
			"service":   "school-management-system",
		})
	})

	// Serve frontend static files via NoRoute handler (avoid wildcard route conflict)
	// If a file exists in ./frontend matching the path serve it, otherwise fall back to index.html
	router.NoRoute(func(c *gin.Context) {
		// Direct root -> serve index
		if c.Request.URL.Path == "/" || c.Request.URL.Path == "" {
			http.ServeFile(c.Writer, c.Request, "./frontend/index.html")
			return
		}

		// Try to serve the exact file in ./frontend/dist (built SPA) first
		fpDist := "./frontend/dist" + c.Request.URL.Path
		if _, err := os.Stat(fpDist); err == nil {
			http.ServeFile(c.Writer, c.Request, fpDist)
			return
		}

		// Then try to serve from ./frontend (dev static files)
		fp := "./frontend" + c.Request.URL.Path
		if _, err := os.Stat(fp); err == nil {
			http.ServeFile(c.Writer, c.Request, fp)
			return
		}

		// Fallback to built index.html if exists, otherwise dev index
		if _, err := os.Stat("./frontend/dist/index.html"); err == nil {
			http.ServeFile(c.Writer, c.Request, "./frontend/dist/index.html")
			return
		}

		http.ServeFile(c.Writer, c.Request, "./frontend/index.html")
	}) // Public routes
	router.POST("/api/auth/login", authHandler.Login)
	router.POST("/api/auth/register", authHandler.Register)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		// User routes
		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.PATCH("/users/:id/status", userHandler.UpdateUserStatus)

		// Profile routes
		api.GET("/profile", userHandler.GetProfile)
		api.PUT("/profile", userHandler.UpdateProfile)

		// Course routes
		api.POST("/courses", courseHandler.CreateCourse)
		api.GET("/courses", courseHandler.GetAllCourses)
		api.GET("/courses/:id", courseHandler.GetCourse)
		api.PUT("/courses/:id", courseHandler.UpdateCourse)
		api.DELETE("/courses/:id", courseHandler.DeleteCourse)
		api.GET("/courses/by-department", courseHandler.GetCoursesByDepartment)

		// Student routes
		api.POST("/students", studentHandler.CreateStudent)
		api.GET("/students", studentHandler.GetAllStudents)
		api.GET("/students/:id", studentHandler.GetStudent)
		api.PUT("/students/:id", studentHandler.UpdateStudent)
		api.DELETE("/students/:id", studentHandler.DeleteStudent)

		// Enrollment routes
		api.POST("/enrollments", enrollmentHandler.EnrollStudent)
		api.GET("/enrollments/:id", enrollmentHandler.GetEnrollment)
		api.PUT("/enrollments/:id/status", enrollmentHandler.UpdateEnrollmentStatus)
		api.DELETE("/enrollments/:id", enrollmentHandler.RemoveEnrollment)
		api.GET("/enrollments/by-student/:studentId", enrollmentHandler.GetStudentEnrollments)
		api.GET("/enrollments/by-course/:courseId", enrollmentHandler.GetCourseEnrollments)

		// Grade routes
		api.POST("/grades", gradeHandler.RecordGrade)
		api.GET("/grades/:id", gradeHandler.GetGrade)
		api.PUT("/grades/:id", gradeHandler.UpdateGrade)
		api.DELETE("/grades/:id", gradeHandler.DeleteGrade)
		api.GET("/grades/by-student/:studentId", gradeHandler.GetStudentGrades)
		api.GET("/grades/by-course/:courseId", gradeHandler.GetCourseGrades)
		api.GET("/grades/average/:studentId", gradeHandler.GetAverageGrade)

		// Attendance routes
		api.POST("/attendance", attendanceHandler.RecordAttendance)
		api.GET("/attendance/:id", attendanceHandler.GetAttendance)
		api.PUT("/attendance/:id", attendanceHandler.UpdateAttendance)
		api.DELETE("/attendance/:id", attendanceHandler.DeleteAttendance)
		api.GET("/attendance/by-student/:studentId", attendanceHandler.GetStudentAttendance)
		api.GET("/attendance/by-course/:courseId", attendanceHandler.GetCourseAttendance)
		api.GET("/attendance/student-course/:studentId/:courseId", attendanceHandler.GetStudentCourseAttendance)
		api.GET("/attendance/stats/:studentId/:courseId", attendanceHandler.GetAttendanceStats)

		// Admin only routes
		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware(models.RoleAdmin))
		{
			admin.DELETE("/users/:id", userHandler.DeleteUser)
			admin.POST("/users", adminHandler.CreateUserAdmin)
			admin.GET("/users", adminHandler.GetAllUsersAdmin)
			admin.GET("/dashboard", adminHandler.GetDashboardStats)
			admin.GET("/health", adminHandler.SystemHealth)

			// Enrollment approval routes
			admin.GET("/enrollments", enrollmentHandler.GetAllEnrollments)
			admin.POST("/enrollments/:id/approve", enrollmentHandler.ApproveEnrollment)
			admin.POST("/enrollments/:id/reject", enrollmentHandler.RejectEnrollment)
		}

		// Teacher routes
		teacher := api.Group("/teacher")
		teacher.Use(middleware.RoleMiddleware(models.RoleTeacher))
		{
			teacher.POST("/grades", gradeHandler.RecordGrade)
			teacher.PUT("/grades/:id", gradeHandler.UpdateGrade)
			teacher.POST("/attendance", attendanceHandler.RecordAttendance)
			teacher.PUT("/attendance/:id", attendanceHandler.UpdateAttendance)
		}

		// Student routes
		student := api.Group("/student")
		student.Use(middleware.RoleMiddleware(models.RoleStudent))
		{
			student.GET("/enrollments", enrollmentHandler.GetStudentEnrollments)
			student.GET("/grades", gradeHandler.GetStudentGrades)
			student.GET("/attendance", attendanceHandler.GetStudentAttendance)
		}
	}

	// Setup server with graceful shutdown
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		appLogger.Infof("Server starting on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	appLogger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		appLogger.Fatal("Server forced to shutdown:", err)
	}

	appLogger.Info("Server exited properly")
}

func createAdminUser(db *gorm.DB, cfg *config.Config, logger *logrus.Logger) {
	var count int64
	db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count)

	if count == 0 {
		adminUser := &models.User{
			FirstName:   "Admin",
			LastName:    "User",
			Email:       "admin@school.com",
			Password:    "admin123", // Change this in production
			Phone:       "1234567890",
			Role:        models.RoleAdmin,
			DateOfBirth: time.Now().AddDate(-30, 0, 0),
			Address:     "School Address",
			IsActive:    true,
		}

		if err := db.Create(adminUser).Error; err != nil {
			logger.Warnf("Failed to create admin user: %v", err)
		} else {
			logger.Info("Admin user created successfully")
		}
	}
}
