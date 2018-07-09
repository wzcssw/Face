package lib

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type callback func()

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Send(data []byte) {
	// Connect to MQ
	conn, err := amqp.Dial("amqp://wzc:123456@192.168.1.10:5672/")
	// conn, err := amqp.Dial("amqp://zhouwei:zhouwei@192.168.1.73:5672/")
	// conn, err := amqp.Dial("amqp://wzc:123456@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Next we create a channel,
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// go!
	q, err := ch.QueueDeclare(
		"face", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	f := func() {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			})
		fmt.Printf("已经发送: %d\n", len(data))
		failOnError(err, "Failed to publish a message")
	}
	f()
}
