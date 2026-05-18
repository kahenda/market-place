package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/service"
)

type MessageHandler struct {
	Service *service.MessageService
}

func NewMessageHandler(svc *service.MessageService) *MessageHandler {
	return &MessageHandler{Service: svc}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	senderID := c.GetString("user_id")

	var req models.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := h.Service.SendMessage(senderID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message sent successfully",
		"data":    msg,
	})
}

func (h *MessageHandler) GetConversation(c *gin.Context) {
	userID := c.GetString("user_id")
	otherUserID := c.Query("with")
	listingID := c.Query("listing_id")

	if otherUserID == "" || listingID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "with and listing_id query params required"})
		return
	}

	messages, err := h.Service.GetConversation(userID, otherUserID, listingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (h *MessageHandler) GetInbox(c *gin.Context) {
	userID := c.GetString("user_id")

	messages, err := h.Service.GetInbox(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
