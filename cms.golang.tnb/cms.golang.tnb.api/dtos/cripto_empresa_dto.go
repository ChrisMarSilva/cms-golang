package dto

type CriptoEmpresaRequest struct {
	ID                 int64   `json:"id,omitempty"`
	Nome               string  `json:"nome,omitempty"`
	Codigo             string  `json:"codigo,omitempty"`
	Fechamento         float64 `json:"fechamento,omitempty"`
	FechamentoAnterior float64 `json:"fechamentoanterior,omitempty"`
	Variacao           float64 `json:"variacao,omitempty"`
	DataHoraAlteracao  string  `json:"datahora,omitempty"`
	Situacao           string  `json:"situacao,omitempty"`
}

type CriptoEmpresaResponse struct {
	ID                 int64   `json:"id,omitempty"`
	Nome               string  `json:"nome,omitempty"`
	Codigo             string  `json:"codigo,omitempty"`
	Fechamento         float64 `json:"fechamento,omitempty"`
	FechamentoAnterior float64 `json:"fechamentoanterior,omitempty"`
	Variacao           float64 `json:"variacao,omitempty"`
	DataHoraAlteracao  string  `json:"datahora,omitempty"`
	Situacao           string  `json:"situacao,omitempty"`
}

type CriptoEmpresaLocal struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
