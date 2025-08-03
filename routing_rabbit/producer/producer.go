package main

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")

	err = ch.ExchangeDeclare(
		"direct_logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	body := "Hello log"
	severity := os.Args[1] // ej: "error", "warn", etc.

	err = ch.Publish(
		"direct_logs",
		severity, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish")
	log.Printf("Sent [%s]: %s", severity, body)
}
