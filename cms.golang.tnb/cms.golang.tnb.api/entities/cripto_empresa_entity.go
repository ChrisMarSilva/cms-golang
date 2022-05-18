package entity

type CriptoEmpresa struct {
	ID                 int64   `json:"id,omitempty" gorm:"column:ID"`
	Nome               string  `json:"nome,omitempty" gorm:"column:NOME"`
	Codigo             string  `json:"codigo,omitempty" gorm:"column:CODIGO"`
	Fechamento         float64 `json:"fechamento,omitempty" gorm:"column:VLRPRECOFECHAMENTO"`
	FechamentoAnterior float64 `json:"fechamentoanterior,omitempty" gorm:"column:VLRPRECOANTERIOR"`
	Variacao           float64 `json:"variacao,omitempty" gorm:"column:VLRVARIACAO"`
	DataHoraAlteracao  string  `json:"datahora,omitempty" gorm:"column:DATAHORAALTERACO"`
	Situacao           string  `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (CriptoEmpresa) TableName() string {
	return "TBCRIPTO_EMPRESA"
}

// func (row CriptoEmpresa) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
