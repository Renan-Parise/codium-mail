package services

import (
	"encoding/json"
	"time"

	"github.com/Renan-Parise/mail/entities"
	"github.com/Renan-Parise/mail/errors"
	"github.com/Renan-Parise/mail/utils"
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

	channel, err := conn.Channel()
	if err != nil {
		return errors.NewServiceError("Failed to open channel: " + err.Error())
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		utils.GetEmailQueueName(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.NewServiceError("Failed to declare queue: " + err.Error())
	}

	err = channel.Confirm(false)
	if err != nil {
		return errors.NewServiceError("Failed to set channel to confirm mode: " + err.Error())
	}

	body, err := json.Marshal(email)
	if err != nil {
		return errors.NewServiceError("Failed to marshal email: " + err.Error())
	}

	confirmChan := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	err = channel.Publish(
		"",
		utils.GetEmailQueueName(),
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		return errors.NewServiceError("Failed to publish message: " + err.Error())
	}

	select {
	case confirm := <-confirmChan:
		if confirm.Ack {
			return nil
		} else {
			return errors.NewServiceError("Message not acknowledged by broker")
		}
	case <-time.After(5 * time.Second):
		return errors.NewServiceError("Timeout waiting for publisher confirmation")
	}
}
