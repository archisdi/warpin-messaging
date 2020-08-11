package repositories

import (
	"encoding/json"
	db "warpin/libs/database"
	"warpin/models"
)

// MessageRepoInterface ...
type MessageRepoInterface interface {
	FindAll() []models.Message
	FindOne(string) models.Message
	Create(models.Message) error
}

// MessageRepo ...
type MessageRepo struct{}

// FindAll ...
func (*MessageRepo) FindAll() []models.Message {
	db := db.Database{"message"}
	byteValue := db.Get()

	var messages []models.Message
	json.Unmarshal(byteValue, &messages)

	return messages
}

// FindOne ...
func (*MessageRepo) FindOne(id string) models.Message {
	db := db.Database{"message"}
	byteValue := db.Get()

	var messages []models.Message
	json.Unmarshal(byteValue, &messages)

	var message models.Message
	for _, i := range messages {
		if i.ID == id {
			message = i
		}
	}

	return message
}

// Create ...
func (*MessageRepo) Create(message models.Message) error {
	db := db.Database{"message"}
	byteValue := db.Get()

	var messages []models.Message
	json.Unmarshal(byteValue, &messages)

	messages = append(messages, message)
	db.Set(messages)

	return nil
}
