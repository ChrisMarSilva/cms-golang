package entity

type AcaoEmpresaSubSetor struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AcaoEmpresaSubSetor) TableName() string {
	return "TBEMPRESA_SUBSETOR"
}

// func (row AcaoEmpresaSubSetor) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
