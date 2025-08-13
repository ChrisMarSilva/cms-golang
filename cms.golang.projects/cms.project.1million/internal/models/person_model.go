package models

import (
	"time"

	"github.com/google/uuid"
)

type PersonModel struct {
	ID          uuid.UUID
	Name        string
	RequestedAt time.Time
}

func NewPersonModel(name string) *PersonModel {
	return &PersonModel{
		ID:          uuid.New(),
		Name:        name,
		RequestedAt: time.Now().UTC(),
	}
}
