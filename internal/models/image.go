package models

import "time"

type Image struct {
	ID        string    `json:"id"`
	ListingID string    `json:"listing_id"`
	URL       string    `json:"url"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}
