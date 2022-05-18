package dto

type UsuarioRequest struct {
}

type UsuarioResponse struct {
}

type UsuarioLocal struct {
	ID         int64  `json:"id" gorm:"column:ID"`
	Nome       string `json:"nome" gorm:"column:NOME"`
	Email      string `json:"email" gorm:"column:EMAIL"`
	Senha      string `json:"senha" gorm:"column:SENHA"`
	DtRegistro string `json:"data_registro" gorm:"column:DTREGISTRO"`
	Tentatia   int32  `json:"tentatia" gorm:"column:TENTATIVA"`
	Foto       string `json:"foto" gorm:"column:FOTO"`
	ChatId     string `json:"chat_id" gorm:"column:CHATID"`
	Tipo       string `json:"tipo" gorm:"column:TIPO"`
	Situacao   string `json:"situacao" gorm:"column:SITUACAO"`
}
