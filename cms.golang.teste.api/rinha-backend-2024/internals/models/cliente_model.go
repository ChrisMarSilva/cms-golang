package models

import (
	"fmt"
)

type Cliente struct {
	Limite int64 `db:"limite"`
	Saldo  int64 `db:"saldo"`
}

func (row Cliente) ToString() string {
	return fmt.Sprintf("limite: %s; saldo: %s", row.Limite, row.Saldo)
}

func NewCliente(limite int64, saldo int64) *Cliente {
	return &Cliente{
		Limite: limite,
		Saldo:  saldo,
	}
}
