package main

import (
	"context"
	"database/sql"
)

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error
}

type dbRepo struct { // dbRepo // defaultRepository
	db *Database //  Db *sql.DB
}

func NewDBRepo(db *Database) *dbRepo {
	return &dbRepo{
		db: db,
	}
}

func (repo *dbRepo) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
	tx, err := repo.db.BeginTx(ctx, nil)
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
