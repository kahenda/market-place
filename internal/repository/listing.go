package repository

import (
	"database/sql"
	"fmt"
	"strings"

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
		INSERT INTO listings (user_id, title, description, price, category, gender, size, condition, area, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'active')
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		listing.UserID,
		listing.Title,
		listing.Description,
		listing.Price,
		listing.Category,
		listing.Gender,
		listing.Size,
		listing.Condition,
		listing.Area,
	).Scan(&listing.ID, &listing.CreatedAt)
}

func (r *ListingRepository) GetListings(category, gender, size, condition, area string) ([]models.Listing, error) {
	base := `SELECT id, user_id, title, description, price, category, gender, size, condition, area, status, created_at FROM listings WHERE status = 'active'`

	args := []interface{}{}
	conditions := []string{}
	i := 1

	if category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", i))
		args = append(args, category)
		i++
	}
	if gender != "" {
		conditions = append(conditions, fmt.Sprintf("gender = $%d", i))
		args = append(args, gender)
		i++
	}
	if size != "" {
		conditions = append(conditions, fmt.Sprintf("size = $%d", i))
		args = append(args, size)
		i++
	}
	if condition != "" {
		conditions = append(conditions, fmt.Sprintf("condition = $%d", i))
		args = append(args, condition)
		i++
	}
	if area != "" {
		conditions = append(conditions, fmt.Sprintf("area = $%d", i))
		args = append(args, area)
		i++
	}

	if len(conditions) > 0 {
		base += " AND " + strings.Join(conditions, " AND ")
	}
	base += " ORDER BY created_at DESC"

	rows, err := r.DB.Query(base, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listings []models.Listing
	for rows.Next() {
		var l models.Listing
		var gender, size, condition, area sql.NullString
		if err := rows.Scan(
			&l.ID, &l.UserID, &l.Title, &l.Description,
			&l.Price, &l.Category, &gender, &size,
			&condition, &area, &l.Status, &l.CreatedAt,
		); err != nil {
			return nil, err
		}
		l.Gender = gender.String
		l.Size = size.String
		l.Condition = condition.String
		l.Area = area.String
		listings = append(listings, l)
	}
	return listings, nil
}
