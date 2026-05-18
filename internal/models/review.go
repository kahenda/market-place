package models

import "time"

type Review struct {
	ID         string    `json:"id"`
	ReviewerID string    `json:"reviewer_id"`
	SellerID   string    `json:"seller_id"`
	ListingID  string    `json:"listing_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateReviewRequest struct {
	SellerID  string `json:"seller_id" binding:"required"`
	ListingID string `json:"listing_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment"`
}
