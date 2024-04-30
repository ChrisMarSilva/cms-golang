package repositories

import (
	"context"
	"database/sql"
	"log"
	"time"

	data "github.com/chrismarsilva/cms.golang.tnb.cripo.database"
	"github.com/chrismarsilva/cms.golang.tnb.cripo.domain/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*models.UserModel, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.UserModel, error)
	GetAll(ctx context.Context, tx *sql.Tx) (*[]models.UserModel, error)

	Create(ctx context.Context, tx *sql.Tx, user *models.UserModel) error
	Update(ctx context.Context, tx *sql.Tx, user *models.UserModel) error
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID)
}

type UserRepository struct {
	db *data.Database
}

func NewUserRepository(db *data.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (this UserRepository) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, id)
	} else {
		row = this.db.QueryRowContext(timeoutCtx, query, id)
	}

	user := &models.UserModel{} // var user models.User
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}

func (this UserRepository) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE email = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, email)
	} else {
		row = this.db.QueryRowContext(timeoutCtx, query, email)
	}

	user := &models.UserModel{}
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}


func (this UserRepository) GetAll(ctx context.Context, tx *sql.Tx) (map[int]models.UserModel, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var rows *sql.Rows
	var err error
	query := `SELECT * FROM users`

	if tx != nil {
		rows, err = tx.QueryContext(timeoutCtx, query)
	} else {
		rows, err = this.db.QueryContext(timeoutCtx, query)
	}
	if err != nil {
		return nil, err
	}

	//var users = make([]models.UserModel, 0)
	users := make(map[int]models.UserModel)

	for rows.Next() {
		var user models.UserModel

		err := rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
		if err != nil {
			log.Println("Erro no Scan:", err.Error())
			return nil, err
		}

		//users = append(users, user)
		users[len(users)] = user
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return users, nil
}

func (this UserRepository) Create(ctx context.Context, tx *sql.Tx, user *models.UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `INSERT INTO users (id, nome, email, password, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, user.ID, user.Nome, user.Email, user.Password, user.IsActive, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (this UserRepository) Update(ctx context.Context, tx *sql.Tx, user *models.UserModel) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `UPDATE users SET nome = ?, password = ?, is_active = ? WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, user.Nome, user.Password, user.IsActive, user.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (this UserRepository) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `DELETE FROM users WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = this.db.Prepare(query)
	}

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(timeoutCtx, id) // .String()
	if err != nil {
		return err
	}

	return nil
}
