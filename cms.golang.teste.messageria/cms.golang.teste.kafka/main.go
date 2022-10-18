package main

// cd 'C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.messageria\cms.golang.teste.kafka'
// https://github.com/codeedu/imersao10/tree/main/simulator

// go mod init github.com/chrismarsilva/cms.golang.teste.kafka
// go get -u github.com/confluentinc/confluent-kafka-go/kafka
// go get -u github.com/joho/godotenv
// go mod tidy

// go run main.go

import (
	"fmt"
	"log"
	_ "time"
	kafka2 "github.com/chrismarsilva/cms.golang.teste.kafka/application/kafka"
	"github.com/chrismarsilva/cms.golang.teste.kafka/infra/kafka"
    "github.com/joho/godotenv"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "route.new-position", producer)
	// for {
	// 	_ = 1
	// }
}
