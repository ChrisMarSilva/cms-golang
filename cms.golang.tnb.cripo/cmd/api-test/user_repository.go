package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*UserEntity, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*UserEntity, error)
	GetAll(ctx context.Context, tx *sql.Tx) (map[int]UserEntity, error) // *[]UserEntity

	Create(ctx context.Context, tx *sql.Tx, user *UserEntity) error
	Update(ctx context.Context, tx *sql.Tx, user *UserEntity) error
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID)
}

type defaultRepository struct {
	db *Database //  Db *sql.DB
}

func NewUserRepository(db *Database) *defaultRepository {
	return &defaultRepository{
		db: db,
	}
}

func (repo defaultRepository) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*UserEntity, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, id)
	} else {
		row = repo.db.QueryRowContext(timeoutCtx, query, id)
	}

	user := &UserEntity{}
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	return user, nil
}

func (repo defaultRepository) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*UserEntity, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var row *sql.Row
	query := `SELECT * FROM users WHERE email = ?`

	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, email)
	} else {
		row = repo.db.QueryRowContext(timeoutCtx, query, email)
	}

	user := &UserEntity{}

	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		log.Println("Erro no Scan:", err.Error())
		return nil, err
	}

	return user, nil
}

func (repo defaultRepository) GetAll(ctx context.Context, tx *sql.Tx) (map[int]UserEntity, error) { // []UserEntity
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var rows *sql.Rows
	var err error
	query := `SELECT * FROM users`

	if tx != nil {
		rows, err = tx.QueryContext(timeoutCtx, query)
	} else {
		rows, err = repo.db.QueryContext(timeoutCtx, query)
	}
	if err != nil {
		return nil, err
	}

	users := make(map[int]UserEntity) //var users = make([]UserEntity, 0)

	for rows.Next() {
		var user UserEntity
		err := rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
		if err != nil {
			log.Println("Erro no Scan:", err.Error())
			return nil, err
		}
		users[len(users)] = user //users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo defaultRepository) Create(ctx context.Context, tx *sql.Tx, user *UserEntity) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `INSERT INTO users (id, nome, email, password, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = repo.db.Prepare(query)
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

func (repo defaultRepository) Update(ctx context.Context, tx *sql.Tx, user *UserEntity) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `UPDATE users SET nome = ?, password = ?, is_active = ? WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = repo.db.Prepare(query)
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

func (repo defaultRepository) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var stmt *sql.Stmt
	var err error
	query := `DELETE FROM users WHERE id = ?`

	if tx != nil {
		stmt, err = tx.Prepare(query)
	} else {
		stmt, err = repo.db.Prepare(query)
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
