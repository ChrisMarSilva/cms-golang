package repository

// import (
// 	"context"
// 	"errors"

// 	database "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/databases"
// 	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
// 	_ "github.com/denisenkom/go-mssqldb"
// 	"github.com/google/uuid"
// )

// type UserRepositorySQLServerSQLC struct {
// 	db *database.Dbtx
// }

// func NewUserRepositorySQLServerSQLC(db *database.Dbtx) *UserRepositorySQLServerSQLC {
// 	return &UserRepositorySQLServerSQLC{db: db}
// }

// // func (repo *UserRepositorySQLServerSQLC) NewUserRepositorySQLServerSQLCWithTx(tx *sql.Tx) *UserRepositorySQLServerSQLC {
// // 	return &UserRepositorySQLServerSQLC{db: *tx}
// // }

// func (repo UserRepositorySQLServerSQLC) GetAll(users *[]entity.User) (err error) {

// 	if repo.db == nil {
// 		return errors.New("db is null")
// 	}

// 	ctx := context.Background()

// 	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) ORDER BY ID"

// 	rows, err := repo.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var user entity.User
// 		if err := rows.Scan(&user.ID, &user.Nome, &user.Status); err != nil {
// 			return err
// 		}
// 		*users = append(*users, user)
// 	}

// 	return nil
// }

// func (repo UserRepositorySQLServerSQLC) Get(user *entity.User, id uuid.UUID) (err error) {

// 	if repo.db == nil {
// 		return errors.New("db is null")
// 	}

// 	ctx := context.Background()

// 	query := "SELECT ID, NOME, STATUS FROM TBUSUARIO WITH(NOLOCK) WHERE ID = ?"

// 	row := repo.db.QueryRowContext(ctx, query, id)

// 	err = row.Scan(&user.ID, &user.Nome, &user.Status)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repo UserRepositorySQLServerSQLC) Create(user *entity.User) (err error) {

// 	if repo.db == nil {
// 		return errors.New("db is null")
// 	}

// 	ctx := context.Background()

// 	query := "INSERT INTO TBUSUARIO (ID, NOME, STATUS) VALUES (?,?,?)"

// 	_, err = repo.db.ExecContext(ctx, query, user.ID, user.Nome, user.Status)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repo UserRepositorySQLServerSQLC) Update(user *entity.User) (err error) {

// 	if repo.db == nil {
// 		return errors.New("db is null")
// 	}

// 	ctx := context.Background()

// 	query := "UPDATE TBUSUARIO SET NOME = ?, STATUS = ? WHERE ID = ?"

// 	_, err = repo.db.ExecContext(ctx, query, user.Nome, user.Status, user.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repo UserRepositorySQLServerSQLC) Delete(id uuid.UUID) (err error) {

// 	if repo.db == nil {
// 		return errors.New("db is null")
// 	}

// 	ctx := context.Background()

// 	query := "DELETE FROM TBUSUARIO WHERE ID = ?"

// 	_, err = repo.db.ExecContext(ctx, query, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
