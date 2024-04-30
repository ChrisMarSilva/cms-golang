package models

import (
	"github.com/google/uuid"
)

type UsersRoleType string

const (
	AdminRole UsersRoleType = "admin"
	UserRole  UsersRoleType = "user"
	GuestRole UsersRoleType = "guest"
)

type UserRoleModel struct {
	ID        uuid.UUID `db:"id"`
	Nome      string    `db:"nome"`
	Descricao string    `db:"descricao"`
}

func NewUserRoleModel(ID uuid.UUID, nome, descricao string) *UserRoleModel {
	return &UserRoleModel{
		ID:        ID,
		Nome:      nome,
		Descricao: descricao,
	}
}
