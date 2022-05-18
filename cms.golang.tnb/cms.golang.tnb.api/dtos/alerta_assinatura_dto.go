package dto

type AlertaAssinaturaRequest struct {
	ID             int64  `json:"id,omitempty"`
	IDUsuario      int64  `json:"idusuario,omitempty"`
	DtHrRegistro   string `json:"dthr_registro,omitempty"`
	DtHrAlteracao  string `json:"dthr_alteracao,omitempty"`
	TipoAlerta     string `json:"tipoalerta,omitempty"`
	TipoAssinatura string `json:"tipoassinatura,omitempty"`
	Situacao       string `json:"situacao,omitempty"`
}

type AlertaAssinaturaResponse struct {
	ID             int64  `json:"id,omitempty"`
	IDUsuario      int64  `json:"idusuario,omitempty"`
	DtHrRegistro   string `json:"dthr_registro,omitempty"`
	DtHrAlteracao  string `json:"dthr_alteracao,omitempty"`
	TipoAlerta     string `json:"tipoalerta,omitempty"`
	TipoAssinatura string `json:"tipoassinatura,omitempty"`
	Situacao       string `json:"situacao,omitempty"`
}

type AlertaAssinaturaLocal struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Nomee string `json:"nomee"`
}
