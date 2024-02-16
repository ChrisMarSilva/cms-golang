package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IClientRepository interface {
	Get(entity *models.Cliente, id int) error
	UpdSaldo(id int, valor int64, tipo string) error
}

type ClientRepository struct {
	db *sqlx.DB
}

func NewClientRepository(db *sqlx.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (repo ClientRepository) Get(entity *models.Cliente, id int) (err error) {
	query := "SELECT limite, saldo FROM cliente WHERE id = $1"

	// row, err := repo.db.QueryContext(context.Background(), query, id)
	// if err != nil {
	// 	return err
	// }
	// defer row.Close()

	// if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
	// 	return err
	// }

	// if row.Next() {
	// 	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	return errors.New("Cliente não localizado.")
	// }

	// if err := row.Err(); err != nil {
	// 	return err
	// }

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), id)
	// row := repo.db.QueryRowContext(context.Background(), query, id)

	if err := row.Err(); err != nil {
		return err
	}

	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
			return errors.New("Cliente não localizado.")
		}
		return err
	}

	return nil
}

func (repo ClientRepository) UpdSaldo(id int, valor int64, tipo string) (err error) {
	var query string

	if tipo == "d" {
		query = "UPDATE cliente SET saldo = saldo - $1 WHERE id = $2"
	} else {
		query = "UPDATE cliente SET saldo = saldo + $1 WHERE id = $2"
	}

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// result, err := repo.db.ExecContext(context.Background(), query, valor, id)
	result, err := stmt.ExecContext(context.Background(), valor, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Saldo do cliente não atualizado.")
	}

	return nil
}
