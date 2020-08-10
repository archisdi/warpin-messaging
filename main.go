package main

import (
	"fmt"
	"time"
	"warpin/models"
)

func main() {
	message := models.Message{
		ID:        "1169",
		Text:      "Hola Amigos",
		CreatedAt: time.Now(),
	}

	fmt.Println(message)
}
