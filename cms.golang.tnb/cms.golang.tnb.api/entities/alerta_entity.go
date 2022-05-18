package entity

type Alerta struct {
	ID               int64  `json:"id,omitempty" gorm:"column:ID"`
	IDUsuario        int64  `json:"idusuario,omitempty" gorm:"column:IDUSUARIO"`
	DtHrRegistro     string `json:"dthr_registro,omitempty" gorm:"column:DTHRREGISTRO"`
	DataEnvio        string `json:"data_envio,omitempty" gorm:"column:DTENVIO"`
	Tipo             string `json:"tipo,omitempty" gorm:"column:TIPO"`
	Mensagem         string `json:"mensagem,omitempty" gorm:"column:MENSAGEM"`
	QtdProc          int64  `json:"qtdproc,omitempty" gorm:"column:QTD_PROC"`
	SituacaoTelegram string `json:"situacaotelegram,omitempty" gorm:"column:SITUACAO_TELEGRAM"`
	SituacaoEmai     string `json:"situacaoemail,omitempty" gorm:"column:SITUACAO_EMAIL"`
}

func (Alerta) TableName() string {
	return "TBALERTA"
}

// func (row Alerta) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
