package repositories

import (
	"encoding/json"
	db "warpin/libs/database"
	"warpin/models"
)

// MessageRepo ...
type MessageRepo struct{}

// FindAll ...
func (repo *MessageRepo) FindAll() []models.Message {
	db := db.Database{"message"}
	byteValue := db.Get()

	var messages []models.Message
	json.Unmarshal(byteValue, &messages)

	return messages
}

// FindOne ...
func (repo *MessageRepo) FindOne(id string) models.Message {
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
func (repo *MessageRepo) Create(message models.Message) models.Message {
	db := db.Database{"message"}
	byteValue := db.Get()

	var messages []models.Message
	json.Unmarshal(byteValue, &messages)

	messages = append(messages, message)
	db.Set(messages)

	return message
}
