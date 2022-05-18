package dto

type BdrEmpresaSegmentoRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSegmentoResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type BdrEmpresaSegmentoLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
