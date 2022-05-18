package entity

type BdrEmpresaSegmento struct {
	ID        int64  `json:"id,omitempty" gorm:"column:ID"`
	Descricao string `json:"descricao,omitempty" gorm:"column:DESCRICAO"`
	Situacao  string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (BdrEmpresaSegmento) TableName() string {
	return "TBBDR_EMPRESA_SEGMENTO"
}

// func (row BdrEmpresaSegmento) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
