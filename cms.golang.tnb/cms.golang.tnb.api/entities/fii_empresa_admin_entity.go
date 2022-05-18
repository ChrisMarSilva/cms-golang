package entity

type FiiEmpresaAdmin struct {
	ID       int64  `json:"id,omitempty" gorm:"column:ID"`
	Nome     string `json:"nome,omitempty" gorm:"column:NOME"`
	CNPJ     string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	Situacao string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (FiiEmpresaAdmin) TableName() string {
	return "TBFII_FUNDOIMOB_ADMIN"
}

// func (row FiiEmpresaAdmin) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
