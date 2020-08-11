package controllers

import (
	"testing"
	"time"

	m "warpin/models"

	mocks "warpin/mocks"

	mock "github.com/stretchr/testify/mock"

	assert "github.com/stretchr/testify/assert"
)

func TestMessageFindAll(t *testing.T) {
	messageRepo := &mocks.MessageRepoInterface{}
	mqtt := &mocks.MqttInterface{}

	messagesMock := []m.Message{
		{"id1", "Hola amigos", time.Now(), time.Now()},
		{"id2", "Hola Hermano", time.Now(), time.Now()},
	}

	messageRepo.On("FindAll").Return(messagesMock).Once()

	c := MessageController{
		MessageRepo: messageRepo,
		Mqtt:        mqtt,
	}

	//
	messages := c.GetMessages()
	assert.Equal(t, 2, len(messages))
}

func TestMessageFindOne(t *testing.T) {
	messageRepo := &mocks.MessageRepoInterface{}
	mqtt := &mocks.MqttInterface{}

	messageMock := m.Message{
		"6969", "Hola amigos", time.Now(), time.Now(),
	}
	messageRepo.On("FindOne", "6969").Return(messageMock, nil).Once()
	messageRepo.On("FindOne", "420").Return(m.Message{}).Once()

	c := MessageController{
		MessageRepo: messageRepo,
		Mqtt:        mqtt,
	}

	message, _ := c.GetMessagesBy("6969")
	assert.Equal(t, "Hola amigos", message.Text)

	_, err := c.GetMessagesBy("420")
	assert.Equal(t, err.Error(), "message not found")
}

func TestMessageCreate(t *testing.T) {
	messageRepo := &mocks.MessageRepoInterface{}
	mqtt := &mocks.MqttInterface{}

	c := MessageController{
		MessageRepo: messageRepo,
		Mqtt:        mqtt,
	}

	text := "Why do birds suddenly appears"
	messageRepo.On("Create", mock.Anything).Return(nil).Once()
	mqtt.On("Publish", "warpin-messages", mock.Anything).Return(nil).Once()

	newMessage, _ := c.PostMessages(struct{ Text string }{Text: text})
	assert.Equal(t, newMessage.Text, text)

	_, error := c.PostMessages(struct{ Text string }{Text: "h"})
	assert.Equal(t, error.Error(), "message must be over 1 character")
}

func TestMessageListen(t *testing.T) {
	messageRepo := &mocks.MessageRepoInterface{}
	mqtt := &mocks.MqttInterface{}

	c := MessageController{
		MessageRepo: messageRepo,
		Mqtt:        mqtt,
	}

	messageMock := m.Message{
		"6969", "Hola amigos", time.Now(), time.Now(),
	}
	mqtt.On("Listen", "warpin-messages", mock.AnythingOfType("*[]string")).Return().Run(func(args mock.Arguments) {
		arg := args.Get(1).(*[]string)
		*arg = append(*arg, messageMock.Serialize())
	}).Once()

	message := c.GetMessagesListen(struct{ Timeout int }{Timeout: 1})
	assert.Equal(t, messageMock.Text, message[0].Text)
}
