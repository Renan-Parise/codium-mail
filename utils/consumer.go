package utils

import (
	"encoding/json"
	"log"

	"github.com/Renan-Parise/codium-mail/agent"
	"github.com/Renan-Parise/codium-mail/entities"
	"github.com/streadway/amqp"
)

func StartConsumer(readyChan chan struct{}) {
	conn, err := amqp.Dial(GetRabbitMQURL())
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	q, err := ch.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %s", err)
	}

	log.Println("Consumer is ready, starting to process messages...")
	close(readyChan)

	for d := range msgs {
		log.Printf("Received a message: %s", string(d.Body))

		var email entities.Email
		if err := json.Unmarshal(d.Body, &email); err != nil {
			log.Printf("Error decoding JSON: %s", err)
			d.Nack(false, false)
			continue
		}

		err := agent.SendEmail(email)
		if err != nil {
			log.Printf("Failed to send email to %s: %s", email.Address, err)
			d.Nack(false, true)
		} else {
			log.Printf("Email sent to %s", email.Address)
			d.Ack(false)
		}
	}

	log.Println("Message channel closed, shutting down consumer.")
}
