package dto

type BdrEmpresaSubSetorRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSubSetorResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSubSetorLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
