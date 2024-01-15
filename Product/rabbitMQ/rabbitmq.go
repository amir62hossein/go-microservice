package rabbitMQ

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	return &RabbitMQ{Channel: ch}
}
