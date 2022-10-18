package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"
	"github.com/chrismarsilva/cms.golang.teste.kafka/application/route"
	"github.com/chrismarsilva/cms.golang.teste.kafka/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// Produce is responsible to publish the positions of each request
func Produce(msg *ckafka.Message) {
	router := route.NewRoute()

	// Example of a json request:
	// 		{"clientId":"1","routeId":"1"}
	//		{"clientId":"2","routeId":"2"}
	//		{"clientId":"3","routeId":"3"}

	err := json.Unmarshal(msg.Value, &router)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = router.LoadPositions()
	if err != nil {
		log.Println(err.Error())
		return
	}

	positions, err := router.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
		return
	}

	producer := kafka.NewKafkaProducer()
	topic := os.Getenv("KafkaProduceTopic")

	for _,  position := range positions {
		kafka.Publish(position, topic, producer)
		time.Sleep(time.Millisecond * 1) // 500 somente para nao ser muito rapido
	}
}