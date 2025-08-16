package services

import (
	"context"
	"log/slog"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

type PersonService struct {
	logger *slog.Logger
	repo   *repositories.PersonRepository
	rdb    *stores.RedisCache
}

func NewPersonService(logger *slog.Logger, repo *repositories.PersonRepository, rdb *stores.RedisCache) *PersonService {
	return &PersonService{
		logger: logger,
		repo:   repo,
		rdb:    rdb,
	}
}

func (s *PersonService) Add(ctx context.Context, request dtos.PersonRequestDto) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.Add")
	defer span.End()
	span.SetAttributes(attribute.String("name", request.Name))

	workers.EventPublisher <- request

	model := models.NewPersonModel(request.Name)

	err := s.repo.Add(ctx, *model)
	if err != nil {
		s.logger.Error("Failed to add person", slog.Any("error", err))
		return err
	}

	payload, err := sonic.Marshal(model)
	if err != nil {
		s.logger.Error("Failed to marshal person data", slog.Any("error", err))
		return err
	}

	err = s.rdb.HSet(ctx, "persons", model.ID.String(), payload)
	if err != nil {
		s.logger.Error("Failed to set person data in cache", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *PersonService) GetAll(ctx context.Context) ([]*dtos.PersonResponseDto, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetAll")
	defer span.End()

	personsModels, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get all persons", slog.Any("error", err))
		return nil, err
	}

	personsDtos := make([]*dtos.PersonResponseDto, 0, len(personsModels))

	for _, personModel := range personsModels {
		personDto := dtos.NewPersonResponseDto(
			personModel.ID,
			personModel.Name,
			personModel.CreatedAt,
		)

		personsDtos = append(personsDtos, personDto)
	}

	// personsData, err := r.rdb.HGetAll(ctx, "persons")
	// if err != nil {
	// 	r.logger.Error("Failed to retrieve persons", slog.Any("error", err))
	// 	return nil, err
	// }

	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// persons := make([]*models.PersonModel, 0, len(personsData))

	// for _, personDataJSON := range personsData {
	// 	wg.Add(1)

	// 	go func(payload string) {
	// 		defer wg.Done()

	// 		var person models.PersonModel
	// 		err := sonic.Unmarshal([]byte(payload), &person)
	// 		if err != nil {
	// 			r.logger.Error("Failed to unmarshal person data", slog.Any("error", err))
	// 			return
	// 		}

	// 		mu.Lock()
	// 		persons = append(persons, &person)
	// 		mu.Unlock()
	// 	}(personDataJSON)
	// }
	// wg.Wait()

	return personsDtos, nil
}

func (s *PersonService) GetByID(ctx context.Context, id uuid.UUID) (*dtos.PersonResponseDto, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetByID")
	defer span.End()
	span.SetAttributes(attribute.String("id", id.String()))

	personDataJSON, _ := s.rdb.HGet(ctx, "persons", id.String())
	if personDataJSON != "" {
		s.logger.Info("Cache hit for person ID", slog.String("id", id.String()))

		var personModel models.PersonModel
		err := sonic.Unmarshal([]byte(personDataJSON), &personModel)
		if err != nil {
			s.logger.Error("Failed to unmarshal person data", slog.Any("error", err))
			return nil, err
		}

		personDto := dtos.NewPersonResponseDto(personModel.ID, personModel.Name, personModel.CreatedAt)
		return personDto, nil
	}

	s.logger.Info("Cache miss for person ID", slog.String("id", id.String()))

	personModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get person by ID", slog.Any("error", err))
		return nil, err
	}

	personDto := dtos.NewPersonResponseDto(personModel.ID, personModel.Name, personModel.CreatedAt)
	return personDto, nil
}

func (s *PersonService) Update(ctx context.Context, id uuid.UUID, request dtos.PersonRequestDto) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.Update")
	defer span.End()
	span.SetAttributes(attribute.String("id", id.String()), attribute.String("name", request.Name))

	model := models.NewPersonModel(request.Name)
	model.ID = id

	err := s.repo.Update(ctx, model)
	if err != nil {
		s.logger.Error("Failed to update person", slog.Any("error", err))
		return err
	}

	payload, err := sonic.Marshal(model)
	if err != nil {
		s.logger.Error("Failed to marshal person data", slog.Any("error", err))
		return err
	}

	err = s.rdb.HSet(ctx, "persons", model.ID.String(), payload)
	if err != nil {
		s.logger.Error("Failed to set person data in cache", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *PersonService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.DeleteByID")
	defer span.End()
	span.SetAttributes(attribute.String("id", id.String()))

	err := s.repo.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete person by ID", slog.Any("error", err))
		return err
	}

	err = s.rdb.HDel(ctx, "persons", id.String())
	if err != nil {
		s.logger.Error("Failed to delete person from cache", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *PersonService) ExistByID(ctx context.Context, id uuid.UUID) (bool, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.ExistByID")
	defer span.End()
	span.SetAttributes(attribute.String("id", id.String()))

	exists, err := s.repo.ExistByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to check if person exists by ID", slog.Any("error", err))
		return false, err
	}

	return exists, nil
}

func (s *PersonService) GetCount(ctx context.Context) (int64, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetCount")
	defer span.End()

	count, err := s.repo.GetCount(ctx)
	if err != nil {
		s.logger.Error("Failed to get person count", slog.Any("error", err))
		return 0, err
	}

	// count, err := r.rdb.HLen(ctx, "persons")
	// if err != nil {
	// 	r.logger.Error("Failed to get person count", slog.Any("error", err))
	// 	return 0, err
	// }

	return count, nil
}
