package service

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/repository"
)

type ImageService struct {
	Repo *repository.ImageRepository
}

func NewImageService(repo *repository.ImageRepository) *ImageService {
	return &ImageService{Repo: repo}
}

func (s *ImageService) UploadImage(listingID string, filePath string) (*models.Image, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		Folder: "kisumu-marketplace",
	})
	if err != nil {
		return nil, err
	}

	image := &models.Image{
		ListingID: listingID,
		URL:       uploadResult.SecureURL,
		SortOrder: 0,
	}

	if err := s.Repo.SaveImage(image); err != nil {
		return nil, err
	}

	return image, nil
}

func (s *ImageService) GetImagesByListing(listingID string) ([]models.Image, error) {
	return s.Repo.GetImagesByListing(listingID)
}
