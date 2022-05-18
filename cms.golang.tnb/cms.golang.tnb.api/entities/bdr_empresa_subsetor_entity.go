package entity

type BdrEmpresaSubSetor struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (BdrEmpresaSubSetor) TableName() string {
	return "TBBDR_EMPRESA_SUBSETOR"
}

// func (row BdrEmpresaSubSetor) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
