package repository

import (
	"context"

	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetAll(ctx context.Context, users *[]entity.User) error
	Get(ctx context.Context, user *entity.User, id uuid.UUID) error
	Create(ctx context.Context, user *entity.User) error
	CreateInBatch(ctx context.Context, users []entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
