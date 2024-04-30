package repositories

import (
	"context"
	"database/sql"

	data "github.com/chrismarsilva/cms.golang.tnb.cripo.database"
)

type IDbRepository interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error
}

type DbRepository struct {
	Db *data.Database
}

func NewDbRepository(db *data.Database) *DbRepository {
	return &DbRepository{
		Db: db,
	}
}

func (this *DbRepository) Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error {
	tx, err := this.Db.BeginTx(ctx, nil)
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
