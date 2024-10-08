package services

import (
	"encoding/json"

	"github.com/Renan-Parise/codium-mail/entities"
	"github.com/Renan-Parise/codium-mail/errors"
	"github.com/Renan-Parise/codium-mail/utils"
	"github.com/streadway/amqp"
)

type EmailService interface {
	PublishEmail(email entities.Email) error
}

type emailService struct{}

func NewEmailService() EmailService {
	return &emailService{}
}

func (s *emailService) PublishEmail(email entities.Email) error {
	conn, err := amqp.Dial(utils.GetRabbitMQURL())
	if err != nil {
		return errors.NewServiceError("Failed to connect to RabbitMQ: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return errors.NewServiceError("Failed to open a channel: " + err.Error())
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
		return errors.NewServiceError("Failed to declare a queue: " + err.Error())
	}

	body, err := json.Marshal(email)
	if err != nil {
		return errors.NewServiceError("Failed to marshal email: " + err.Error())
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}
