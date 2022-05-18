package entity

type AcaoEmpresaSegmento struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AcaoEmpresaSegmento) TableName() string {
	return "TBEMPRESA_SEGMENTO"
}

// func (row AcaoEmpresaSegmento) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
