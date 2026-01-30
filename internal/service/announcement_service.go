package service

import (
	"school-management-system/internal/models"
	"school-management-system/internal/repository"
)

type AnnouncementService interface {
	Create(announcement *models.Announcement) error
	GetByID(id uint) (*models.Announcement, error)
	GetAll(page, limit int) ([]models.Announcement, int64, error)
	GetActive(page, limit int) ([]models.Announcement, int64, error)
	GetByAudience(audience string, page, limit int) ([]models.Announcement, int64, error)
	Update(announcement *models.Announcement) error
	Delete(id uint) error
}

type announcementService struct {
	repo repository.AnnouncementRepository
}

func NewAnnouncementService(repo repository.AnnouncementRepository) AnnouncementService {
	return &announcementService{repo: repo}
}

func (s *announcementService) Create(announcement *models.Announcement) error {
	return s.repo.Create(announcement)
}

func (s *announcementService) GetByID(id uint) (*models.Announcement, error) {
	return s.repo.FindByID(id)
}

func (s *announcementService) GetAll(page, limit int) ([]models.Announcement, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *announcementService) GetActive(page, limit int) ([]models.Announcement, int64, error) {
	return s.repo.FindActive(page, limit)
}

func (s *announcementService) GetByAudience(audience string, page, limit int) ([]models.Announcement, int64, error) {
	return s.repo.FindByAudience(audience, page, limit)
}

func (s *announcementService) Update(announcement *models.Announcement) error {
	return s.repo.Update(announcement)
}

func (s *announcementService) Delete(id uint) error {
	return s.repo.Delete(id)
}
