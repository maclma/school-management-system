package main

import (
	"fmt"
	"log"
	"os"
	"school-management-system/internal/config"
	"school-management-system/internal/models"
	"school-management-system/pkg/database"
)

func main() {
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("DB_PATH", "school.db")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	var user models.User
	if err := db.Where("email = ?", "admin@school.com").First(&user).Error; err != nil {
		log.Fatal("Admin user not found:", err)
	}

	user.Password = "admin123"
	if err := db.Save(&user).Error; err != nil {
		log.Fatal("Failed to update password:", err)
	}

	fmt.Println("Admin password reset to admin123")
}
