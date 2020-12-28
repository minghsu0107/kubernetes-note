package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	url := os.Args[1]
	messageCount, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to parse second arg as messageCount : int")
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/*
		When you want a single message to be delivered to a single queue, you can
		publish to the default exchange with the routingKey of the queue name.  This is
		because every declared queue gets an implicit route to the default exchange.

		Publishings can be undeliverable when the mandatory flag is true and no queue is
		bound that matches the routing key, or when the immediate flag is true and no
		consumer on the matched queue is ready to accept the delivery.
	*/

	for i := 0; i < messageCount; i++ {
		body := fmt.Sprintf("Hello World: %d", i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key

			// If mandatory is set and after running the bindings the message was placed on zero queues
			// then the message is returned to the sender (with a basic.return).
			// If mandatory had not been set under the same circumstances the server would silently drop the message
			false, // mandatory

			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}
}
