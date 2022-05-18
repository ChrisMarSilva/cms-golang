package entity

type AcaoEmpresaSetor struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AcaoEmpresaSetor) TableName() string {
	return "TBEMPRESA_SETOR"
}

// func (row AcaoEmpresaSetor) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
