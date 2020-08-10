package controllers

import (
	"errors"
	"warpin/repositories"

	m "warpin/models"
)

// MessageController ...
type MessageController struct {
}

// GetMessages returns all messages in database
func (c *MessageController) GetMessages() []m.Message {
	messageRepo := repositories.MessageRepo{}
	messages := messageRepo.FindAll()

	return messages
}

// GetMessagesBy returns specific messages in database by id
func (c *MessageController) GetMessagesBy(id string) (m.Message, error) {
	messageRepo := repositories.MessageRepo{}
	message := messageRepo.FindOne(id)

	if message.ID == "" {
		return message, errors.New("message not found")
	}

	return message, nil
}

// PostMessages creates a new message
func (c *MessageController) PostMessages(payload struct{ Text string }) m.Message {
	messageRepo := repositories.MessageRepo{}
	message := m.NewMessage(payload.Text)

	messageRepo.Create(message)

	return message
}
