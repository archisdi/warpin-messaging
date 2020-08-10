package models

import "time"

// Message ...
type Message struct {
	ID        string
	Text      string
	CreatedAt time.Time
}
