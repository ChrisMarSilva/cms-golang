package models

// type UsuarioCarteiraCripto struct {
// 	gorm.Model
// 	ID            int      `gorm:"primary_key;auto_increment"`
// 	IDCarteira    int      `gorm:"type:int;not null;index"`
// 	IDCripto      int      `gorm:"type:int;not null;index"`
// 	Quant         *int     `gorm:"type:int"`
// 	VLRPrecoMedio *float64 `gorm:"type:decimal(20,10)"`
// 	VLRPrecoTeto  *float64 `gorm:"type:decimal(17,2)"`
// 	PercentBalac  *float64 `gorm:"type:decimal(17,2)"`
// 	NotaBalac     *int     `gorm:"type:int"`
// 	Situacao      string   `gorm:"type:varchar(1);not null;index"`
// }

// func (u *UsuarioCarteiraCripto) SituacaoDescr() string {
// 	switch u.Situacao {
// 	case "A":
// 		return "Ativo"
// 	case "I":
// 		return "Inativo"
// 	case "F":
// 		return "Finalizado"
// 	default:
// 		return "Desconhecido"
// 	}
// }
