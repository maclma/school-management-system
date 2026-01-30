package database

import (
	"fmt"
	"log"
	"school-management-system/internal/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	glebarez "github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Configure GORM logger based on environment
	gormLogger := logger.Default
	if cfg.AppEnv == "production" {
		gormLogger = gormLogger.LogMode(logger.Warn)
	}

	// Retry connection (useful when starting app before database)
	maxRetries := 5
	start := time.Now()
	for i := 0; i < maxRetries; i++ {
		log.Printf("DB connect attempt %d/%d", i+1, maxRetries)
		if cfg.DBDriver == "sqlite" || cfg.DBDriver == "sqlite3" {
			db, err = gorm.Open(glebarez.Open(cfg.DBPath), &gorm.Config{
				Logger: gormLogger,
			})
		} else {
			dsn := fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Kolkata",
				cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: gormLogger,
			})
		}

		if err == nil {
			elapsed := time.Since(start)
			log.Printf("DB connected after %s (attempt %d)", elapsed.String(), i+1)
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Create indexes for performance optimization
	if err := CreateIndexes(db); err != nil {
		log.Printf("Warning: Failed to create indexes: %v", err)
	} else {
		log.Println("Database indexes created successfully")
	}

	// Optimize query settings
	OptimizeQueries(db)

	DB = db
	log.Println("Database connected successfully")
	return db, nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
