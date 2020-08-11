package controllers

import (
	"encoding/json"
	"errors"
	"time"
	mqtt "warpin/libs/mqtt"
	repo "warpin/repositories"

	m "warpin/models"
)

// MessageController ...
type MessageController struct {
	MessageRepo repo.MessageRepoInterface
	Mqtt        mqtt.MqttInterface
}

// GetMessages returns all messages in database
func (c *MessageController) GetMessages() []m.Message {
	messages := c.MessageRepo.FindAll()
	return messages
}

// GetMessagesListen listen for realtime message for given time
func (c *MessageController) GetMessagesListen(query struct{ Timeout int }) []m.Message {
	response := []string{}
	timeOut := 10

	if query.Timeout != 0 {
		timeOut = query.Timeout
	}

	c.Mqtt.Listen("warpin-messages", &response)
	time.Sleep(time.Duration(timeOut) * time.Second)

	messages := []m.Message{}
	for _, raw := range response {
		var message m.Message
		json.Unmarshal([]byte(raw), &message)
		messages = append(messages, message)
	}

	return messages
}

// GetMessagesBy returns specific messages in database by id
func (c *MessageController) GetMessagesBy(id string) (m.Message, error) {
	message := c.MessageRepo.FindOne(id)

	if message.ID == "" {
		return message, errors.New("message not found")
	}

	return message, nil
}

// PostMessages creates a new message
func (c *MessageController) PostMessages(payload struct{ Text string }) (m.Message, error) {
	var message m.Message
	if len(payload.Text) <= 1 {
		return message, errors.New("message must be over 1 character")
	}

	message = m.NewMessage(payload.Text)

	c.MessageRepo.Create(message)
	c.Mqtt.Publish("warpin-messages", message.Serialize())

	return message, nil
}
