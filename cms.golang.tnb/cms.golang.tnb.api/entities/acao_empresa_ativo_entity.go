package entity

type AcaoEmpresaAtivo struct {
	ID         int64  `json:"id,omitempty" gorm:"column:ID"`
	IDEmpresa  int64  `json:"idempresa,omitempty" gorm:"column:IDEMPRESA"`
	Codigo     string `json:"codigo,omitempty" gorm:"column:CODIGO"`
	CodigoIsin string `json:"codigoisin,omitempty" gorm:"column:CODISIN"`
	Tipo       string `json:"tipo,omitempty" gorm:"column:TIPO"`
	Situacao   string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AcaoEmpresaAtivo) TableName() string {
	return "TBEMPRESA_ATIVO"
}

// func (row AcaoEmpresaAtivo) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
