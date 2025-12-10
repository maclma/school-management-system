package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, userData map[string]interface{}) (*models.User, error)
	DeleteUser(id uint) error
	GetAllUsers(page, limit int, role models.UserRole) ([]models.User, int64, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, userData map[string]interface{}) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if firstName, ok := userData["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := userData["last_name"].(string); ok {
		user.LastName = lastName
	}
	if phone, ok := userData["phone"].(string); ok {
		user.Phone = phone
	}

	err = s.userRepo.Update(user)
	return user, err
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *userService) GetAllUsers(page, limit int, role models.UserRole) ([]models.User, int64, error) {
	return s.userRepo.FindAll(page, limit, role)
}
