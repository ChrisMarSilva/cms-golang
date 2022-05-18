package entity

type FiiEmpresa struct {
	ID          int64  `json:"id,omitempty" gorm:"column:ID"`
	IDTipo      int64  `json:"idtipo,omitempty" gorm:"column:IDFIITIPO"`
	IDAdmin     int64  `json:"idadmin,omitempty" gorm:"column:IDFIIADMIN"`
	Nome        string `json:"nome,omitempty" gorm:"column:NOME"`
	RazaoSocial string `json:"razaosocial,omitempty" gorm:"column:RAZAOSOCIAL"`
	CNPJ        string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	Codigo      string `json:"codigo,omitempty" gorm:"column:CODIGO"`
	CodigoIsin  string `json:"codigoisin,omitempty" gorm:"column:CODISIN"`
	Situacao    string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (FiiEmpresa) TableName() string {
	return "TBFII_FUNDOIMOB"
}

// func (row FiiEmpresa) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
