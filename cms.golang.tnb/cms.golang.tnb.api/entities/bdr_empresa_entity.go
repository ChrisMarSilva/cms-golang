package entity

type BdrEmpresa struct {
	ID          int64  `json:"id,omitempty" gorm:"column:ID"`
	IDSetor     int64  `json:"idsetor,omitempty" gorm:"column:IDSETOR"`
	IDSubSetor  int64  `json:"idsubsetor,omitempty" gorm:"column:IDSUBSETOR"`
	IDSegmento  int64  `json:"idsegmento,omitempty" gorm:"column:IDSEGMENTO"`
	Nome        string `json:"nome,omitempty" gorm:"column:NOME"`
	RazaoSocial string `json:"razaosocial,omitempty" gorm:"column:RAZAOSOCIAL"`
	CNPJ        string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	Atividade   string `json:"atividade,omitempty" gorm:"column:ATIVIDADE"`
	CodigoCVM   string `json:"codigocvm,omitempty" gorm:"column:CODCVM"`
	SituacaoCVM string `json:"situacaocvm,omitempty" gorm:"column:SITCVM"`
	Codigo      string `json:"codigo,omitempty" gorm:"column:CODIGO"`
	CodigoIsin  string `json:"codigoisin,omitempty" gorm:"column:CODISIN"`
	Tipo        string `json:"tipo,omitempty" gorm:"column:TIPO"`
	Situacao    string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (BdrEmpresa) TableName() string {
	return "TBBDR_EMPRESA"
}

// func (row BdrEmpresa) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
