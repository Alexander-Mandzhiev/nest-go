package main

import (
	"consumer/internal/errors"
	"consumer/internal/rabbitmq"
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, ch := rabbitmq.ConnectMQ()
	defer rabbitmq.CloseMQ(conn, ch)

	q, err := ch.QueueDeclare("values", true, false, false, false, nil)
	errors.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(1, 0, false)
	errors.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(q.Name, "values", false, false, false, false, nil)
	errors.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			log.Printf("1.Received a message: %s", d.Body)
			res := d.Body

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          []byte(res),
				})
			errors.FailOnError(err, "Failed to publish a message")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
