package models

import "time"

type Message struct {
	ID         string    `json:"id"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	ListingID  string    `json:"listing_id"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}

type SendMessageRequest struct {
	ReceiverID string `json:"receiver_id" binding:"required"`
	ListingID  string `json:"listing_id" binding:"required"`
	Body       string `json:"body" binding:"required"`
}
