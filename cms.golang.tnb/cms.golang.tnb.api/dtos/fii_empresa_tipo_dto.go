package dto

type FiiEmpresaTipoRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type FiiEmpresaTipoResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type FiiEmpresaTipoLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
