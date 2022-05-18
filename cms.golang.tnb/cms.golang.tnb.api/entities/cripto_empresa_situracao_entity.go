package entity

import (
	"fmt"
)

type CriptoEmpresaSituacao struct {
	Codigo    string `json:"codigo,omitempty" gorm:"column:CODIGO"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
}

func (CriptoEmpresaSituacao) TableName() string {
	return "TBCRIPTO_EMPRESA_ST"
}

func (row CriptoEmpresaSituacao) ToString() string {
	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
}
