package main

import (
	"fmt"
	"order-service/helper"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	orderQueue, err := ch.QueueDeclare("ORDER", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	message, err := ch.Consume(orderQueue.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// productQueue, err := ch.QueueDeclare("PRODUCT", true, false, false, false, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// ctx := context.Background()
	forever := make(chan struct{})

	go func() {
		for d := range message {
			order, err := helper.SaveOrder(d.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(order)
		}
	}()

	<-forever

}
