package controllers

import (
	"fmt"
	"testing"
	"warpin/libs/mqtt"
	"warpin/repositories"
)

func TestMessageFindAll(t *testing.T) {
	c := MessageController{
		MessageRepo: repositories.MessageRepo{},
		Mqtt:        mqtt.Mqtt{},
	}

	messages := c.GetMessages()
	for _, message := range messages {
		fmt.Println(message.Serialize())
	}

	t.Log("success")
}
