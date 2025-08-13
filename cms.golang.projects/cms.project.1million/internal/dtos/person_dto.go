package dtos

import (
	"time"

	"github.com/google/uuid"
)

type PersonRequestDto struct {
	Name string `json:"name"`
}

type PersonResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	RequestedAt time.Time `json:"requested_at"`
}

type PersonResponseBasicDto struct {
	Name string `json:"name"`
}
