package dto

type EtfEmpresaRequest struct {
	ID          int64  `json:"id,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	Fundo       string `json:"fundo,omitempty"`
	Indice      string `json:"indice,omitempty"`
	Nome        string `json:"nome,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type EtfEmpresaResponse struct {
	ID          int64  `json:"id,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	Fundo       string `json:"fundo,omitempty"`
	Indice      string `json:"indice,omitempty"`
	Nome        string `json:"nome,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type EtfEmpresaLocal struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Nomee string `json:"nomee"`
}
