package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/service"
)

type ReviewHandler struct {
	Service *service.ReviewService
}

func NewReviewHandler(svc *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: svc}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	reviewerID := c.GetString("user_id")

	var req models.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.Service.CreateReview(reviewerID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review submitted successfully",
		"review":  review,
	})
}

func (h *ReviewHandler) GetReviewsBySeller(c *gin.Context) {
	sellerID := c.Param("seller_id")

	reviews, err := h.Service.GetReviewsBySeller(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
