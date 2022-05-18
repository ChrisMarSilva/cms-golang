package dto

type BdrEmpresaRequest struct {
	ID          int64  `json:"id,omitempty"`
	IDSetor     int64  `json:"idsetor,omitempty"`
	IDSubSetor  int64  `json:"idsubsetor,omitempty"`
	IDSegmento  int64  `json:"idsegmento,omitempty"`
	Nome        string `json:"nome,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Atividade   string `json:"atividade,omitempty"`
	CodigoCVM   string `json:"codigocvm,omitempty"`
	SituacaoCVM string `json:"situacaocvm,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Tipo        string `json:"tipo,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type BdrEmpresaResponse struct {
	ID          int64  `json:"id,omitempty"`
	IDSetor     int64  `json:"idsetor,omitempty"`
	IDSubSetor  int64  `json:"idsubsetor,omitempty"`
	IDSegmento  int64  `json:"idsegmento,omitempty"`
	Nome        string `json:"nome,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Atividade   string `json:"atividade,omitempty"`
	CodigoCVM   string `json:"codigocvm,omitempty"`
	SituacaoCVM string `json:"situacaocvm,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Tipo        string `json:"tipo,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type BdrEmpresaLocal struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
