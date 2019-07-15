package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var (
	amqpURI  = flag.String("amqp", "amqp://docker-server.cloudapp.net:5672", "AQMP_URI")
	exchange = "exchange"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Declaring the connections as Globals

func initAmqp() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		"test-exchange", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare the Exchange")
	return conn, ch
}

func AMQPInit() {
	flag.Parse()
	initAmqp()
}

func PublishToQueue(uuidpayload string) {
	_, ch := initAmqp()
	AMQPInit()

	err := ch.Publish(
		exchange,
		"",
		false,
		true,
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "test/plain",
			Body:         []byte(uuidpayload),
			Timestamp:    time.Now(),
		})
	failOnError(err, "Failed to Publish on RabbitMQ")
}
