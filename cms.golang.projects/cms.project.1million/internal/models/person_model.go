package models

import (
	"time"

	"github.com/google/uuid"
)

type PersonModel struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func NewPersonModel(name string) *PersonModel {
	return &PersonModel{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}
