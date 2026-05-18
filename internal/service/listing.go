package service

import (
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/repository"
)

type ListingService struct {
	Repo *repository.ListingRepository
}

func NewListingService(repo *repository.ListingRepository) *ListingService {
	return &ListingService{Repo: repo}
}

func (s *ListingService) CreateListing(userID string, req *models.CreateListingRequest) (*models.Listing, error) {
	listing := &models.Listing{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Gender:      req.Gender,
		Size:        req.Size,
		Condition:   req.Condition,
		Area:        req.Area,
	}
	if err := s.Repo.CreateListing(listing); err != nil {
		return nil, err
	}
	return listing, nil
}

func (s *ListingService) GetListings(category, gender, size, condition, area string) ([]models.Listing, error) {
	return s.Repo.GetListings(category, gender, size, condition, area)
}
