package dto

type AcaoEmpresaSegmentoRequest struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSegmentoResponse struct {
	ID        int64  `json:"id,omitempty"`
	Descricao string `json:"descricao,omitempty"`
	Situacao  string `json:"situacao,omitempty"`
}

type AcaoEmpresaSegmentoLocal struct {
	ID        int64  `json:"id"`
	Descricao string `json:"descricao"`
}
