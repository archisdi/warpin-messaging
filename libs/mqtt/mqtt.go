package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client ...
var Client *mqtt.Client

// Mqtt ...
type Mqtt struct{}

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

// Listen ...
func (m *Mqtt) Listen(topic string, response *[]string) {
	client := *Client
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		*response = append(*response, string(msg.Payload()))
	})
}

// Publish ...
func (m *Mqtt) Publish(topic string, message string) {
	client := *Client
	client.Publish(topic, 0, false, message)
}

// Initialize ..
func Initialize(ClientID string, uri *url.URL) {
	if Client == nil {
		client := connect(ClientID, uri)
		Client = &client
	}
}
