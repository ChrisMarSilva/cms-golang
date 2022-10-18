package kafka

import (
	"log"
	"os"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KafkaBootstrapServers"),
		//"session.timeout.ms": 6000,
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return p
}

// Publish is simple function created to publish new message to kafka
func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
		//Headers:      []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}

	err := producer.Produce(message, nil)
	if err != nil {
		// if err.(kafka.Error).Code() == kafka.ErrQueueFull {
		// 	// Producer queue is full, wait 1s for messages
		// 	// to be delivered then try again.
		// 	time.Sleep(time.Second)
		// 	continue
		// }
		return err
	}

	// // Flush and close the producer and the events channel
	// for p.Flush(10000) > 0 {
	// 	fmt.Print("Still waiting to flush outstanding messages\n", err)
	// }
	// p.Close()

	return nil
}
