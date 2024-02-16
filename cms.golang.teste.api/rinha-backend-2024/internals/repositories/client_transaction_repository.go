package repositories

import (
	"context"
	"errors"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IClientTransactionRepository interface {
	GetAll(entities *[]models.ClienteTransacao, idcliente int) error
	Add(idcliente int, valor int64, tipo string, descricao string) error
}

type ClientTransactionRepository struct {
	db *sqlx.DB
}

func NewClientTransactionRepository(db *sqlx.DB) *ClientTransactionRepository {
	return &ClientTransactionRepository{db: db}
}

func (repo ClientTransactionRepository) GetAll(entities *[]models.ClienteTransacao, idcliente int) (err error) {
	query := "SELECT valor, tipo, descricao, dthrregistro FROM cliente_transacao WHERE cliente_id = $1 ORDER BY id DESC LIMIT 10" // ?

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// rows, err := repo.db.QueryxContext(context.Background(), query, idcliente)
	rows, err := stmt.QueryContext(context.Background(), idcliente)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		entity := models.ClienteTransacao{}
		if err := rows.Scan(&entity.Valor, &entity.Tipo, &entity.Descricao, &entity.DtHrRegistro); err != nil { // if err := rows.StructScan(&entity); err != nil {
			return err
		}
		*entities = append(*entities, entity)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (repo ClientTransactionRepository) Add(idcliente int, valor int64, tipo string, descricao string) (err error) {
	query := "INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao) Values ($1, $2, $3, $4)"

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//result, err := repo.db.ExecContext(context.Background(), query, idcliente, valor, tipo, descricao)
	result, err := stmt.ExecContext(context.Background(), idcliente, valor, tipo, descricao)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("Transação do cliente não inserida.")
	}

	return nil
}
