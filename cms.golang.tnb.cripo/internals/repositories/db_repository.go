package repositories

import (
	"context"
	"database/sql"
)

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error
}

type defaultRepository struct { // dbRepo
	Db *sql.DB
}

func NewDBRepo(db *sql.DB) *defaultRepository {
	return &defaultRepository{
		Db: db,
	}
}

func (repo *defaultRepository) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
	tx, err := repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() error {
		if err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		return nil
	}()

	if err := operation(ctx, tx); err != nil {
		return err
	}

	return nil
}
