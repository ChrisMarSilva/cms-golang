package main

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) GetByEmail(ctx context.Context, email string) (*UserModel, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := repo.db.QueryRowContext(ctx, query, email)

	user := &UserModel{}
	err := row.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	log.Error("Erro no ErrNoRows:", err.Error())
		// 	return nil, fmt.Errorf("No user found with Email '%s'", email)
		// }
		log.Error("Erro no Scan:", err.Error())
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return user, nil
}

func (repo UserRepository) Create(ctx context.Context, user *UserModel) error {
	query := "INSERT INTO users (id, nome, email, password, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Nome, user.Email, user.Password, user.IsActive, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Update(ctx context.Context, user *UserModel) error {
	query := "UPDATE users SET nome = ?, password = ?, is_active = ? WHERE id = ?"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Nome, user.Password, user.IsActive, user.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id) // .String()
	if err != nil {
		return err
	}

	return nil
}
