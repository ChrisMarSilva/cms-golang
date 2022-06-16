package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

// go get github.com/wagslane/go-rabbitmq
// https://github.com/wagslane/go-rabbitmq

var consumerName = "example"

func main() {

	producer()
	consumer()

	fmt.Print("FIM")
}

func producer() {

	publisher, err := rabbitmq.NewPublisher("amqp://user:pass@localhost",rabbitmq.Config{}, rabbitmq.WithPublisherOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := publisher.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// err = publisher.Publish([]byte("hello, world"), []string{"routing_key"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// err = publisher.Publish(
	// 	[]byte("hello, world"),
	// 	[]string{"routing_key"},
	// 	rabbitmq.WithPublishOptionsContentType("application/json"),
	// 	rabbitmq.WithPublishOptionsMandatory,
	// 	rabbitmq.WithPublishOptionsPersistentDelivery,
	// 	rabbitmq.WithPublishOptionsExchange("events"),
	// )

	returns := publisher.NotifyReturn()
	go func() {
		for r := range returns {
			log.Printf("message returned from server: %s", string(r.Body))
		}
	}()

	confirmations := publisher.NotifyPublish()
	go func() {
		for c := range confirmations {
			log.Printf("message confirmed from server. tag: %v, ack: %v", c.DeliveryTag, c.Ack)
		}
	}()
	

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err = publisher.Publish(
				[]byte("hello, world"),
				[]string{"routing_key"},
				rabbitmq.WithPublishOptionsContentType("application/json"),
				rabbitmq.WithPublishOptionsMandatory,
				rabbitmq.WithPublishOptionsPersistentDelivery,
				rabbitmq.WithPublishOptionsExchange("amq.topic"),
			)
			if err != nil {
				log.Println(err)
			}
		case <-done:
			fmt.Println("stopping publisher")
			return
		}
	}

}

func consumer() {

	consumer, err := rabbitmq.NewConsumer("amqp://guest:guest@localhost:5672/", rabbitmq.Config{}, rabbitmq.WithConsumerOptionsLogging)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	// defer consumer.Close()

	defer func() {
		err := consumer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		"my_queue",
		[]string{"routing_key", "routing_key_2"},
		rabbitmq.WithConsumeOptionsConcurrency(10),
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsQuorum,
		rabbitmq.WithConsumeOptionsBindingExchangeName("events"),
		rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsConsumerName(consumerName),
	)

	if err != nil {
		log.Fatal(err)
	}

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")

}

