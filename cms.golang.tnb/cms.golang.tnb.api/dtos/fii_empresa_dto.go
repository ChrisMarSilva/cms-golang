package dto

type FiiEmpresaRequest struct {
	ID          int64  `json:"id,omitempty"`
	IDTipo      int64  `json:"idtipo,omitempty"`
	IDAdmin     int64  `json:"idadmin,omitempty"`
	Nome        string `json:"nome,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type FiiEmpresaResponse struct {
	ID          int64  `json:"id,omitempty"`
	IDTipo      int64  `json:"idtipo,omitempty"`
	IDAdmin     int64  `json:"idadmin,omitempty"`
	Nome        string `json:"nome,omitempty"`
	RazaoSocial string `json:"razaosocial,omitempty"`
	CNPJ        string `json:"cnpj,omitempty"`
	Codigo      string `json:"codigo,omitempty"`
	CodigoIsin  string `json:"codigoisin,omitempty"`
	Situacao    string `json:"situacao,omitempty"`
}

type FiiEmpresaLocal struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Nomee string `json:"nomee"`
}
