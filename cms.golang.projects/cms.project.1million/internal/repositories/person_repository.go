package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PersonRepository struct {
	RedisClient *redis.Client
	Key         string
}

func NewPersonRepository(redisClient *redis.Client) *PersonRepository {
	return &PersonRepository{
		RedisClient: redisClient,
		Key:         "persons",
	}
}

func (r *PersonRepository) Add(ctx context.Context, model models.PersonModel) error {
	payload, err := sonic.Marshal(model)
	if err != nil {
		return fmt.Errorf("failed to marshal person data: %w", err)
	}

	return r.RedisClient.HSet(ctx, r.Key, model.ID.String(), payload).Err()
}

func (r *PersonRepository) GetAll(ctx context.Context) ([]*models.PersonModel, error) {
	// HGetAll: Retorna todos os campos e valores do hash armazenado em key.
	personsData, err := r.RedisClient.HGetAll(ctx, r.Key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve persons: %w", err)
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
				return // continue
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
	personDataJSON, err := r.RedisClient.HGet(ctx, r.Key, id.String()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve person by ID: %w", err)
	}

	var person models.PersonModel
	err = sonic.Unmarshal([]byte(personDataJSON), &person)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal person data: %w", err)
	}

	return &person, nil
}

func (r *PersonRepository) GetCount(ctx context.Context) (int64, error) {
	count, err := r.RedisClient.HLen(ctx, r.Key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get person count: %w", err)
	}

	return count, nil
}
