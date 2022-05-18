package dto

type BdrEmpresaSetorRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSetorResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSetorLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
