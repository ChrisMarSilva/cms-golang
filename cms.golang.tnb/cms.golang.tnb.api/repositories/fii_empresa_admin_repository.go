package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type FiiEmpresaAdminRepository struct {
}

func NewFiiEmpresaAdminRepository() *FiiEmpresaAdminRepository {
	return &FiiEmpresaAdminRepository{}
}

func (repo FiiEmpresaAdminRepository) GetLista(db *gorm.DB, rows *[]entity.FiiEmpresaAdmin) (err error) {
	//err := s.bd.Raw("SELECT ID AS ID, NOME AS Nome FROM TBFII_FUNDOIMOB_ADMIN ORDER BY NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
