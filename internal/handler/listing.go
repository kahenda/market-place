package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/service"
)

type ListingHandler struct {
	Service *service.ListingService
}

func NewListingHandler(svc *service.ListingService) *ListingHandler {
	return &ListingHandler{Service: svc}
}

func (h *ListingHandler) CreateListing(c *gin.Context) {
	userID := c.GetString("user_id")

	var req models.CreateListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listing, err := h.Service.CreateListing(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Listing created successfully",
		"listing": listing,
	})
}

func (h *ListingHandler) GetListings(c *gin.Context) {
	category := c.Query("category")

	var listings []models.Listing
	var err error

	if category != "" {
		listings, err = h.Service.GetListingsByCategory(category)
	} else {
		listings, err = h.Service.GetAllListings()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
	})
}
