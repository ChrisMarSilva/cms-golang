package entities

import (
	"fmt"
)

type CriptoEmpresaSituacao struct {
	Codigo    string `json:"codigo" gorm:"primaryKey; column:CODIGO"`
	Descricao string `json:"descricao" gorm:"column:DESCRICAO"`
}

func (CriptoEmpresaSituacao) TableName() string {
	return "TBCRIPTO_EMPRESA_ST"
}

func (row CriptoEmpresaSituacao) ToString() string {
	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
}

// type CriptoEmpresaSituacao struct {
// 	Codigo    string `json:"codigo"`
// 	Descricao string `json:"descricao"`
// }
