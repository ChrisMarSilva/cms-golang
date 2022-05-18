package entity

type EtfEmpresa struct {
	ID          int64  `json:"id,omitempty" gorm:"column:ID"`
	RazaoSocial string `json:"razaosocial,omitempty" gorm:"column:RAZAOSOCIAL"`
	Fundo       string `json:"fundo,omitempty" gorm:"column:FUNDO"`
	Indice      string `json:"indice,omitempty" gorm:"column:INDICE"`
	Nome        string `json:"nome,omitempty" gorm:"column:NOME"`
	CNPJ        string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	Codigo      string `json:"codigo,omitempty" gorm:"column:CODIGO"`
	CodigoIsin  string `json:"codigoisin,omitempty" gorm:"column:CODISIN"`
	Situacao    string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (EtfEmpresa) TableName() string {
	return "TBETF_INDICE"
}

// func (row EtfEmpresa) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
