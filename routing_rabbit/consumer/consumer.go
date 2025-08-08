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
		"topic_logs", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to declare exchange")

	q, err := ch.QueueDeclare(
		"",    // name ("" = let RabbitMQ assign a random name)
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare queue")

	// Bindings: pasa como argumentos los binding keys que quieres
	severities := os.Args[1:]
	if len(severities) == 0 {
		log.Fatalf("Usage: go run consumer.go [error] [warn]")
	}

	for _, s := range severities {
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // binding key
			"topic_logs", // exchange
			false,
			nil,
		)
		failOnError(err, "Failed to bind queue")
	}

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	failOnError(err, "Failed to register consumer")

	log.Printf("[*] Waiting for logs. Binding keys: %v", severities)
	for d := range msgs {
		log.Printf("Received [%s]: %s", d.RoutingKey, d.Body)
	}
}
