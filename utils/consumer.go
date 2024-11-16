package utils

import (
	"encoding/json"

	"github.com/Renan-Parise/mail/agent"
	"github.com/Renan-Parise/mail/entities"
	"github.com/Renan-Parise/mail/errors"
	"github.com/streadway/amqp"
)

func StartConsumer(readyChan chan struct{}) {
	conn, err := amqp.Dial(GetRabbitMQURL())
	if err != nil {
		errors.NewConsumerError("Failed to connect to RabbitMQ: " + err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		errors.NewConsumerError("Failed to open a channel: " + err.Error())
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
		errors.NewConsumerError("Failed to declare a queue: " + err.Error())
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
		errors.NewConsumerError("Failed to register a consumer: " + err.Error())
	}

	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		errors.NewConsumerError("Failed to set QoS: " + err.Error())
	}

	close(readyChan)

	for d := range msgs {
		var email entities.Email
		if err := json.Unmarshal(d.Body, &email); err != nil {
			errors.NewConsumerError("Failed to unmarshal email: " + err.Error())
			d.Nack(false, false)
			continue
		}

		err := agent.SendEmail(email)
		if err != nil {
			errors.NewConsumerError("Failed to send email: " + err.Error())
			d.Nack(false, true)
		} else {
			d.Ack(false)
		}
	}
}
