package stores

import (
	"log/slog"
	"os"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQ struct {
	conn      *rabbitmq.Conn
	Publisher *rabbitmq.Publisher
	Consumer  *rabbitmq.Consumer
}

func NewRabbitMQ(config *utils.Config) *RabbitMQ {
	url := config.RabbitMqUrl // fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.RabbitMqUser, config.RabbitMqPass, config.RabbitMqHost, config.RabbitMqPort, config.RabbitMqVhost)
	conn, err := rabbitmq.NewConn(url, rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		slog.Error("RabbitMQ connection error", "error", err)
		os.Exit(1)
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsConfirm,
		rabbitmq.WithPublisherOptionsExchangeName(""),
	)
	if err != nil {
		slog.Error("RabbitMQ publisher error", "error", err)
		os.Exit(1)
	}

	consumer, err := rabbitmq.NewConsumer(
		conn,
		config.RabbitMqQueue,
		rabbitmq.WithConsumerOptionsConcurrency(config.NumConsumerWorkers),
		rabbitmq.WithConsumerOptionsExchangeName(""),
	)
	if err != nil {
		slog.Error("RabbitMQ consumer error", "error", err)
		os.Exit(1)
	}

	return &RabbitMQ{
		Publisher: publisher,
		Consumer:  consumer,
	}
}

func (r *RabbitMQ) Close() {
	r.Publisher.Close()
	r.Consumer.Close()
	r.conn.Close()
}
