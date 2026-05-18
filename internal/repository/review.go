package repository

import (
	"database/sql"
	"github.com/kahenda/marketplace/internal/models"
)

type ReviewRepository struct {
	DB *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	query := `
		INSERT INTO reviews (reviewer_id, seller_id, listing_id, rating, comment)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		review.ReviewerID,
		review.SellerID,
		review.ListingID,
		review.Rating,
		review.Comment,
	).Scan(&review.ID, &review.CreatedAt)
}

func (r *ReviewRepository) GetReviewsBySeller(sellerID string) ([]models.Review, error) {
	query := `
		SELECT id, reviewer_id, seller_id, listing_id, rating, comment, created_at
		FROM reviews WHERE seller_id = $1
		ORDER BY created_at DESC`
	rows, err := r.DB.Query(query, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var rv models.Review
		if err := rows.Scan(&rv.ID, &rv.ReviewerID, &rv.SellerID, &rv.ListingID, &rv.Rating, &rv.Comment, &rv.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, rv)
	}
	return reviews, nil
}
