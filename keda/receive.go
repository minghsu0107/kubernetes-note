package main

import (
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

/*
All deliveries in AMQP must be acknowledged. It is expected of the consumer to
call Delivery.Ack after it has successfully processed the delivery. If the
consumer is cancelled or the channel or connection is closed any unacknowledged
deliveries will be requeued at the end of the same queue.

Acknowledging a message tells RabbitMQ that it has been taken care of and RabbitMQ can delete it now
*/

func main() {
	url := os.Args[1]
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// A Channel is the application session that is opened for each piece of your app
	// to communicate with the RabbitMQ broker
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to set QoS")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable, if true then the queue definition will survive a server restart, not the messages in it
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// prefetching so that producer won't overwhelm consumer
	/*
		With a prefetch count greater than zero, the server will deliver that many
		messages to consumers before acknowledgments are received.  The server ignores
		this option when consumers are started with noAck because no acknowledgments
		are expected or sent.

		With a prefetch size greater than zero, the server will try to keep at least
		that many bytes of deliveries flushed to the network before receiving
		acknowledgments from the consumers. This option is ignored when consumers are
		started with noAck.

		When global is true, these Qos settings apply to all existing and future
		consumers on all channels on the same connection.  When false, the Channel.Qos
		settings will apply to all existing and future consumers on this channel.
	*/
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer, empty string will cause the library to generate a unique identity
		false,  // auto-ack, when autoAck (also known as noAck) is true, the server will acknowledge deliveries to this consumer prior to writing the delivery to the network
		// When autoAck is true, the consumer should not call Delivery.Ack. Automatically acknowledging deliveries means that some deliveries may get lost

		false, // exclusive, when exclusive is true, the server will ensure that this is the sole consumer from this queue
		false, // no-local, not supported by RabbitMQ
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	/*
		Inflight messages, limited by Channel.Qos will be buffered until received from
		the returned chan.
		When the Channel or Connection is closed, all buffered and inflight messages will
		be dropped.
		When the consumer tag is cancelled, all inflight messages will be delivered until
		the returned chan is closed.
	*/
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			time.Sleep(1 * time.Second)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
