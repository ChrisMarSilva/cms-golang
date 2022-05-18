package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.messageria.pulsar
// go get -u "github.com/apache/pulsar-client-go/pulsar"
// go mod tidy

// go run main.go

func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
		// URL: "pulsar://localhost:6650,localhost:6651,localhost:6652",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("#1: Could not instantiate Pulsar client: %v", err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{Topic: "my-topic"})
	if err != nil {
		log.Fatal("#2: client.CreateProducer Erro:", err)
	}

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-1")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-2")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-3")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-4")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-5")})
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{Payload: []byte("hello-6")})
	if err != nil {
		log.Println("#3: Failed to publish message", err)
	}
	defer producer.Close()

	log.Println("#4: Published message")

	// topic1 := "topic-1"
	// topic2 := "topic-2"
	// topics := []string{topic1, topic2}
	// consumer, err := client.Subscribe(pulsar.ConsumerOptions{ Topics:           topics, SubscriptionName: "multi-topic-sub", })

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{Topic: "my-topic", SubscriptionName: "my-sub", Type: pulsar.Shared})
	if err != nil {
		log.Fatal("#5: client.Subscribe Erro:", err)
	}
	defer consumer.Close()

	ctx, canc := context.WithTimeout(context.Background(), 5*time.Second)
	// ctx, canc := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer canc()

	msg, err := consumer.Receive(ctx)
	if err != nil {
		log.Fatal("#6: consumer.Receive Erro:", err)
	}
	fmt.Printf("#7: Received message\n")
	fmt.Printf("#7.1: Received message msgId: %#v'\n", msg.ID())
	fmt.Printf("#7.2: Received message content: '%s'\n", string(msg.Payload()))

	msg, err = consumer.Receive(ctx)
	if err != nil {
		log.Fatal("#8: consumer.Receive Erro:", err)
	}
	fmt.Printf("#9: Received message\n")
	fmt.Printf("#9.1: Received message msgId: %#v'\n", msg.ID())
	fmt.Printf("#9.2: Received message content: '%s'\n", string(msg.Payload()))

	// for i := 0; i < 10; i++ {
	// 	msg, err := consumer.Receive(context.Background())
	// 	if err != nil {
	// 		log.Fatal("#10: consumer.Receive Erro:", err)
	// 	}
	// 	log.Printf("#11: Received message msgId: %#v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
	// 	consumer.Ack(msg)
	// }
	// if err := consumer.Unsubscribe(); err != nil {
	// 	log.Fatal("#12: consumer.Unsubscribe Erro:", err)
	// }

	// channel := make(chan pulsar.ConsumerMessage, 100)
	// options := pulsar.ConsumerOptions{ Topic:            "topic-1", SubscriptionName: "my-subscription", Type:             pulsar.Shared, }
	// options.MessageChannel = channel
	// consumer, err := client.Subscribe(options)
	// for cm := range channel {
	//     msg := cm.Message
	//     fmt.Printf("Received message  msgId: %v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
	//     consumer.Ack(msg)
	// }

	reader, err := client.CreateReader(pulsar.ReaderOptions{Topic: "my-topic", StartMessageID: pulsar.EarliestMessageID()})
	if err != nil {
		log.Fatal("#13: client.CreateReader Erro:", err)
	}
	defer reader.Close()

	for reader.HasNext() {
		msg, err := reader.Next(context.Background())
		if err != nil {
			log.Fatal("#14: reader.Next Erro:", err)
		}
		fmt.Printf("#15: Received message msgId: %#v -- content: '%s'\n", msg.ID(), string(msg.Payload()))
	}

}
