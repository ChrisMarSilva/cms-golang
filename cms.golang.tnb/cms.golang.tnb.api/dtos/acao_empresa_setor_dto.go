package dto

type AcaoEmpresaSetorRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSetorResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSetorLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
