package users

import (
	"context"

	"github.com/ChrisMarSilva/cms.golang.teste.bd.sql.vs.orm/entities"
)

type UserRepository interface {
	Create(ctx context.Context, newUser entities.User) error
	GetUserByEmail(ctx context.Context, email string) ([]entities.User, error)
	GetByID(ctx context.Context, ID int64) (*entities.UserWithoutPassword, error)
	GetAll(ctx context.Context) ([]entities.UserWithoutPassword, error)
}
