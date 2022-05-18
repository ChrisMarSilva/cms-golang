package entity

type CorretoraLista struct {
	ID           int64  `json:"id,omitempty" gorm:"column:ID"`
	Nome         string `json:"nome,omitempty" gorm:"column:NOME"`
	CNPJ         string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	ImportatNota string `json:"importatnota,omitempty" gorm:"column:IMPORTAR_NOTA"`
	Situacao     string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (CorretoraLista) TableName() string {
	return "TBCORRETORA_LISTA"
}

// func (row CorretoraLista) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
