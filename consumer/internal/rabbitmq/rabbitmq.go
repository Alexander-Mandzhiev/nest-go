package rabbitmq

import (
	"consumer/internal/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Body      string
	QueueName string
}

func ConnectMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errors.FailOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	ch, err := conn.Channel()
	errors.FailOnError(err, "Failed to open a channel")
	// defer ch.Close()
	return conn, ch
}

func CloseMQ(conn *amqp.Connection, channel *amqp.Channel) {
	defer conn.Close()    //rabbit mq close
	defer channel.Close() //rabbit mq channel close
}
