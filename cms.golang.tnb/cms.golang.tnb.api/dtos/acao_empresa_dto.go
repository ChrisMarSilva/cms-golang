package dto

type AcaoEmpresaRequest struct {
	ID           int64  `json:"id,omitempty"`
	IDSetor      int64  `json:"idsetor,omitempty"`
	IDSubSetor   int64  `json:"idsubsetor,omitempty"`
	IDSegmento   int64  `json:"idsegmento,omitempty"`
	Nome         string `json:"nome,omitempty"`
	NomeResumido string `json:"nomeresumido,omitempty"`
	RazaoSocial  string `json:"razaosocial,omitempty"`
	CNPJ         string `json:"cnpj,omitempty"`
	Atividade    string `json:"atividade,omitempty"`
	CodigoCVM    string `json:"codigocvm,omitempty"`
	SituacaoCVM  string `json:"situacaocvm,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Site         string `json:"site,omitempty"`
	TipoMercado  string `json:"tipomercado,omitempty"`
	Situacao     string `json:"situacao,omitempty"`
}

type AcaoEmpresaResponse struct {
	ID           int64  `json:"id,omitempty"`
	IDSetor      int64  `json:"idsetor,omitempty"`
	IDSubSetor   int64  `json:"idsubsetor,omitempty"`
	IDSegmento   int64  `json:"idsegmento,omitempty"`
	Nome         string `json:"nome,omitempty"`
	NomeResumido string `json:"nomeresumido,omitempty"`
	RazaoSocial  string `json:"razaosocial,omitempty"`
	CNPJ         string `json:"cnpj,omitempty"`
	Atividade    string `json:"atividade,omitempty"`
	CodigoCVM    string `json:"codigocvm,omitempty"`
	SituacaoCVM  string `json:"situacaocvm,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Site         string `json:"site,omitempty"`
	TipoMercado  string `json:"tipomercado,omitempty"`
	Situacao     string `json:"situacao,omitempty"`
}

type AcaoEmpresaLocal struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
