package repository

import (
	"database/sql"
	"github.com/kahenda/marketplace/internal/models"
)

type MessageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

func (r *MessageRepository) SendMessage(msg *models.Message) error {
	query := `
		INSERT INTO messages (sender_id, receiver_id, listing_id, body)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`
	return r.DB.QueryRow(query,
		msg.SenderID,
		msg.ReceiverID,
		msg.ListingID,
		msg.Body,
	).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *MessageRepository) GetConversation(userID1, userID2, listingID string) ([]models.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, listing_id, body, created_at
		FROM messages
		WHERE listing_id = $3
		AND (
			(sender_id = $1 AND receiver_id = $2) OR
			(sender_id = $2 AND receiver_id = $1)
		)
		ORDER BY created_at ASC`
	rows, err := r.DB.Query(query, userID1, userID2, listingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.ListingID, &m.Body, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func (r *MessageRepository) GetInbox(userID string) ([]models.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, listing_id, body, created_at
		FROM messages
		WHERE receiver_id = $1
		ORDER BY created_at DESC`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.ListingID, &m.Body, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}
