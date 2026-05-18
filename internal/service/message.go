package service

import (
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/repository"
)

type MessageService struct {
	Repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{Repo: repo}
}

func (s *MessageService) SendMessage(senderID string, req *models.SendMessageRequest) (*models.Message, error) {
	msg := &models.Message{
		SenderID:   senderID,
		ReceiverID: req.ReceiverID,
		ListingID:  req.ListingID,
		Body:       req.Body,
	}
	if err := s.Repo.SendMessage(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (s *MessageService) GetConversation(userID1, userID2, listingID string) ([]models.Message, error) {
	return s.Repo.GetConversation(userID1, userID2, listingID)
}

func (s *MessageService) GetInbox(userID string) ([]models.Message, error) {
	return s.Repo.GetInbox(userID)
}
