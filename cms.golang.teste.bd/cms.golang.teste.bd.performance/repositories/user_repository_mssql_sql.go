package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
)

type UserRepositorySQLServerSQL struct {
	db *sql.DB
}

func NewUserRepositorySQLServerSQL(db *sql.DB) *UserRepositorySQLServerSQL {
	return &UserRepositorySQLServerSQL{db: db}
}

func (repo UserRepositorySQLServerSQL) GetUserRepository() (err error) {
	return nil
}

func (repo UserRepositorySQLServerSQL) GetAll(ctx context.Context, users *[]entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) ORDER BY ID"

	rows, err := repo.db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Nome, &user.Status)
		if err != nil {
			return err
		}
		*users = append(*users, user)
	}

	return nil
}

func (repo UserRepositorySQLServerSQL) Get(ctx context.Context, user *entity.User, id uuid.UUID) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) WHERE ID = ?"

	row, err := repo.db.Query(query, id)
	if err != nil {
		return err
	}

	defer row.Close()

	row.Next()

	err = row.Scan(&user.ID, &user.Nome, &user.Status)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQL) Create(ctx context.Context, user *entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// query := `INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES (@ID, @NOME, @STATUS);`
	// stmt, err := repo.db.Prepare(query)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()
	// row := stmt.QueryRowContext(ctx, sql.Named("ID", user.ID), sql.Named("NOME", user.Nome), sql.Named("STATUS", user.Status))
	// err = row.Scan()
	// if err != nil {
	// 	return err
	// }

	query := fmt.Sprintf("INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES ('%s','%s','%s');", user.ID, user.Nome, user.Status)
	_, err = repo.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQL) CreateInBatch(ctx context.Context, users []entity.User) (err error) {

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	for _, user := range users {
		// err = tx.QueryRow("SELECT id FROM pages WHERE url = $1", r.URL.Path).Scan(&id)
		// _, err = tx.Exec("UPDATE pages SET visitors = visitors + 1 WHERE id = $1", id)
		query := fmt.Sprintf("INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES ('%s','%s','%s');", user.ID, user.Nome, user.Status)
		_, err = tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo UserRepositorySQLServerSQL) Update(ctx context.Context, user *entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// query := fmt.Sprintf("UPDATE TBUSUARIO SET NOME = @NOME, STATUS = @STATUS WHERE ID = @ID")
	// _, err = repo.db.ExecContext(ctx, query, sql.Named("NOME", user.Nome), sql.Named("STATUS", user.Status), sql.Named("ID", user.ID))
	// if err != nil {
	// 	return err
	// }

	query := fmt.Sprintf("UPDATE TBUSUARIO SET NOME = '%s', STATUS = '%s' WHERE ID = '%s';", user.Nome, user.Status, user.ID)

	_, err = repo.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQL) Delete(ctx context.Context, id uuid.UUID) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// query := fmt.Sprintf("DELETE FROM TBUSUARIO WHERE ID = @ID;")
	// _, err = repo.db.ExecContext(ctx, query, sql.Named("ID", id))
	// if err != nil {
	// 	return err
	// }

	query := fmt.Sprintf("DELETE FROM TBUSUARIO WHERE ID = '%s';", id)

	_, err = repo.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
