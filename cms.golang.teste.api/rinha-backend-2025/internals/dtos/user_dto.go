package dtos

import (
	"github.com/google/uuid"
)

type UserRequestDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseDto struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
