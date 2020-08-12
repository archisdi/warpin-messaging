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

// MqttInterface ...
type MqttInterface interface {
	Listen(string, *[]string)
	Publish(string, string)
}

// Mqtt ...
type Mqtt struct{}

func connect(clientID string, uri *url.URL) (mqtt.Client, error) {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
		return client, err
	}
	return client, nil
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
func (*Mqtt) Listen(topic string, response *[]string) {
	client := *Client
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		*response = append(*response, string(msg.Payload()))
	})
}

// Publish ...
func (*Mqtt) Publish(topic string, message string) {
	client := *Client
	client.Publish(topic, 0, false, message)
}

// Initialize ..
func Initialize(ClientID string, uri string) error {
	if Client == nil {
		URI, _ := url.Parse(uri)
		client, err := connect(ClientID, URI)
		if err != nil {
			return err
		}
		Client = &client
	}
	return nil
}
