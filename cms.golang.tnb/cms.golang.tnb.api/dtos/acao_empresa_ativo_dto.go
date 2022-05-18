package dto

type AcaoEmpresaAtivoRequest struct {
	ID         int64  `json:"id,omitempty"`
	IDEmpresa  int64  `json:"idempresa,omitempty"`
	Codigo     string `json:"codigo,omitempty"`
	CodigoIsin string `json:"codigoisin,omitempty"`
	Tipo       string `json:"tipo,omitempty"`
	Situacao   string `json:"situacao,omitempty"`
}

type AcaoEmpresaAtivoResponse struct {
	ID         int64  `json:"id,omitempty"`
	IDEmpresa  int64  `json:"idempresa,omitempty"`
	Codigo     string `json:"codigo,omitempty"`
	CodigoIsin string `json:"codigoisin,omitempty"`
	Tipo       string `json:"tipo,omitempty"`
	Situacao   string `json:"situacao,omitempty"`
}
