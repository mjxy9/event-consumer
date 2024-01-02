package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQ(queueName string) (*RabbitMQ, error) {
	rabbitMq := &RabbitMQ{queueName: queueName}

	err := rabbitMq.connect()

	if err != nil {
		return nil, err
	}

	return rabbitMq, nil
}

func (r *RabbitMQ) connect() error {

	dsn := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(dsn)

	if err != nil {
		log.Printf("could not establish amqp connection, error: %v", err)
		return err
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Printf("could not connect with channel, error: %v", err)
		return err
	}

	r.channel = channel
	return nil
}

func (r *RabbitMQ) Consume(eventch chan amqp.Delivery) error {

	events, err := r.channel.Consume(r.queueName, "event-consumer", false, false, false, false, nil)

	if err != nil {
		log.Printf("could not consume messages, error: %v", err)
		return err
	}

	for event := range events {
		eventch <- event
	}

	return nil
}
