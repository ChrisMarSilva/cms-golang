package entity

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

const (
	UserActive   string = "A"
	UserInactive        = "I"
)

type User struct {
	ID     uuid.UUID `json:"id,omitempty" gorm:"column:ID" gorm:"type:uuid" bson:"_id,omitempty"`
	Nome   string    `json:"nome,omitempty" gorm:"column:NOME" bson:"nome,omitempty"`
	Status string    `json:"status,omitempty" gorm:"column:STATUS" bson:"status,omitempty"`
}

func (User) TableName() string {
	return "TBUSUARIO"
}

func (row User) ToString() string {
	return fmt.Sprintf("id: %s; nome: %s; status: %s", row.ID.String(), row.Nome, row.StatusString())
}

func (row User) StatusString() string {
	switch row.Status {
	case UserActive:
		return "Active"
	case UserInactive:
		return "Inactive"
	}
	return "unknown"
}

func (row User) Validate() error {
	if row.ID == uuid.Nil {
		return errors.New("id is NIL")
	}
	if row.ID.String() == "" {
		return errors.New("id is required")
	}
	if row.Nome == "" {
		return errors.New("nome is required")
	}
	if row.Status == "" {
		return errors.New("status is required")
	}
	return nil
}
