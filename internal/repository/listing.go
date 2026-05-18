package repository

import (
	"database/sql"
	"github.com/kahenda/marketplace/internal/models"
)

type ListingRepository struct {
	DB *sql.DB
}

func NewListingRepository(db *sql.DB) *ListingRepository {
	return &ListingRepository{DB: db}
}

func (r *ListingRepository) CreateListing(listing *models.Listing) error {
	query := `
		INSERT INTO listings (user_id, title, description, price, category, status)
		VALUES ($1, $2, $3, $4, $5, 'active')
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		listing.UserID,
		listing.Title,
		listing.Description,
		listing.Price,
		listing.Category,
	).Scan(&listing.ID, &listing.CreatedAt)
}

func (r *ListingRepository) GetAllListings() ([]models.Listing, error) {
	query := `SELECT id, user_id, title, description, price, category, status, created_at FROM listings WHERE status = 'active' ORDER BY created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []models.Listing
	for rows.Next() {
		var l models.Listing
		if err := rows.Scan(&l.ID, &l.UserID, &l.Title, &l.Description, &l.Price, &l.Category, &l.Status, &l.CreatedAt); err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}

func (r *ListingRepository) GetListingsByCategory(category string) ([]models.Listing, error) {
	query := `SELECT id, user_id, title, description, price, category, status, created_at FROM listings WHERE status = 'active' AND category = $1 ORDER BY created_at DESC`
	rows, err := r.DB.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []models.Listing
	for rows.Next() {
		var l models.Listing
		if err := rows.Scan(&l.ID, &l.UserID, &l.Title, &l.Description, &l.Price, &l.Category, &l.Status, &l.CreatedAt); err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}
