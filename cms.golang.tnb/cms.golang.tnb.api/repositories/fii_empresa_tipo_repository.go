package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type FiiEmpresaTipoRepository struct {
}

func NewFiiEmpresaTipoRepository() *FiiEmpresaTipoRepository {
	return &FiiEmpresaTipoRepository{}
}

func (repo FiiEmpresaTipoRepository) GetLista(db *gorm.DB, rows *[]entity.FiiEmpresaTipo) (err error) {
	//err := s.bd.Raw("SELECT ID AS ID, DESCRICAO AS Descricao FROM TBFII_FUNDOIMOB_TIPO ORDER BY DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
