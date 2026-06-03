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
	c.JSON(http.StatusCreated, gin.H{"message": "Listing created successfully", "listing": listing})
}

func (h *ListingHandler) GetListingByID(c *gin.Context) {
	id := c.Param("id")
	listing, err := h.Service.GetListingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if listing == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"listing": listing})
}

func (h *ListingHandler) GetListings(c *gin.Context) {
	category := c.Query("category")
	gender := c.Query("gender")
	size := c.Query("size")
	condition := c.Query("condition")
	area := c.Query("area")

	listings, err := h.Service.GetListings(category, gender, size, condition, area)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"listings": listings, "count": len(listings)})
}

func (h *ListingHandler) GetOptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"categories": models.Categories,
		"genders":    models.Genders,
		"sizes":      models.Sizes,
		"conditions": models.Conditions,
		"areas":      models.KisumuAreas,
	})
}
