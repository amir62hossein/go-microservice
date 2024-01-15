package rabbitMQ

import "github.com/rabbitmq/amqp091-go"

type RabbitClient struct {
	Channel *amqp091.Channel
}

func NewRabbitClient() *RabbitClient {

	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	return &RabbitClient{Channel: channel}
}
