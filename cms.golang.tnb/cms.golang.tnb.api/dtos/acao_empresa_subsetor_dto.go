package dto

type AcaoEmpresaSubSetorRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSubSetorResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSubSetorLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
