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

/*

type Repository struct {
	Conn      *pgxpool.Pool
	ChPessoas chan rinha.Pessoa
}

var Repo *Repository

func NewRepository(Conn *pgxpool.Pool, Cache *redis.Client) *Repository {
	if Repo == nil {
		Repo = &Repository{Conn: Conn, Cache: Cache, ChPessoas: make(chan rinha.Pessoa)}
	}
	return Repo
}

func (r *Repository) Create(ctx context.Context, pessoa rinha.Pessoa) error {
	r.ChPessoas <- pessoa
	return nil
}

func (r *Repository) Insert(pessoas []rinha.Pessoa) error {
	if len(pessoas) == 0 {
		return nil
	}

	_, err := r.Conn.CopyFrom(
		context.Background(),
		pgx.Identifier{"pessoas"},
		[]string{"id", "apelido", "nome", "nascimento", "stack", "search_index"},
		pgx.CopyFromSlice(len(pessoas), func(i int) ([]any, error) {
			p := pessoas[i]
			index := fmt.Sprintf("%s %s %s", strings.ToLower(p.Apelido), strings.ToLower(p.Nome), strings.ToLower(strings.Join(p.Stack, " ")))
			return []any{p.ID, p.Apelido, p.Nome, p.Nascimento.Time, p.Stack, index}, nil
		}),
	)

	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.ConstraintName == "pessoas_apelido_key" {
		slog.Error("algum apelido ja existe")
		return pgErr
	}

	return err
}

func (r *Repository) FindOne(ctx context.Context, id uuid.UUID) (rinha.Pessoa, error) {
	var pessoa rinha.Pessoa
	var nascimento time.Time
	err = r.Conn.QueryRow(ctx, `SELECT id, apelido, nome, nascimento, stack FROM pessoas WHERE id = $1 LIMIT 1`, id).Scan(&pessoa.ID, &pessoa.Apelido, &pessoa.Nome, &nascimento, &pessoa.Stack)
	pessoa.Nascimento = rinha.Date{Time: nascimento}
	if err == pgx.ErrNoRows {
		return pessoa, nil
	}
	return pessoa, err
}

func (r *Repository) FindByTermo(ctx context.Context, t string) ([]rinha.Pessoa, error) {
	pessoas := []rinha.Pessoa{}
	rows, err := r.Conn.Query(ctx, ` SELECT id, apelido, nome, nascimento, stack FROM pessoas WHERE search_index ILIKE '%' || $1 || '%' LIMIT 50 `, strings.ToLower(t))
	if err != nil {
		return pessoas, err
	}
	defer rows.Close()x
	for rows.Next() {
		var pessoa rinha.Pessoa
		var nascimento time.Time
		err := rows.Scan(&pessoa.ID, &pessoa.Apelido, &pessoa.Nome, &nascimento, &pessoa.Stack)
		if err != nil {
			slog.Error(err.Error())
		}
		pessoas = append(pessoas, pessoa)
	}
	return pessoas, err
}

*/
