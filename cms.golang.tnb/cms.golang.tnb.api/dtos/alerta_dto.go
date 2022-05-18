package dto

type AlertaRequest struct {
	ID               int64  `json:"id,omitempty"`
	IDUsuario        int64  `json:"idusuario,omitempty"`
	DtHrRegistro     string `json:"dthr_registro,omitempty"`
	DataEnvio        string `json:"data_envio,omitempty"`
	Tipo             string `json:"tipo,omitempty"`
	Mensagem         string `json:"mensagem,omitempty"`
	QtdProc          int64  `json:"qtdproc,omitempty"`
	SituacaoTelegram string `json:"situacaotelegram,omitempty"`
	SituacaoEmai     string `json:"situacaoemail,omitempty"`
}

type AlertaResponse struct {
	ID               int64  `json:"id,omitempty"`
	IDUsuario        int64  `json:"idusuario,omitempty"`
	DtHrRegistro     string `json:"dthr_registro,omitempty"`
	DataEnvio        string `json:"data_envio,omitempty"`
	Tipo             string `json:"tipo,omitempty"`
	Mensagem         string `json:"mensagem,omitempty"`
	QtdProc          int64  `json:"qtdproc,omitempty"`
	SituacaoTelegram string `json:"situacaotelegram,omitempty"`
	SituacaoEmai     string `json:"situacaoemail,omitempty"`
}

type AlertaLocal struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Nomee string `json:"nomee"`
}
