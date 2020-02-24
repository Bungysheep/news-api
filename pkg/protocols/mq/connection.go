package mq

import (
	"log"
	"net"
	"os"
	"time"

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

	for i := 0; i < configs.NUMBERDIALATTEMPT; i++ {
		MqConnection, err = amqp.Dial(resolveRabbitMQURL)
		if err != nil {
			opErr, ok := err.(*net.OpError)
			if !ok && opErr.Op != "dial" {
				return err
			}
		}
		time.Sleep(5 * time.Second)
	}

	return err
}

func resolveRabbitMQURL() (string, error) {
	resolveRabbitMQURL := os.Getenv("RABBITMQ_URL")
	if resolveRabbitMQURL != "" {
		return resolveRabbitMQURL, nil
	}

	return configs.RABBITMQURL, nil
}
