package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.intensivo
// go get -u github.com/stretchr/testify
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/rabbitmq/amqp091-go
// go mod tidy

// docker-compose down
// docker-compose up -d --build
// docker-compose up -d

// go run main.go

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 1_000; i++ { // 100 // 1_000 // 1_0000_000
		order := GenerateOrders()
		Publish(ch, order)
		time.Sleep(500 * time.Millisecond)
		// time.Sleep(1 * time.Second)
	}
}

func GenerateOrders() entity.Order {
	return entity.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax:   rand.Float64() * 10,
	}
}

func Publish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
