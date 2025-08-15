package services

import (
	"context"
	"log/slog"

	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
	"github.com/google/uuid"
)

type PersonService struct {
	repo *repositories.PersonRepository
}

func NewPersonService(repo *repositories.PersonRepository) *PersonService {
	return &PersonService{
		repo: repo,
	}
}

func (s *PersonService) Add(ctx context.Context, request dtos.PersonRequestDto) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.Add")
	defer span.End()

	// err := s.repo.Add(ctx, *model)
	// if err != nil {
	// 	slog.Error("Failed to add person to repository", slog.Any("error", err))
	// 	return nil, err
	// }

	workers.EventPublisher <- request

	return nil
}

func (s *PersonService) GetAll(ctx context.Context) ([]*dtos.PersonResponseDto, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetAll")
	defer span.End()

	personsModels, err := s.repo.GetAll(ctx)
	if err != nil {
		slog.Error("Failed to get all persons", slog.Any("error", err))
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
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetByID")
	defer span.End()

	personModel, err := s.repo.GetByID(ctx, id)
	if err != nil {
		slog.Error("Failed to get person by ID", slog.Any("error", err))
		return nil, err
	}

	personDto := &dtos.PersonResponseDto{ID: personModel.ID, Name: personModel.Name, RequestedAt: personModel.RequestedAt}
	return personDto, nil
}

func (s *PersonService) GetCount(ctx context.Context) (int64, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonService.GetCount")
	defer span.End()

	count, err := s.repo.GetCount(ctx)
	if err != nil {
		slog.Error("Failed to get person count", slog.Any("error", err))
		return 0, err
	}

	return count, nil
}
