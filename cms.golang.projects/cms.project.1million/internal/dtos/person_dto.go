package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PersonRequestDto struct {
	Name string `json:"name"`
}

type PersonResponseDto struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPersonResponseDto(id uuid.UUID, name string, createdAt time.Time) *PersonResponseDto {
	return &PersonResponseDto{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}
