package models

import (
	"time"

	"github.com/google/uuid"
)

// Message ...
type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewMessage returns new message
func NewMessage(text string) Message {
	return Message{
		ID:        uuid.New(),
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
