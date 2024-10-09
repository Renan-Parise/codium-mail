package utils

import (
	"github.com/streadway/amqp"
)

func EnsureQueueExists() error {
	conn, err := amqp.Dial(GetRabbitMQURL())
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
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
