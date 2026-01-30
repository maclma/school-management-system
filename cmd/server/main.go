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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	appLogger := logger.InitLogger(cfg)
	appLogger.Info("Starting School Management System")

	gin.SetMode(cfg.GinMode)

	// Measure DB connect time
	dbConnectStart := time.Now()
	db, err := database.ConnectDB(cfg)
	if err != nil {
		appLogger.Fatal("Failed to connect to database:", err)
	}
	appLogger.Infof("Database connected in %s", time.Since(dbConnectStart).String())
	defer database.CloseDB()

	// Auto migrate models with timing
	appLogger.Info("Running database migrations...")
	migStart := time.Now()
	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Course{},
		&models.Enrollment{},
		&models.Grade{},
		&models.Attendance{},
		&models.Assignment{},
		&models.AssignmentSubmission{},
		&models.SystemSetting{},
		&models.AuditLog{},
		&models.Notification{},
		&models.Announcement{},
		&models.Message{},
		&models.Payment{},
		&models.TimeTable{},
		&models.GradeTranscript{},
		&models.Backup{},
		&models.ImportBatch{},
		&models.AssignmentRubric{},
		&models.RubricScore{},
	)
	if err != nil {
		appLogger.Fatal("Failed to migrate database:", err)
	}
	appLogger.Infof("Database migrations completed in %s", time.Since(migStart).String())

	// Create admin user if needed
	adminStart := time.Now()
	createAdminUser(db, cfg, appLogger)
	appLogger.Infof("Admin user check completed in %s", time.Since(adminStart).String())

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	courseRepo := repository.NewCourseRepository()
	studentRepo := repository.NewStudentRepository()
	enrollmentRepo := repository.NewEnrollmentRepository()
	gradeRepo := repository.NewGradeRepository()
	attendanceRepo := repository.NewAttendanceRepository()
	teacherRepo := repository.NewTeacherRepository()
	assignmentRepo := repository.NewAssignmentRepository()
	assignmentSubmissionRepo := repository.NewAssignmentSubmissionRepository()

	// New feature repositories
	systemSettingRepo := repository.NewSystemSettingRepository()
	// auditLogRepo := repository.NewAuditLogRepository() // Used internally by middleware
	notificationRepo := repository.NewNotificationRepository()
	announcementRepo := repository.NewAnnouncementRepository()
	messageRepo := repository.NewMessageRepository()
	paymentRepo := repository.NewPaymentRepository()
	timetableRepo := repository.NewTimeTableRepository()
	gradeTranscriptRepo := repository.NewGradeTranscriptRepository()
	backupRepo := repository.NewBackupRepository()
	importBatchRepo := repository.NewImportBatchRepository()
	rubricRepo := repository.NewAssignmentRubricRepository()
	rubricScoreRepo := repository.NewRubricScoreRepository()

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiry)
	userService := service.NewUserService(userRepo)
	courseService := service.NewCourseService(courseRepo)
	studentService := service.NewStudentService(studentRepo)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	gradeService := service.NewGradeService(gradeRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo)
	teacherService := service.NewTeacherService(teacherRepo)
	assignmentService := service.NewAssignmentService(assignmentRepo)
	assignmentSubmissionService := service.NewAssignmentSubmissionService(assignmentSubmissionRepo)

	// New feature services
	systemSettingService := service.NewSystemSettingService(systemSettingRepo)
	// auditLogService := service.NewAuditLogService(auditLogRepo) // Used internally by middleware
	notificationService := service.NewNotificationService(notificationRepo)
	announcementService := service.NewAnnouncementService(announcementRepo)
	messageService := service.NewMessageService(messageRepo)
	paymentService := service.NewPaymentService(paymentRepo)
	timetableService := service.NewTimeTableService(timetableRepo)
	gradeTranscriptService := service.NewGradeTranscriptService(gradeTranscriptRepo)
	backupService := service.NewBackupService(backupRepo)
	importBatchService := service.NewImportBatchService(importBatchRepo)
	emailHost := os.Getenv("SMTP_HOST")
	emailPort := os.Getenv("SMTP_PORT")
	emailAddr := os.Getenv("SMTP_EMAIL")
	emailName := os.Getenv("SMTP_NAME")
	emailPass := os.Getenv("SMTP_PASS")
	emailService := service.NewEmailService(emailHost, emailPort, emailAddr, emailName, emailPass)
	searchService := service.NewSearchService(announcementRepo, paymentRepo, studentRepo)
	exportService := service.NewExportService(db)
	attendanceAutomationService := service.NewAttendanceAutomationService(emailService)
	gradeAutoCalculationService := service.NewGradeAutoCalculationService(gradeTranscriptService, emailService)

	// New feature handlers
	systemSettingHandler := handlers.NewSystemSettingHandler(systemSettingService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	messageHandler := handlers.NewMessageHandler(messageService)
	announcementHandler := handlers.NewAnnouncementHandler(announcementService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	timetableHandler := handlers.NewTimeTableHandler(timetableService)
	gradeTranscriptHandler := handlers.NewGradeTranscriptHandler(gradeTranscriptService)
	backupHandler := handlers.NewBackupHandler(backupService)
	importBatchHandler := handlers.NewImportBatchHandler(importBatchService)
	searchHandler := handlers.NewSearchHandler(searchService)
	exportHandler := handlers.NewExportHandler(exportService)
	attendanceAutomationHandler := handlers.NewAttendanceAutomationHandler(attendanceAutomationService)
	gradeAutoCalcHandler := handlers.NewGradeAutoCalcHandler(gradeAutoCalculationService)
	rubricHandler := handlers.NewRubricHandler(rubricRepo, rubricScoreRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	courseHandler := handlers.NewCourseHandler(courseService)
	studentHandler := handlers.NewStudentHandler(studentService)
	enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentService, studentService)
	gradeHandler := handlers.NewGradeHandler(gradeService, studentService)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService, studentService)
	adminHandler := handlers.NewAdminHandler(userService, courseService, studentService)
	teacherHandler := handlers.NewTeacherHandler(teacherService)
	assignmentHandler := handlers.NewAssignmentHandler(assignmentService, assignmentSubmissionService, studentService)

	// Setup router with comprehensive middleware
	router := gin.New()

	// Recovery and logging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Request tracking and security
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ValidationMiddleware())
	router.Use(middleware.MaxBodySizeMiddleware(10 * 1024 * 1024)) // 10MB limit

	// Rate limiting on public routes
	router.Use(middleware.APIRateLimit())

	// Health endpoint (public, no auth required)
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "timestamp": time.Now().Unix(), "service": "school-management-system"})
	})

	// Public routes
	public := router.Group("/api/auth")
	public.Use(middleware.AuthRateLimit()) // Stricter rate limiting for auth
	{
		public.POST("/login", authHandler.Login)
		public.POST("/register", authHandler.Register)
	}

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(authService))
	{
		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.PATCH("/users/:id/status", userHandler.UpdateUserStatus)

		api.GET("/profile", userHandler.GetProfile)
		api.PUT("/profile", userHandler.UpdateProfile)

		api.POST("/courses", courseHandler.CreateCourse)
		api.GET("/courses", courseHandler.GetAllCourses)
		api.GET("/courses/:id", courseHandler.GetCourse)
		api.PUT("/courses/:id", courseHandler.UpdateCourse)
		api.DELETE("/courses/:id", courseHandler.DeleteCourse)
		api.GET("/courses/by-department", courseHandler.GetCoursesByDepartment)

		api.POST("/students", studentHandler.CreateStudent)
		api.GET("/students", studentHandler.GetAllStudents)
		api.GET("/students/:id", studentHandler.GetStudent)
		api.PUT("/students/:id", studentHandler.UpdateStudent)
		api.DELETE("/students/:id", studentHandler.DeleteStudent)

		api.POST("/enrollments", enrollmentHandler.EnrollStudent)
		api.GET("/enrollments/:id", enrollmentHandler.GetEnrollment)
		api.PUT("/enrollments/:id/status", enrollmentHandler.UpdateEnrollmentStatus)
		api.DELETE("/enrollments/:id", enrollmentHandler.RemoveEnrollment)
		api.GET("/enrollments/by-student/:studentId", enrollmentHandler.GetStudentEnrollments)
		api.GET("/enrollments/by-course/:courseId", enrollmentHandler.GetCourseEnrollments)

		api.POST("/grades", gradeHandler.RecordGrade)
		api.GET("/grades/:id", gradeHandler.GetGrade)
		api.PUT("/grades/:id", gradeHandler.UpdateGrade)
		api.DELETE("/grades/:id", gradeHandler.DeleteGrade)
		api.GET("/grades/by-student/:studentId", gradeHandler.GetStudentGrades)
		api.GET("/grades/by-course/:courseId", gradeHandler.GetCourseGrades)
		api.GET("/grades/average/:studentId", gradeHandler.GetAverageGrade)

		api.POST("/attendance", attendanceHandler.RecordAttendance)
		api.GET("/attendance/:id", attendanceHandler.GetAttendance)
		api.PUT("/attendance/:id", attendanceHandler.UpdateAttendance)
		api.DELETE("/attendance/:id", attendanceHandler.DeleteAttendance)
		api.GET("/attendance/by-student/:studentId", attendanceHandler.GetStudentAttendance)
		api.GET("/attendance/by-course/:courseId", attendanceHandler.GetCourseAttendance)
		api.GET("/attendance/student-course/:studentId/:courseId", attendanceHandler.GetStudentCourseAttendance)
		api.GET("/attendance/stats/:studentId/:courseId", attendanceHandler.GetAttendanceStats)

		// Assignments
		api.POST("/assignments", assignmentHandler.CreateAssignment)
		api.GET("/assignments/:id", assignmentHandler.GetAssignment)
		api.PUT("/assignments/:id", assignmentHandler.UpdateAssignment)
		api.DELETE("/assignments/:id", assignmentHandler.DeleteAssignment)
		api.GET("/assignments/course/:course_id", assignmentHandler.GetAssignmentsByCourse)

		// Assignment submissions
		api.POST("/assignments/submit", assignmentHandler.SubmitAssignment)
		api.GET("/submissions/assignment/:assignment_id", assignmentHandler.GetSubmissionsByAssignment)
		api.PUT("/submissions/:submission_id/grade", assignmentHandler.GradeSubmission)

		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware(models.RoleAdmin))
		{
			admin.DELETE("/users/:id", userHandler.DeleteUser)
			admin.POST("/users", adminHandler.CreateUserAdmin)
			admin.GET("/users", adminHandler.GetAllUsersAdmin)
			admin.GET("/dashboard", adminHandler.GetDashboardStats)
			admin.GET("/health", adminHandler.SystemHealth)

			admin.GET("/enrollments", enrollmentHandler.GetAllEnrollments)
			admin.POST("/enrollments/:id/approve", enrollmentHandler.ApproveEnrollment)
			admin.POST("/enrollments/:id/reject", enrollmentHandler.RejectEnrollment)

			admin.POST("/teachers", teacherHandler.CreateTeacher)
			admin.GET("/teachers", teacherHandler.GetAllTeachers)
			admin.GET("/teachers/:id", teacherHandler.GetTeacher)
			admin.PUT("/teachers/:id", teacherHandler.UpdateTeacher)
			admin.DELETE("/teachers/:id", teacherHandler.DeleteTeacher)
			admin.GET("/teachers/by-department", teacherHandler.GetTeachersByDepartment)
			admin.GET("/teachers/:id/courses", teacherHandler.GetTeacherCourses)
		}

		teacher := api.Group("/teacher")
		teacher.Use(middleware.RoleMiddleware(models.RoleTeacher))
		{
			teacher.POST("/grades", gradeHandler.RecordGrade)
			teacher.PUT("/grades/:id", gradeHandler.UpdateGrade)
			teacher.POST("/attendance", attendanceHandler.RecordAttendance)
			teacher.PUT("/attendance/:id", attendanceHandler.UpdateAttendance)

			teacher.GET("/assignments", assignmentHandler.GetAssignmentsByTeacher)
			teacher.GET("/submissions/assignment/:assignment_id", assignmentHandler.GetSubmissionsByAssignment)
			teacher.PUT("/submissions/:submission_id/grade", assignmentHandler.GradeSubmission)

			// Notifications
			api.POST("/notifications", notificationHandler.Create)
			api.GET("/notifications", notificationHandler.GetMyNotifications)
			api.GET("/notifications/unread", notificationHandler.GetUnread)
			api.PUT("/notifications/:id/read", notificationHandler.MarkAsRead)
			api.PUT("/notifications/mark-all-read", notificationHandler.MarkAllAsRead)
			api.DELETE("/notifications/:id", notificationHandler.Delete)

			// Messages
			api.POST("/messages", messageHandler.SendMessage)
			api.GET("/messages/inbox", messageHandler.GetInbox)
			api.GET("/messages/conversation/:user_id", messageHandler.GetConversation)
			api.GET("/messages/unread", messageHandler.CountUnread)
			api.PUT("/messages/:id/read", messageHandler.MarkAsRead)

			// Announcements
			api.GET("/announcements", announcementHandler.GetAll)
			api.GET("/announcements/active", announcementHandler.GetActive)
			api.POST("/announcements", announcementHandler.Create)
			api.PUT("/announcements/:id", announcementHandler.Update)
			api.DELETE("/announcements/:id", announcementHandler.Delete)

			// Advanced Search
			api.GET("/search/announcements", searchHandler.SearchAnnouncements)
			api.GET("/search/payments", searchHandler.SearchPayments)
			api.GET("/search/students", searchHandler.SearchStudents)
			api.GET("/search/grades", searchHandler.SearchGradesByRange)
			api.GET("/search/overdue-payments", searchHandler.SearchOverduePayments)

			// CSV Exports
			api.GET("/export/payments", exportHandler.ExportPaymentsCSV)
			api.GET("/export/grades", exportHandler.ExportGradesCSV)
			api.GET("/export/attendance", exportHandler.ExportAttendanceCSV)
			api.GET("/export/transcript/:student_id", exportHandler.ExportStudentTranscript)
			api.GET("/export/enrollments", exportHandler.ExportEnrollments)

			// Attendance Automation
			api.GET("/attendance/stats/course/:course_id", attendanceAutomationHandler.GetAttendanceStats)
			api.GET("/attendance/percentage/:student_id/:course_id", attendanceAutomationHandler.GetStudentAttendancePercentage)
			api.POST("/attendance/check-low", attendanceAutomationHandler.CheckLowAttendance)
			api.GET("/attendance/low/:threshold", attendanceAutomationHandler.GetStudentsWithLowAttendance)
			api.GET("/attendance/report/:course_id", attendanceAutomationHandler.GetAttendanceReport)

			// Grade Auto-Calculation
			api.POST("/grades/auto", gradeAutoCalcHandler.RecordGradeWithAutoCalc)
			api.GET("/grades/course-average/:course_id", gradeAutoCalcHandler.GetCourseAverage)
			api.GET("/grades/distribution/:course_id", gradeAutoCalcHandler.GetGradeDistribution)
			api.GET("/grades/student-stats/:student_id", gradeAutoCalcHandler.GetStudentGradeStats)

			// Rubrics
			api.POST("/rubrics", rubricHandler.CreateRubric)
			api.GET("/rubrics/:id", rubricHandler.GetRubric)
			api.GET("/rubrics/assignment/:assignment_id", rubricHandler.GetRubricsByAssignment)
			api.PUT("/rubrics/:id", rubricHandler.UpdateRubric)
			api.DELETE("/rubrics/:id", rubricHandler.DeleteRubric)
			api.POST("/rubrics/score", rubricHandler.ScoreSubmission)
			api.GET("/rubrics/score/:submission_id", rubricHandler.GetSubmissionScore)

			// Timetable
			api.GET("/timetable", timetableHandler.GetAll)
			api.GET("/timetable/course/:course_id", timetableHandler.GetByCourseID)
			api.GET("/timetable/teacher/:teacher_id", timetableHandler.GetByTeacherID)
			api.GET("/timetable/day/:day", timetableHandler.GetByDay)
			api.POST("/timetable", timetableHandler.Create)
			api.PUT("/timetable/:id", timetableHandler.Update)
			api.DELETE("/timetable/:id", timetableHandler.Delete)

			// Grade Transcripts
			api.GET("/transcripts/student/:student_id", gradeTranscriptHandler.GetByStudentID)
			api.GET("/transcripts/latest/:student_id", gradeTranscriptHandler.GetLatest)
			api.GET("/transcripts/gpa/:student_id", gradeTranscriptHandler.GetGPA)

			// Payments
			api.POST("/payments", paymentHandler.Create)
			api.GET("/payments/student/:student_id", paymentHandler.GetByStudent)
			api.GET("/payments", paymentHandler.GetAll)
			api.PUT("/payments/:id", paymentHandler.Update)
			api.GET("/payments/balance/:student_id", paymentHandler.GetStudentBalance)

			// System Settings (admin only)
			admin.GET("/settings", systemSettingHandler.GetAll)
			admin.GET("/settings/:key", systemSettingHandler.GetByKey)
			admin.POST("/settings", systemSettingHandler.Create)
			admin.PUT("/settings/:id", systemSettingHandler.Update)
			admin.DELETE("/settings/:id", systemSettingHandler.Delete)

			// Backups (admin only)
			admin.GET("/backups", backupHandler.GetAll)
			admin.GET("/backups/latest", backupHandler.GetLatest)
			admin.GET("/backups/:id", backupHandler.GetByID)
			admin.DELETE("/backups/:id", backupHandler.Delete)

			// Import batches (admin only)
			admin.GET("/imports", importBatchHandler.GetAll)
			admin.GET("/imports/:id", importBatchHandler.GetByID)
			admin.GET("/imports/status/:status", importBatchHandler.GetByStatus)
			admin.DELETE("/imports/:id", importBatchHandler.Delete)
		}

		student := api.Group("/student")
		student.Use(middleware.RoleMiddleware(models.RoleStudent))
		{
			student.GET("/enrollments", enrollmentHandler.GetMyEnrollments)
			student.GET("/grades", gradeHandler.GetMyGrades)
			student.GET("/attendance", attendanceHandler.GetMyAttendance)

			student.GET("/assignments/submissions", assignmentHandler.GetSubmissionsByStudent)
		}
	}

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		appLogger.Infof("Server starting on port %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatalf("Failed to start server: %v", err)
		}
	}()

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

	// Use ADMIN_PASSWORD/ADMIN_EMAIL env vars for creation or reset in non-production
	adminPass := os.Getenv("ADMIN_PASSWORD")
	if adminPass == "" {
		adminPass = "admin123"
	}
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		adminEmail = "admin@school.com"
	}

	if count == 0 {
		adminUser := &models.User{
			FirstName:   "Admin",
			LastName:    "User",
			Email:       adminEmail,
			Password:    adminPass,
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
			if cfg.AppEnv != "production" {
				logger.Infof("Admin credentials -> email: %s password: %s", adminUser.Email, adminPass)
			}
		}
		return
	}

	// If admin exists but ADMIN_PASSWORD env provided and not production, update password and log
	if envPass := os.Getenv("ADMIN_PASSWORD"); envPass != "" && cfg.AppEnv != "production" {
		var admin models.User
		q := db.Model(&models.User{}).Where("role = ?", models.RoleAdmin)
		if adminEmail != "" {
			q = q.Where("email = ?", adminEmail)
		}
		if err := q.First(&admin).Error; err == nil {
			admin.Password = envPass
			if err := db.Save(&admin).Error; err == nil {
				logger.Infof("Admin password reset for %s", admin.Email)
				logger.Infof("Admin credentials -> email: %s password: %s", admin.Email, envPass)
			} else {
				logger.Warnf("Failed to update admin password: %v", err)
			}
		}
	}
}
