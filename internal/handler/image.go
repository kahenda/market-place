package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/marketplace/internal/service"
)

type ImageHandler struct {
	Service *service.ImageService
}

func NewImageHandler(svc *service.ImageService) *ImageHandler {
	return &ImageHandler{Service: svc}
}

func (h *ImageHandler) UploadImage(c *gin.Context) {
	listingID := c.Param("listing_id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file required"})
		return
	}

	tmpPath := filepath.Join("/tmp", file.Filename)
	if err := c.SaveUploadedFile(file, tmpPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	image, err := h.Service.UploadImage(listingID, tmpPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"image":   image,
	})
}

func (h *ImageHandler) GetImages(c *gin.Context) {
	listingID := c.Param("listing_id")

	images, err := h.Service.GetImagesByListing(listingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": images})
}
