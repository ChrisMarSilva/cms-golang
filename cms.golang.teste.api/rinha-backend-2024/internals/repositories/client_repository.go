package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IClientRepository interface {
	Get(ctx context.Context, entity *models.Cliente, id int) error
	// GetWithPrepare(entity *models.Cliente, id int) error
	// GetWithouPrepare(entity *models.Cliente, id int) error
	// GetWithPgx(entity *models.Cliente, id int) error

	// UpdSaldoWithPrepare(tx *sqlx.Tx, id int, valor int64, tipo string) error
	// UpdSaldoWithouPrepare(tx *sqlx.Tx, id int, valor int64, tipo string) error
	// UpdSaldoWithPgx(tx *sql.Tx, id int, valor int64, tipo string) error
	UpdSaldo(ctx context.Context, tx pgx.Tx, id int, valor int64, tipo string) error
}

type ClientRepository struct {
	// db *sqlx.DB
	// writer *sqlx.DB
	// reader *sqlx.DB
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db: db}
}

func (repo ClientRepository) Get(ctx context.Context, entity *models.Cliente, id int) (err error) {
	query := "SELECT limite, saldo FROM cliente WHERE id = $1"
	row := repo.db.QueryRow(ctx, query, id)

	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
			return errors.New("Cliente não localizado.")
		}
		return err
	}

	return nil
}

// func (repo ClientRepository) GetWithPrepare(entity *models.Cliente, id int) (err error) {
// 	query := "SELECT limite, saldo FROM cliente WHERE id = $1"
//
// 	// repo.db.QueryRowxContext(context.Background(),query)
// 	// row, err := repo.db.QueryContext(context.Background(), query, id)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// defer row.Close()
//
// 	// if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
// 	// 	return err
// 	// }
//
// 	// if row.Next() {
// 	// 	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
// 	// 		return err
// 	// 	}
// 	// } else {
// 	// 	return errors.New("Cliente não localizado.")
// 	// }
//
// 	// if err := row.Err(); err != nil {
// 	// 	return err
// 	// }
//
// 	stmt, err := repo.db.PrepareContext(context.Background(), query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	row := stmt.QueryRowContext(context.Background(), id)
// 	// row := repo.db.QueryRowContext(context.Background(), query, id)
//
// 	if err := row.Err(); err != nil {
// 		return err
// 	}
//
// 	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
// 		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
// 			return errors.New("Cliente não localizado.")
// 		}
// 		return err
// 	}
//
// 	return nil
// }

// func (repo ClientRepository) GetWithouPrepare(entity *models.Cliente, id int) (err error) {
// 	query := "SELECT limite, saldo FROM cliente WHERE id = $1"
//
// 	row := repo.db.QueryRowContext(context.Background(), query, id)
// 	if err := row.Err(); err != nil {
// 		return err
// 	}
//
// 	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
// 		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
// 			return errors.New("Cliente não localizado.")
// 		}
// 		return err
// 	}
//
// 	return nil
// }

// func (repo ClientRepository) GetWithPgx(entity *models.Cliente, id int) (err error) {
// 	query := "SELECT limite, saldo FROM cliente WHERE id = $1"
// 	row := repo.dbPgx.QueryRow(context.Background(), query, id)
//
// 	if err := row.Scan(&entity.Limite, &entity.Saldo); err != nil {
// 		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
// 			return errors.New("Cliente não localizado.")
// 		}
// 		return err
// 	}
//
// 	return nil
// }

func (repo ClientRepository) UpdSaldo(ctx context.Context, tx pgx.Tx, id int, valor int64, tipo string) (err error) {
	var query string

	if tipo == "d" {
		query = "UPDATE cliente SET saldo = saldo - $1 WHERE id = $2"
	} else {
		query = "UPDATE cliente SET saldo = saldo + $1 WHERE id = $2"
	}

	result, err := tx.Exec(ctx, query, valor, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		return errors.New("Saldo do cliente não atualizado.")
	}

	return nil
}

// func (repo ClientRepository) UpdSaldoWithPrepare(tx *sqlx.Tx, id int, valor int64, tipo string) (err error) {
// 	var query string
//
// 	if tipo == "d" {
// 		query = "UPDATE cliente SET saldo = saldo - $1 WHERE id = $2"
// 	} else {
// 		query = "UPDATE cliente SET saldo = saldo + $1 WHERE id = $2"
// 	}
//
// 	//stmt, err := repo.db.PrepareContext(context.Background(), query)
// 	stmt, err := tx.PrepareContext(context.Background(), query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	// result, err := repo.db.ExecContext(context.Background(), query, valor, id)
// 	result, err := stmt.ExecContext(context.Background(), valor, id)
// 	if err != nil {
// 		return err
// 	}
//
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
//
// 	if rowsAffected != 1 {
// 		return errors.New("Saldo do cliente não atualizado.")
// 	}
//
// 	return nil
// }

// func (repo ClientRepository) UpdSaldoWithouPrepare(tx *sqlx.Tx, id int, valor int64, tipo string) (err error) {
// 	var query string
//
// 	if tipo == "d" {
// 		query = "UPDATE cliente SET saldo = saldo - $1 WHERE id = $2"
// 	} else {
// 		query = "UPDATE cliente SET saldo = saldo + $1 WHERE id = $2"
// 	}
//
// 	result, err := repo.db.ExecContext(context.Background(), query, valor, id)
// 	if err != nil {
// 		return err
// 	}
//
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
//
// 	if rowsAffected != 1 {
// 		return errors.New("Saldo do cliente não atualizado.")
// 	}
//
// 	return nil
// }

// func (repo ClientRepository) UpdSaldoWithPgx(tx *sql.Tx, id int, valor int64, tipo string) (err error) {
// 	var query string
//
// 	if tipo == "d" {
// 		query = "UPDATE cliente SET saldo = saldo - $1 WHERE id = $2"
// 	} else {
// 		query = "UPDATE cliente SET saldo = saldo + $1 WHERE id = $2"
// 	}
//
// 	result, err := repo.dbPgx.Exec(context.Background(), query, valor, id)
// 	if err != nil {
// 		return err
// 	}
//
// 	rowsAffected := result.RowsAffected()
// 	if rowsAffected != 1 {
// 		return errors.New("Saldo do cliente não atualizado.")
// 	}
//
// 	return nil
// }
