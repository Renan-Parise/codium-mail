package services

import (
	"encoding/json"

	"github.com/Renan-Parise/codium-mail/entities"
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
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
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
		return err
	}

	body, err := json.Marshal(email)
	if err != nil {
		return err
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
