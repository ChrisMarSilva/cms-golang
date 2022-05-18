package dto

type CorretoraListaRequest struct {
	ID           int64  `json:"id,omitempty"`
	Nome         string `json:"nome,omitempty"`
	CNPJ         string `json:"cnpj,omitempty"`
	ImportatNota string `json:"importatnota,omitempty"`
	Situacao     string `json:"situacao,omitempty"`
}

type CorretoraListaResponse struct {
	ID           int64  `json:"id,omitempty"`
	Nome         string `json:"nome,omitempty"`
	CNPJ         string `json:"cnpj,omitempty"`
	ImportatNota string `json:"importatnota,omitempty"`
	Situacao     string `json:"situacao,omitempty"`
}

type CorretoraListaLocal struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
