package entity

type AcaoEmpresa struct {
	ID           int64  `json:"id,omitempty" gorm:"column:ID"`
	IDSetor      int64  `json:"idsetor,omitempty" gorm:"column:IDSETOR"`
	IDSubSetor   int64  `json:"idsubsetor,omitempty" gorm:"column:IDSUBSETOR"`
	IDSegmento   int64  `json:"idsegmento,omitempty" gorm:"column:IDSEGMENTO"`
	Nome         string `json:"nome,omitempty" gorm:"column:NOME"`
	NomeResumido string `json:"nomeresumido,omitempty" gorm:"column:NOMRESUMIDO"`
	RazaoSocial  string `json:"razaosocial,omitempty" gorm:"column:RAZAOSOCIAL"`
	CNPJ         string `json:"cnpj,omitempty" gorm:"column:CNPJ"`
	Atividade    string `json:"atividade,omitempty" gorm:"column:ATIVIDADE"`
	CodigoCVM    string `json:"codigocvm,omitempty" gorm:"column:CODCVM"`
	SituacaoCVM  string `json:"situacaocvm,omitempty" gorm:"column:SITCVM"`
	Slug         string `json:"slug,omitempty" gorm:"column:SLUG"`
	Site         string `json:"site,omitempty" gorm:"column:SITE"`
	TipoMercado  string `json:"tipomercado,omitempty" gorm:"column:TIPO_MERCADO"`
	Situacao     string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AcaoEmpresa) TableName() string {
	return "TBEMPRESA"
}

// func (row CriptoEmpresaSituacao) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
