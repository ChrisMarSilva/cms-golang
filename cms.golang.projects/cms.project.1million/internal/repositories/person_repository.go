package repositories

import (
	"context"
	"log/slog"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/google/uuid"
)

type PersonRepository struct {
	redisCache *stores.RedisCache
}

func NewPersonRepository(redisCache *stores.RedisCache) *PersonRepository {
	return &PersonRepository{
		redisCache: redisCache,
	}
}

func (r *PersonRepository) Add(ctx context.Context, model models.PersonModel) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.Add")
	defer span.End()

	payload, err := sonic.Marshal(model)
	if err != nil {
		slog.Error("Failed to marshal person data", slog.Any("error", err))
		return err
	}

	return r.redisCache.HSet(ctx, "persons", model.ID.String(), payload)
}

func (r *PersonRepository) GetAll(ctx context.Context) ([]*models.PersonModel, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetAll")
	defer span.End()

	personsData, err := r.redisCache.HGetAll(ctx, "persons")
	if err != nil {
		slog.Error("Failed to retrieve persons", slog.Any("error", err))
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	persons := make([]*models.PersonModel, 0, len(personsData))

	for _, personDataJSON := range personsData {
		wg.Add(1)
		payload := personDataJSON

		go func() {
			defer wg.Done()

			var person models.PersonModel
			err := sonic.Unmarshal([]byte(payload), &person)
			if err != nil {
				slog.Error("Failed to unmarshal person data", slog.Any("error", err))
				return
			}

			mu.Lock()
			persons = append(persons, &person)
			mu.Unlock()
		}()
	}

	wg.Wait()

	return persons, nil
}

func (r *PersonRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.PersonModel, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetByID")
	defer span.End()

	personDataJSON, err := r.redisCache.HGet(ctx, "persons", id.String())
	if err != nil {
		slog.Error("Failed to retrieve person by ID", slog.Any("error", err))
		return nil, err
	}

	var person models.PersonModel
	err = sonic.Unmarshal([]byte(personDataJSON), &person)
	if err != nil {
		slog.Error("Failed to unmarshal person data", slog.Any("error", err))
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) GetCount(ctx context.Context) (int64, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetCount")
	defer span.End()

	return r.redisCache.HLen(ctx, "persons")
}
