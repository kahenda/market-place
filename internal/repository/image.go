package repository

import (
	"database/sql"
	"github.com/kahenda/marketplace/internal/models"
)

type ImageRepository struct {
	DB *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

func (r *ImageRepository) SaveImage(image *models.Image) error {
	query := `
		INSERT INTO images (listing_id, image_url, sort_order)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		image.ListingID,
		image.URL,
		image.SortOrder,
	).Scan(&image.ID, &image.CreatedAt)
}

func (r *ImageRepository) GetImagesByListing(listingID string) ([]models.Image, error) {
	query := `
		SELECT id, listing_id, image_url, sort_order, created_at
		FROM images WHERE listing_id = $1
		ORDER BY sort_order ASC`
	rows, err := r.DB.Query(query, listingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var img models.Image
		if err := rows.Scan(&img.ID, &img.ListingID, &img.URL, &img.SortOrder, &img.CreatedAt); err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}
