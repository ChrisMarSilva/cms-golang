package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/google/uuid"
)

type PersonRepository struct {
	logger *slog.Logger
	db     *stores.Database
	//redisCache *stores.RedisCache
}

func NewPersonRepository(logger *slog.Logger, db *stores.Database, redisCache *stores.RedisCache) *PersonRepository {
	return &PersonRepository{
		logger: logger,
		db:     db,
		//redisCache: redisCache,
	}
}

func (r *PersonRepository) Add(ctx context.Context, model models.PersonModel) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.Add")
	defer span.End()

	query := "INSERT INTO \"TbPerson\" (id, name, created_at) Values ($1, $2, $3)"
	result, err := r.db.Conn.Exec(ctx, query, model.ID, model.Name, model.CreatedAt)
	if err != nil {
		r.logger.Error("Failed to insert client transaction", slog.Any("error", err))
		return err
	}

	if result.RowsAffected() != 1 {
		r.logger.Error("Unexpected number of rows affected", slog.Int64("rowsAffected", result.RowsAffected()))
		return errors.New("Failed to insert person")
	}

	// payload, err := sonic.Marshal(model)
	// if err != nil {
	// 	r.logger.Error("Failed to marshal person data", slog.Any("error", err))
	// 	return err
	// }

	// err = r.redisCache.HSet(ctx, "persons", model.ID.String(), payload)
	// if err != nil {
	// 	r.logger.Error("Failed to set person data in cache", slog.Any("error", err))
	// 	return err
	// }

	return nil
}

func (r *PersonRepository) GetAll(ctx context.Context) ([]*models.PersonModel, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetAll")
	defer span.End()

	query := "SELECT id, name, created_at FROM \"TbPerson\" ORDER BY created_at"
	rows, err := r.db.Conn.Query(ctx, query)
	if err != nil {
		r.logger.Error("Failed to query persons from database", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()

	persons := make([]*models.PersonModel, 0)

	for rows.Next() {
		var person models.PersonModel
		if err := rows.Scan(&person.ID, &person.Name, &person.CreatedAt); err != nil {
			r.logger.Error("Failed to scan person data", slog.Any("error", err))
			return nil, err
		}

		persons = append(persons, &person)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Failed to iterate over persons", slog.Any("error", err))
		return nil, err
	}

	// personsData, err := r.redisCache.HGetAll(ctx, "persons")
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

	return persons, nil
}

func (r *PersonRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.PersonModel, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetByID")
	defer span.End()

	var person models.PersonModel

	query := "SELECT id, name, created_at FROM \"TbPerson\" WHERE id = $1"
	row := r.db.Conn.QueryRow(ctx, query, id)

	if err := row.Scan(&person.ID, &person.Name, &person.CreatedAt); err != nil {
		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
			r.logger.Error("Person not found", slog.Any("error", err))
			return nil, errors.New("Not found.")
		}
		r.logger.Error("Failed to scan person data", slog.Any("error", err))
		return nil, err
	}

	// personDataJSON, err := r.redisCache.HGet(ctx, "persons", id.String())
	// if err != nil {
	// 	r.logger.Error("Failed to retrieve person by ID", slog.Any("error", err))
	// 	return nil, err
	// }

	// err = sonic.Unmarshal([]byte(personDataJSON), &person)
	// if err != nil {
	// 	r.logger.Error("Failed to unmarshal person data", slog.Any("error", err))
	// 	return nil, err
	// }

	return &person, nil
}

func (r *PersonRepository) Update(ctx context.Context, model *models.PersonModel) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.Update")
	defer span.End()

	query := "UPDATE \"TbPerson\" SET name = $1 WHERE id = $2"
	result, err := r.db.Conn.Exec(ctx, query, model.Name, model.ID)
	if err != nil {
		r.logger.Error("Failed to update person data", slog.Any("error", err))
		return err
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("No rows updated", slog.Any("id", model.ID))
		return errors.New("Not found.")
	}

	// payload, err := sonic.Marshal(model)
	// if err != nil {
	// 	r.logger.Error("Failed to marshal person data", slog.Any("error", err))
	// 	return err
	// }

	// err = r.redisCache.HSet(ctx, "persons", model.ID.String(), payload)
	// if err != nil {
	// 	r.logger.Error("Failed to update person data in cache", slog.Any("error", err))
	// 	return err
	// }

	return nil
}

func (r *PersonRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.DeleteByID")
	defer span.End()

	query := "DELETE FROM \"TbPerson\" WHERE id = $1"
	result, err := r.db.Conn.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete person by ID", slog.Any("error", err))
		return err
	}

	if result.RowsAffected() == 0 {
		r.logger.Error("No rows deleted", slog.Any("id", id))
		return errors.New("Not found.")
	}

	// err := r.redisCache.HDel(ctx, "persons", id.String())
	// if err != nil {
	// 	r.logger.Error("Failed to delete person by ID", slog.Any("error", err))
	// 	return err
	// }

	return nil
}

func (r *PersonRepository) ExistByID(ctx context.Context, id uuid.UUID) (bool, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.ExistByID")
	defer span.End()

	var exists bool

	query := "SELECT exists(SELECT 1 FROM \"TbPerson\" WHERE id = $1)"
	row := r.db.Conn.QueryRow(ctx, query, id)
	if err := row.Scan(&exists); err != nil {
		r.logger.Error("Failed to check if person exists", slog.Any("error", err))
		return false, err
	}

	return exists, nil
}

func (r *PersonRepository) GetCount(ctx context.Context) (int64, error) {
	ctx, span := utils.Tracer.Start(ctx, "PersonRepository.GetCount")
	defer span.End()

	var count int64

	query := "SELECT count(1) FROM \"TbPerson\""
	row := r.db.Conn.QueryRow(ctx, query)
	if err := row.Scan(&count); err != nil {
		r.logger.Error("Failed to get person count", slog.Any("error", err))
		return 0, err
	}

	// count, err := r.redisCache.HLen(ctx, "persons")
	// if err != nil {
	// 	r.logger.Error("Failed to get person count", slog.Any("error", err))
	// 	return 0, err
	// }

	return count, nil
}
