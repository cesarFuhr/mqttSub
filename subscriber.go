package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Subscriber struct {
	c MQTT.Client
}

func defaultHandler(c MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC -> %s\n", msg.Topic())
	fmt.Printf("MSG -> %s\n", msg.Payload())
}

func (s *Subscriber) Connect(brokerUrl string) {
	opts := MQTT.NewClientOptions().AddBroker(brokerUrl)
	id, _ := uuid.NewUUID()
	clientName := fmt.Sprintf("Sub-%s", id)
	opts.SetClientID(clientName)
	opts.SetDefaultPublishHandler(defaultHandler)

	s.c = MQTT.NewClient(opts)
	if token := s.c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (s *Subscriber) Subscribe(topic string) error {
	if token := s.c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return token.Error()
	}
	return nil
}

func (s *Subscriber) Unsubscribe(topic string) error {
   if token := s.c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
	   fmt.Println(token.Error())
	   return token.Error()
   }
	return nil
}
