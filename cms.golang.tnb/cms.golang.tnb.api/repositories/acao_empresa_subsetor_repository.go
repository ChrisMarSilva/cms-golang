package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AcaoEmpresaSubSetorRepository struct {
}

func NewAcaoEmpresaSubSetorRepository() *AcaoEmpresaSubSetorRepository {
	return &AcaoEmpresaSubSetorRepository{}
}

func (repo AcaoEmpresaSubSetorRepository) GetLista(db *gorm.DB, rows *[]entity.AcaoEmpresaSubSetor) (err error) {
	// 	err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBEMPRESA_SUBSETOR S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
