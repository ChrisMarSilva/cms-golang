package entities

import (
	"github.com/google/uuid"
)

type UsersRoleType string

const (
	AdminRole UsersRoleType = "admin"
	UserRole  UsersRoleType = "user"
	GuestRole UsersRoleType = "guest"
)

type UserRoleEntity struct {
	ID        uuid.UUID `db:"id"`
	Nome      string    `db:"nome"`
	Descricao string    `db:"descricao"`
}

func NewUserRoleEntity(ID uuid.UUID, nome, descricao string) *UserRoleEntity {
	return &UserRoleEntity{
		ID:        ID,
		Nome:      nome,
		Descricao: descricao,
	}
}
