package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type LoginRepository struct {
}

func NewLoginRepository() *LoginRepository {
	return &LoginRepository{}
}

func (repo LoginRepository) GetEntrar(db *gorm.DB, row *entity.Usuario) (err error) {
	// err := s.bd.Raw("SELECT ID, NOME, EMAIL, SENHA, DTREGISTRO, TENTATIVA, FOTO, CHATID, TIPO, SITUACAO FROM TBUSUARIO WHERE EMAIL = ? OR EMAIL LIKE ?  ", txtEmail, txtEmail+"@%").Scan(&user).Error
	err = db.Find(row).Error
	if err != nil {
		return err
	}
	return nil
}
