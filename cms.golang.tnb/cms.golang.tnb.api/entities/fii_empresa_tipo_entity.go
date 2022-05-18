package entity

type FiiEmpresaTipo struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (FiiEmpresaTipo) TableName() string {
	return "TBFII_FUNDOIMOB_TIPO"
}

// func (row FiiEmpresaTipo) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
