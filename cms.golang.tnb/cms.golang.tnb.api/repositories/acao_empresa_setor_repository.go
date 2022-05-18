package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AcaoEmpresaSetorRepository struct {
}

func NewAcaoEmpresaSetorRepository() *AcaoEmpresaSetorRepository {
	return &AcaoEmpresaSetorRepository{}
}

func (repo AcaoEmpresaSetorRepository) GetLista(db *gorm.DB, rows *[]entity.AcaoEmpresaSetor) (err error) {
	//	err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBEMPRESA_SETOR S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
