package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleTeacher UserRole = "teacher"
	RoleStudent UserRole = "student"
	RoleParent  UserRole = "parent"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FirstName    string    `gorm:"size:100;not null" json:"first_name"`
	LastName     string    `gorm:"size:100;not null" json:"last_name"`
	Email        string    `gorm:"size:100;unique;not null" json:"email"`
	Password     string    `gorm:"size:255;not null" json:"-"`
	Phone        string    `gorm:"size:20" json:"phone"`
	Role         UserRole  `gorm:"type:user_role;not null" json:"role"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	Address      string    `gorm:"type:text" json:"address"`
	ProfileImage string    `json:"profile_image"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relations
	Student *Student `json:"student,omitempty"`
	Teacher *Teacher `json:"teacher,omitempty"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" && !isHashed(u.Password) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func isHashed(password string) bool {
	return len(password) == 60 && password[0] == '$'
}
