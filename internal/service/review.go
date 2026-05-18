package service

import (
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/repository"
)

type ReviewService struct {
	Repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) CreateReview(reviewerID string, req *models.CreateReviewRequest) (*models.Review, error) {
	review := &models.Review{
		ReviewerID: reviewerID,
		SellerID:   req.SellerID,
		ListingID:  req.ListingID,
		Rating:     req.Rating,
		Comment:    req.Comment,
	}
	if err := s.Repo.CreateReview(review); err != nil {
		return nil, err
	}
	return review, nil
}

func (s *ReviewService) GetReviewsBySeller(sellerID string) ([]models.Review, error) {
	return s.Repo.GetReviewsBySeller(sellerID)
}
