package repository

import (
	"school-management-system/internal/models"
	"school-management-system/pkg/database"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	FindByID(id uint) (*models.Payment, error)
	FindByStudentID(studentID uint, page, limit int) ([]models.Payment, int64, error)
	FindByStatus(status string, page, limit int) ([]models.Payment, int64, error)
	FindAll(page, limit int) ([]models.Payment, int64, error)
	Update(payment *models.Payment) error
	Delete(id uint) error
	SumByStudent(studentID uint) (float64, error)
	SumByStatus(status string) (float64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{db: database.DB}
}

func (r *paymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) FindByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Preload("Student").First(&payment, id).Error
	return &payment, err
}

func (r *paymentRepository) FindByStudentID(studentID uint, page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("student_id = ?", studentID).Count(&total).
		Preload("Student").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error
	return payments, total, err
}

func (r *paymentRepository) FindByStatus(status string, page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64
	offset := (page - 1) * limit
	err := r.db.Where("status = ?", status).Count(&total).
		Preload("Student").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error
	return payments, total, err
}

func (r *paymentRepository) FindAll(page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64
	offset := (page - 1) * limit
	err := r.db.Count(&total).
		Preload("Student").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&payments).Error
	return payments, total, err
}

func (r *paymentRepository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}

func (r *paymentRepository) SumByStudent(studentID uint) (float64, error) {
	var total float64
	err := r.db.Model(&models.Payment{}).
		Where("student_id = ? AND status = ?", studentID, "paid").
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&total)
	return total, err
}

func (r *paymentRepository) SumByStatus(status string) (float64, error) {
	var total float64
	err := r.db.Model(&models.Payment{}).
		Where("status = ?", status).
		Select("COALESCE(SUM(amount), 0)").
		Row().
		Scan(&total)
	return total, err
}
