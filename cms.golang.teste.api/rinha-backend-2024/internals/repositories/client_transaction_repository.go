package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IClientTransactionRepository interface {
	GetAll(ctx context.Context, entities *map[int]models.ClienteTransacao, idcliente int) error
	// GetAllWithPrepare(entities *map[int]models.ClienteTransacao, idcliente int) error
	Add(ctx context.Context, tx pgx.Tx, idcliente int, valor int64, tipo string, descricao string) error
	// AddWithPrepare(tx *sqlx.Tx, idcliente int, valor int64, tipo string, descricao string) error
	SaveTransactionBatch(transactionBatch []models.ClienteTransacao) error
}

type ClientTransactionRepository struct {
	// db *sqlx.DB
	// writer *sqlx.DB
	// reader *sqlx.DB
	db *pgxpool.Pool
}

func NewClientTransactionRepository(db *pgxpool.Pool) *ClientTransactionRepository {
	return &ClientTransactionRepository{db: db}
}

func (repo ClientTransactionRepository) GetAll(ctx context.Context, entities *map[int]models.ClienteTransacao, idcliente int) (err error) {
	*entities = make(map[int]models.ClienteTransacao, 0)

	query := "SELECT valor, tipo, descricao, dthrregistro FROM cliente_transacao WHERE cliente_id = $1 ORDER BY id DESC LIMIT 10"
	rows, err := repo.db.Query(ctx, query, idcliente)
	if err != nil {
		return err
	}
	defer rows.Close()

	var totalRows int
	for rows.Next() {
		totalRows++ // Obter o número total de linhas retornadas pela consulta
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("erro ao contar linhas: %w", err)
	}

	//entities := make([]*models.ClienteTransacao, 0, totalRows)
	// if err := rows.Err(); err != nil {
	//     return fmt.Errorf("erro ao iterar sobre transacoes: %w", err)
	// }

	i := 0
	for rows.Next() {
		//entity := models.ClienteTransacao{}
		var entity models.ClienteTransacao
		if err := rows.Scan(&entity.Valor, &entity.Tipo, &entity.Descricao, &entity.DtHrRegistro); err != nil {
			return err
		}
		(*entities)[i] = entity
		//entities = append(entities, &entity)
		i++
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("erro ao iterar sobre transacoes: %w", err)
	}

	return nil
}

// func (repo ClientTransactionRepository) GetAllWithPrepare(entities *map[int]models.ClienteTransacao, idcliente int) (err error) {
// 	*entities = make(map[int]models.ClienteTransacao, 0)
//
// 	query := "SELECT valor, tipo, descricao, dthrregistro FROM cliente_transacao WHERE cliente_id = $1 ORDER BY id DESC LIMIT 10"
// 	//repo.db.QueryxContext(context.Background(),query)
//
// 	//users *[]entity.User
// 	// rows, err := repo.db.Queryx(query)
// 	// defer rows.Close()
// 	// for rows.Next() {
// 	// 	var user entity.User
// 	// 	if err := rows.StructScan(&user); err != nil { return err }
// 	// 	*users = append(*users, user)
// 	// }
//
// 	stmt, err := repo.db.PrepareContext(context.Background(), query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	// rows, err := repo.db.QueryxContext(context.Background(), query, idcliente)
// 	rows, err := stmt.QueryContext(context.Background(), idcliente)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()
//
// 	i := 0
// 	for rows.Next() {
// 		entity := models.ClienteTransacao{}
// 		if err := rows.Scan(&entity.Valor, &entity.Tipo, &entity.Descricao, &entity.DtHrRegistro); err != nil { // if err := rows.StructScan(&entity); err != nil {
// 			return err
// 		}
// 		(*entities)[i] = entity
// 		i++
// 	}
//
// 	if err := rows.Err(); err != nil {
// 		return err
// 	}
//
// 	return nil
// }

func (repo ClientTransactionRepository) Add(ctx context.Context, tx pgx.Tx, idcliente int, valor int64, tipo string, descricao string) (err error) {
	query := "INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao) Values ($1, $2, $3, $4)"

	result, err := tx.Exec(ctx, query, idcliente, valor, tipo, descricao)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		return errors.New("Transação do cliente não inserida.")
	}

	return nil
}

// func (repo ClientTransactionRepository) AddWithPrepare(tx *sqlx.Tx, idcliente int, valor int64, tipo string, descricao string) (err error) {
// 	query := "INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao) Values ($1, $2, $3, $4)"
//
// 	// res := repo.db.MustExec(query, user.ID, user.Nome, user.Status)
// 	// _, err = res.RowsAffected()
//
// 	//result, err := repo.db.ExecContext(context.Background(), query, idcliente, valor, tipo, descricao)
//
// 	//stmt, err := repo.db.PrepareContext(context.Background(), query)
// 	stmt, err := tx.PrepareContext(context.Background(), query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	result, err := stmt.ExecContext(context.Background(), idcliente, valor, tipo, descricao)
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
// 		return errors.New("Transação do cliente não inserida.")
// 	}
//
// 	return nil
// }

func (repo ClientTransactionRepository) SaveTransactionBatch(transactionBatch []models.ClienteTransacao) error {

	sql := "INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao) Values %s"
	//INSERT INTO cliente_transacao (cliente_id, valor, tipo, descricao) Values ($1, $2, $3, $4)

	params := []interface{}{}
	paramSql := []string{}

	for index, transaction := range transactionBatch {
		i := index * 4
		params = append(params, transaction.IdCliente, transaction.Valor, transaction.Tipo, transaction.Descricao, transaction.DtHrRegistro)
		paramSql = append(paramSql, fmt.Sprintf("($%d, $%d, $%d, $%d)", i+1, i+2, i+3, i+4))
	}

	_, err := repo.db.Exec(context.Background(), fmt.Sprintf(sql, strings.Join(paramSql, ",")), params...)
	if err != nil {
		log.Println(fmt.Sprintf("error executing insert %v", err))
		return err
	}

	return nil
}
