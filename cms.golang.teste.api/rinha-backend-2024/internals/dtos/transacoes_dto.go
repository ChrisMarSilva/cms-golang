package dtos

import (
	"errors"
)

type TransacaoRequestDto struct {
	Valor     int64  `json:"valor,omitempty"`
	Tipo      string `json:"tipo,omitempty"`
	Descricao string `json:"descricao,omitempty"`
}

// func NewTransacaoRequestDto(valor int64, tipo string, descricao string) *TransacaoRequestDto {
// 	return &TransacaoRequestDto{
// 		Valor:     valor,
// 		Tipo:      tipo,
// 		Descricao: descricao,
// 	}
// }

func (request *TransacaoRequestDto) Valido() error {
	if request.Valor < 1 || request.Valor > 9223372036854775807 {
		return errors.New("Campo valor inválido.")
	}

	if request.Tipo == "" || len(request.Tipo) != 1 {
		return errors.New("Campo tipo inválido.")
	}

	if request.Tipo != "d" && request.Tipo != "c" {
		return errors.New("Campo tipo inválido.")
	}

	if request.Descricao == "" || len(request.Descricao) > 10 {
		return errors.New("Campo descrição inválido.")
	}

	return nil
}

type TransacaoResponseDto struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}
