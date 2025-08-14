package services

import (
	"context"

	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
	"github.com/google/uuid"
	//amqp "github.com/rabbitmq/amqp091-go"
)

type PersonService struct {
	Repo           *repositories.PersonRepository
	RabbitMQClient *stores.RabbitMQ
}

func NewPersonService(repo *repositories.PersonRepository, rabbitMQClient *stores.RabbitMQ) *PersonService {
	return &PersonService{
		Repo:           repo,
		RabbitMQClient: rabbitMQClient,
	}
}

func (s *PersonService) Add(ctx context.Context, request dtos.PersonRequestDto) error {
	//model := models.NewPersonModel(request.Name)

	// payload, err := sonic.Marshal(request)
	// if err != nil {
	// 	return err
	// }

	// q, err := s.RabbitMQClient.Channel.QueueDeclare(s.Config.RabbitMqQueue, true, false, false, false, amqp.Table{})
	// if err != nil {
	// 	return err
	// }

	// if s.RabbitMQClient.Conn.IsClosed() || s.RabbitMQClient.Channel.IsClosed() {
	// 	s.RabbitMQClient.CloseConnection()
	// 	s.RabbitMQClient = stores.NewRabbitMQConnection(s.Config)
	// }

	// message := amqp.Publishing{ContentType: "application/json", DeliveryMode: amqp.Persistent, Body: payload}
	//err = s.RabbitMQClient.Channel.PublishWithContext(ctx, "", s.Config.RabbitMqQueue, false, false, message)
	// err = s.RabbitMQClient.Publisher.PublishWithContext(
	// 	ctx,
	// 	payload,
	// 	[]string{s.Config.RabbitMqQueue},
	// 	rabbitmq.WithPublishOptionsContentType("application/json"),
	// 	rabbitmq.WithPublishOptionsMandatory,
	// 	rabbitmq.WithPublishOptionsPersistentDelivery,
	// 	rabbitmq.WithPublishOptionsExchange(""),
	// )
	// if err != nil {
	// 	return err
	// }

	workers.EventPublisher <- request

	// err := s.Repo.Add(ctx, *model)
	// if err != nil {
	// 	return nil, err
	// }

	//person := &dtos.PersonResponseDto{ID: model.ID, Name: model.Name, RequestedAt: model.RequestedAt}
	//return person, nil
	return nil
}

func (s *PersonService) GetAll(ctx context.Context) ([]*dtos.PersonResponseDto, error) {
	personsModels, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	personsDtos := make([]*dtos.PersonResponseDto, 0, len(personsModels))

	for _, personModel := range personsModels {
		personDto := &dtos.PersonResponseDto{ID: personModel.ID, Name: personModel.Name, RequestedAt: personModel.RequestedAt}
		personsDtos = append(personsDtos, personDto)
	}

	return personsDtos, nil
}

func (s *PersonService) GetByID(ctx context.Context, id uuid.UUID) (*dtos.PersonResponseDto, error) {
	personModel, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	personDto := &dtos.PersonResponseDto{ID: personModel.ID, Name: personModel.Name, RequestedAt: personModel.RequestedAt}
	return personDto, nil
}

func (s *PersonService) GetCount(ctx context.Context) (int64, error) {
	count, err := s.Repo.GetCount(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
