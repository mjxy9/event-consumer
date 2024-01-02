package main

import (
	"encoding/json"
	"os"

	"github.com/emejotaw/event-consumer/pkg/events/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	queueName := "event-queue"
	rabbitMQ, err := rabbitmq.NewRabbitMQ(queueName)

	if err != nil {
		panic(err)
	}

	deliverych := make(chan amqp.Delivery)
	go rabbitMQ.Consume(deliverych)

	encoder := json.NewEncoder(os.Stdout)

	for delivery := range deliverych {
		event := string(delivery.Body)
		delivery.Ack(false)
		encoder.Encode(event)
	}
}
