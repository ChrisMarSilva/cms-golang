package dto

type FiiEmpresaAdminRequest struct {
	ID       int64  `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	CNPJ     string `json:"cnpj,omitempty"`
	Situacao string `json:"situacao,omitempty"`
}

type FiiEmpresaAdminResponse struct {
	ID       int64  `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	CNPJ     string `json:"cnpj,omitempty"`
	Situacao string `json:"situacao,omitempty"`
}

type FiiEmpresaAdminLocal struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
