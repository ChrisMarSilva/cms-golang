package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/chrismarsilva/rinha-backend-2025/internals/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (repo UsersRepository) GetAll(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)

	query := `SELECT "id", "name", "email" FROM "TbUser" ORDER BY "id" LIMIT 15`
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			slog.Error(err.Error())
			continue
			// return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das linhas: %w", err)
	}

	return users, nil
}

func (repo *UsersRepository) GetByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	var user models.User

	query := `SELECT id, name, email FROM "TbUser" WHERE id = $1 LIMIT 1`
	row := repo.db.QueryRow(ctx, query, id)

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found.")
		}
		return user, err
	}

	return user, nil
}

func (r *UsersRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := `SELECT "id", "name", "email" FROM "TbUser" WHERE "email" = $1`
	row := r.db.QueryRow(ctx, query, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows) {
			return nil, nil // errors.New("user not found.")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) Create(ctx context.Context, model models.User) error {
	query := `INSERT INTO "TbUser" (id, name, email) Values ($1, $2, $3)`

	result, err := r.db.Exec(ctx, query, model.ID, model.Name, model.Email)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		return errors.New("Usuario não inserido.")
	}

	return nil
}
