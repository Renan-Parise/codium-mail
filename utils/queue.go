package utils

import (
	"github.com/Renan-Parise/codium-mail/errors"
	"github.com/streadway/amqp"
)

func EnsureQueueExists() error {
	conn, err := amqp.Dial(GetRabbitMQURL())
	if err != nil {
		return errors.NewServiceError("Failed to connect to RabbitMQ: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return errors.NewServiceError("Failed to open channel: " + err.Error())
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}
