package repository

import (
	"context"
	"errors"

	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepositorySQLServerSQLX struct {
	db *sqlx.DB
}

func NewUserRepositorySQLServerSQLX(db *sqlx.DB) *UserRepositorySQLServerSQLX {
	return &UserRepositorySQLServerSQLX{db: db}
}

func (repo UserRepositorySQLServerSQLX) GetUserRepository() (err error) {
	return nil
}

func (repo UserRepositorySQLServerSQLX) GetAll(ctx context.Context, users *[]entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) ORDER BY ID"

	// err = repo.db.Select(users, query)
	// if err != nil {
	// 	return err
	// }

	rows, err := repo.db.Queryx(query)

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.StructScan(&user)
		if err != nil {
			return err
		}
		*users = append(*users, user)
	}

	return nil
}

func (repo UserRepositorySQLServerSQLX) Get(ctx context.Context, user *entity.User, id uuid.UUID) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) WHERE ID = ?"

	err = repo.db.Get(user, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQLX) Create(ctx context.Context, user *entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// query := fmt.Sprintf("INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES ('%s','%s','%s');", user.ID, user.Nome, user.Status)
	// _, err = repo.db.Exec(query)

	query := "INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES ($1, $2, $3)"

	res := repo.db.MustExec(query, user.ID, user.Nome, user.Status)

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQLX) CreateInBatch(ctx context.Context, users []entity.User) (err error) {

	tx := repo.db.MustBegin()
	// tx, err := repo.db.Begin()
	// if err != nil {
	// 	return err
	// }

	query := "INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES ($1, $2, $3)"

	for _, user := range users {
		res := repo.db.MustExec(query, user.ID, user.Nome, user.Status)
		_, err = res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo UserRepositorySQLServerSQLX) Update(ctx context.Context, user *entity.User) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	//query := fmt.Sprintf("UPDATE TBUSUARIO SET NOME = '%s', STATUS = '%s' WHERE ID = '%s';", user.Nome, user.Status, user.ID)
	//_, err = repo.db.Exec(query)

	query := "UPDATE TBUSUARIO SET NOME = $1, STATUS = $2 WHERE ID = $3"
	res := repo.db.MustExec(query, user.Nome, user.Status, user.ID)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositorySQLServerSQLX) Delete(ctx context.Context, id uuid.UUID) (err error) {

	if repo.db == nil {
		return errors.New("db is null")
	}

	err = repo.db.PingContext(ctx)
	if err != nil {
		return err
	}

	// query := fmt.Sprintf("DELETE FROM TBUSUARIO WHERE ID = '%s';", id)
	// _, err = repo.db.Exec(query)

	query := "DELETE FROM TBUSUARIO WHERE ID = $1"
	res := repo.db.MustExec(query, id)
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
