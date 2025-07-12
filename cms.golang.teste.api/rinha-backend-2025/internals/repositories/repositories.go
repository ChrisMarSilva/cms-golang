package repositories

import (
	"context"

	"github.com/chrismarsilva/rinha-backend-2025/internals/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	User interface {
		GetAll(ctx context.Context) ([]models.User, error)
		GetByID(ctx context.Context, id uuid.UUID) (models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		Create(ctx context.Context, model models.User) error
		//Update(ctx context.Context, model models.User) error
		//Delete(ctx context.Context, id uuid.UUID) error
	}
}

func New(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: NewUsersRepository(db),
	}
}
