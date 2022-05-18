package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type FiiEmpresaRepository struct {
}

func NewFiiEmpresaRepository() *FiiEmpresaRepository {
	return &FiiEmpresaRepository{}
}

func (repo FiiEmpresaRepository) GetLista(db *gorm.DB, rows *[]entity.FiiEmpresa) (err error) {
	//err := s.bd.Raw("SELECT E.ID AS ID, E.NOME AS Nome, E.NOME AS Nomee FROM TBFII_FUNDOIMOB E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
