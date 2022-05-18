package dto

type CriptoEmpresaSituacaoRequest struct {
	Codigo    string `json:"codigo,omitempty"`
	Descricao string `json:"descricao,omitempty"`
}

type CriptoEmpresaSituacaoResponse struct {
	Codigo    string `json:"codigo,omitempty"`
	Descricao string `json:"descricao,omitempty"`
}
