package utils

import (
	"encoding/json"
	"os"

	"github.com/Renan-Parise/codium-mail/entities"
	"github.com/streadway/amqp"
)

func GetRabbitMQURL() string {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		panic("RABBITMQ_URL is not set")
	}

	return url
}

func StartConsumer() {
	conn, err := amqp.Dial(GetRabbitMQURL())
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

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
		true,  // Auto-acknowledge
		false, // Non-exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			var email entities.Email
			if err := json.Unmarshal(d.Body, &email); err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			log.Printf("Received an email request: %+v", email)

			err := SendEmail(email)
			if err != nil {
				log.Printf("Failed to send email to %s: %s", email.Address, err)
			} else {
				log.Printf("Email sent to %s", email.Address)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}

func SendEmail(email entities.Email) error {
	// Implement your email sending logic here.
	// For example, using SMTP or an external service.
	// This is a placeholder implementation.
	log.Printf("Sending email to %s with subject '%s'", email.Address, email.Subject)
	return nil
}
