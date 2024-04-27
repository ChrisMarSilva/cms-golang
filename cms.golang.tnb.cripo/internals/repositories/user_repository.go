package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/entities"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*entities.UserEntity, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*entities.UserEntity, error)
	GetAll(ctx context.Context, tx *sql.Tx) (*[]entities.UserEntity, error)

	Create(ctx context.Context, tx *sql.Tx, user *entities.UserEntity) error
	Update(ctx context.Context, tx *sql.Tx, user *entities.UserEntity) error
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID)
}

type defaultRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *defaultRepository {
	return &defaultRepository{
		db: db,
	}
}

func (repo defaultRepository) GetById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*entities.UserEntity, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE id = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, id)
	} else {
		row = repo.db.QueryRowContext(timeoutCtx, query, id)
	}

	user := &entities.UserEntity{} // var user models.User
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

func (repo defaultRepository) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*entities.UserEntity, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT * FROM users WHERE email = ?`

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(timeoutCtx, query, email)
	} else {
		row = repo.db.QueryRowContext(timeoutCtx, query, email)
	}

	user := &entities.UserEntity{}
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

func (repo defaultRepository) GetAll(ctx context.Context, tx *sql.Tx) (*[]entities.UserEntity, error) {
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

	//var users = make([]entities.UserEntity, 0)
	users := make(map[int]entities.UserEntity)

	for rows.Next() {
		var user entities.UserEntity

		err := rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Password, &user.IsActive, &user.CreatedAt)
		if err != nil {
			log.Error("Erro no Scan:", err.Error())
			return nil, err
		}

		//users = append(users, user)
		users[len(users)] = *user
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	//log.Info("ID:", user.ID, "Nome:", user.Nome, "Email:", user.Email, "Password:", user.Password, "IsActive:", user.IsActive, "Created_at:", user.Created_at)
	return users, nil
}

func (repo defaultRepository) Create(ctx context.Context, tx *sql.Tx, user *entities.UserEntity) error {
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

func (repo defaultRepository) Update(ctx context.Context, tx *sql.Tx, user *entities.UserEntity) error {
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
