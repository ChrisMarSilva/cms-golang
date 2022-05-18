package entity

type AlertaAssinatura struct {
	ID             int64  `json:"id,omitempty" gorm:"column:ID"`
	IDUsuario      int64  `json:"idusuario,omitempty" gorm:"column:IDUSUARIO"`
	DtHrRegistro   string `json:"dthr_registro,omitempty" gorm:"column:DTHRREGISTRO"`
	DtHrAlteracao  string `json:"dthr_alteracao,omitempty" gorm:"column:DTHRALTERACAO"`
	TipoAlerta     string `json:"tipoalerta,omitempty" gorm:"column:TIPO_ALERTA"`
	TipoAssinatura string `json:"tipoassinatura,omitempty" gorm:"column:TIPO_ASSINATURA"`
	Situacao       string `json:"situacao,omitempty" gorm:"column:SITUACAO"`
}

func (AlertaAssinatura) TableName() string {
	return "TBALERTA_ASSINATURA"
}

// func (row AlertaAssinatura) ToString() string {
// 	return fmt.Sprintf("codigo: %s; descricao: %s", row.Codigo, row.Descricao)
// }
