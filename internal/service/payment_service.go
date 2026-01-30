package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type PaymentService interface {
	Create(payment *models.Payment) error
	GetByID(id uint) (*models.Payment, error)
	GetByStudentID(studentID uint, page, limit int) ([]models.Payment, int64, error)
	GetByStatus(status string, page, limit int) ([]models.Payment, int64, error)
	GetAll(page, limit int) ([]models.Payment, int64, error)
	Update(payment *models.Payment) error
	Delete(id uint) error
	GetStudentBalance(studentID uint) (float64, error)
	GetTotalRevenue(status string) (float64, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) Create(payment *models.Payment) error {
	return s.repo.Create(payment)
}

func (s *paymentService) GetByID(id uint) (*models.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *paymentService) GetByStudentID(studentID uint, page, limit int) ([]models.Payment, int64, error) {
	return s.repo.FindByStudentID(studentID, page, limit)
}

func (s *paymentService) GetByStatus(status string, page, limit int) ([]models.Payment, int64, error) {
	return s.repo.FindByStatus(status, page, limit)
}

func (s *paymentService) GetAll(page, limit int) ([]models.Payment, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *paymentService) Update(payment *models.Payment) error {
	return s.repo.Update(payment)
}

func (s *paymentService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *paymentService) GetStudentBalance(studentID uint) (float64, error) {
	return s.repo.SumByStudent(studentID)
}

func (s *paymentService) GetTotalRevenue(status string) (float64, error) {
	return s.repo.SumByStatus(status)
}
