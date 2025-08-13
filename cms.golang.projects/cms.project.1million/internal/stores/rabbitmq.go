package stores

import (
	"log"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/wagslane/go-rabbitmq"
	//amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	// Conn    *amqp.Connection
	// Channel *amqp.Channel
	Conn      *rabbitmq.Conn
	Publisher *rabbitmq.Publisher
	Consumer  *rabbitmq.Consumer
}

func NewRabbitMQConnection(config *utils.Config) *RabbitMQ {
	url := config.RabbitMqUrl
	//url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.RabbitMqUser, config.RabbitMqPass, config.RabbitMqHost, config.RabbitMqPort, config.RabbitMqVhost)

	//conn, err := amqp.Dial(url) // amqp.DialConfig(url, amqp.Config{Heartbeat: 30 * time.Second})
	conn, err := rabbitmq.NewConn(url, rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		log.Fatal("RabbitMQ connect:", err)
	}

	// ch, err := conn.Channel()
	// if err != nil {
	// 	log.Fatal("RabbitMQ channel:", err)
	// }

	// _, err = ch.QueueDeclare(config.RabbitMqQueue, true, false, false, false, amqp.Table{})
	// if err != nil {
	// 	log.Fatal("RabbitMQ queue declare:", err)
	// }

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(""),
	)
	if err != nil {
		log.Fatal("RabbitMQ publisher:", err)
	}

	consumer, err := rabbitmq.NewConsumer(
		conn,
		config.RabbitMqQueue,
		rabbitmq.WithConsumerOptionsExchangeName(""),
		//rabbitmq.WithConsumerOptionsRoutingKey(""),
	)
	if err != nil {
		log.Fatal("RabbitMQ consumer:", err)
	}

	log.Println("RabbitMQ connection established successfully")
	return &RabbitMQ{
		Conn: conn,
		//Channel: ch,
		Publisher: publisher,
		Consumer:  consumer,
	}
}

func (rabbitMQ *RabbitMQ) CloseConnection() {
	log.Println("Closing RabbitMQ connection...")
	//rabbitMQ.Channel.Close()
	rabbitMQ.Publisher.Close()
	rabbitMQ.Consumer.Close()
	rabbitMQ.Conn.Close()
}
