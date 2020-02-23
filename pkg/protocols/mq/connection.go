package mq

import (
	"log"
	"os"

	"github.com/bungysheep/news-api/pkg/configs"
	"github.com/streadway/amqp"
)

var (
	// MqConnection - RabbitMQ connection
	MqConnection *amqp.Connection
)

// CreateMqConnection - Create rabbitmq connection
func CreateMqConnection() error {
	log.Printf("Creating rabbitmq connection...")

	resolveRabbitMQURL, err := resolveRabbitMQURL()
	if err != nil {
		return err
	}

	conn, err := amqp.Dial(resolveRabbitMQURL)
	if err != nil {
		return err
	}

	MqConnection = conn

	return nil
}

func resolveRabbitMQURL() (string, error) {
	resolveRabbitMQURL := os.Getenv("RABBITMQ_URL")
	if resolveRabbitMQURL != "" {
		return resolveRabbitMQURL, nil
	}

	return configs.RABBITMQURL, nil
}
