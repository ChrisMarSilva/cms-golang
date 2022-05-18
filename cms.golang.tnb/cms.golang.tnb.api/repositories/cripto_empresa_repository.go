package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type CriptoEmpresaRepository struct {
}

func NewCriptoEmpresaRepository() *CriptoEmpresaRepository {
	return &CriptoEmpresaRepository{}
}

func (repo CriptoEmpresaRepository) GetLista(db *gorm.DB, rows *[]entity.CriptoEmpresa) (err error) {
	//err := s.bd.Raw("SELECT E.ID AS ID, E.NOME AS Nome FROM TBCRIPTO_EMPRESA E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
