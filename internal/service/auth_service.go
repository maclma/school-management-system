package service

import (
	"errors"
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
	"school-management-system/pkg/logger"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type AuthService interface {
	Login(email, password string) (string, *models.User, error)
	Register(user *models.User) error
	GenerateToken(user *models.User) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	RefreshToken(tokenString string) (string, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
	jwtExpiry int
	logger    *logrus.Logger
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string, jwtExpiry int) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
		logger:    logger.GetLogger(),
	}
}

func (s *authService) Login(email, password string) (string, *models.User, error) {
	s.logger.WithField("email", email).Info("Login attempt")

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		s.logger.WithError(err).WithField("email", email).Warn("User not found")
		return "", nil, errors.New("invalid credentials")
	}

	if !user.CheckPassword(password) {
		s.logger.WithField("email", email).Warn("Invalid password")
		return "", nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		s.logger.WithField("email", email).Warn("Inactive account login attempt")
		return "", nil, errors.New("account is inactive")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate token")
		return "", nil, errors.New("failed to generate token")
	}

	s.logger.WithField("email", email).Info("Login successful")
	return token, user, nil
}

func (s *authService) Register(user *models.User) error {
	// Normalize email
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		s.logger.WithField("email", user.Email).Warn("Registration attempt with existing email")
		return errors.New("email already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.WithError(err).WithField("email", user.Email).Error("Failed checking existing user")
		return errors.New("failed to register user")
	}

	err = s.userRepo.Create(user)
	if err != nil {
		s.logger.WithError(err).WithField("email", user.Email).Error("Failed to register user")
		return errors.New("failed to register user")
	}

	s.logger.WithField("email", user.Email).WithField("role", user.Role).Info("User registered successfully")
	return nil
}

func (s *authService) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * time.Duration(s.jwtExpiry)).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		s.logger.WithError(err).Warn("Token validation failed")
		return nil, err
	}

	return token, nil
}

func (s *authService) RefreshToken(tokenString string) (string, error) {
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token claims")
	}

	userID := uint(claims["user_id"].(float64))
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	return s.GenerateToken(user)
}
