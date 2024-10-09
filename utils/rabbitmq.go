package utils

import (
	"os"
)

func GetRabbitMQURL() string {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		panic("RABBITMQ_URL is not set")
	}
	return url
}

func GetEmailQueueName() string {
	return "email_queue"
}
