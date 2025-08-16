package services

import (
	"context"
	"log/slog"

	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
	"github.com/google/uuid"
)

type PersonService struct {
	logger *slog.Logger
	repo   *repositories.PersonRepository
}

func NewPersonService(logger *slog.Logger, repo *repositories.PersonRepository) *PersonService {
	return &PersonService{
		logger: logger,
		repo:   repo,
	}
}

func (s *PersonService) Add(ctx context.Context, request dtos.PersonRequestDto) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.Add")
	defer span.End()

	workers.EventPublisher <- request

	model := models.NewPersonModel(request.Name)

	err := s.repo.Add(ctx, *model)
	if err != nil {
		s.logger.Error("Failed to add person", slog.Any("error", err))
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

	return personsDtos, nil
}

func (s *PersonService) GetByID(ctx context.Context, id uuid.UUID) (*dtos.PersonResponseDto, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetByID")
	defer span.End()

	personModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get person by ID", slog.Any("error", err))
		return nil, err
	}

	personDto := dtos.NewPersonResponseDto(
		personModel.ID,
		personModel.Name,
		personModel.CreatedAt,
	)

	return personDto, nil
}

func (s *PersonService) Update(ctx context.Context, id uuid.UUID, request dtos.PersonRequestDto) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.Update")
	defer span.End()

	model := models.NewPersonModel(request.Name)
	model.ID = id

	err := s.repo.Update(ctx, model)
	if err != nil {
		s.logger.Error("Failed to update person", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *PersonService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.DeleteByID")
	defer span.End()

	err := s.repo.DeleteByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete person by ID", slog.Any("error", err))
		return err
	}

	return nil
}

func (s *PersonService) ExistByID(ctx context.Context, id uuid.UUID) (bool, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.ExistByID")
	defer span.End()

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

	return count, nil
}
