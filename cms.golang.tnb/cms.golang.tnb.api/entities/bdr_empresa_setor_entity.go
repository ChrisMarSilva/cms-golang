package entity

type BdrEmpresaSetor struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (BdrEmpresaSetor) TableName() string {
	return "TBBDR_EMPRESA_SETOR"
}

// func (row BdrEmpresaSetor) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
