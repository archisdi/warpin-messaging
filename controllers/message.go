package controllers

import (
	"time"

	m "warpin/models"

	"github.com/google/uuid"
)

// MessageController ...
type MessageController struct {
}

// GetMessages returns all messages in database
func (c *MessageController) GetMessages() []m.Message {
	return []m.Message{
		m.Message{
			ID:        uuid.New(),
			Text:      "Warung Pintar",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// GetMessagesBy returns specific messages in database by id
func (c *MessageController) GetMessagesBy(id string) m.Message {
	return m.Message{
		ID:        uuid.New(),
		Text:      "Warung Pintar",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// PostMessages creates a new message
func (c *MessageController) PostMessages(payload struct{ Text string }) m.Message {
	println("Received Message: " + payload.Text)
	return m.NewMessage(payload.Text)
}
