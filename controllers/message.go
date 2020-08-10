package controllers

import (
	"errors"
	mqtt "warpin/libs/mqtt"
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
func (c *MessageController) PostMessages(payload struct{ Text string }) (m.Message, error) {
	message := m.NewMessage(payload.Text)
	if len(message.Text) <= 1 {
		return message, errors.New("message must be over 1 character")
	}

	// persist new message
	messageRepo := repositories.MessageRepo{}
	messageRepo.Create(message)

	// publish to mqtt server
	mqtt.Publish("warpin-messaging", message.String())

	return message, nil
}
