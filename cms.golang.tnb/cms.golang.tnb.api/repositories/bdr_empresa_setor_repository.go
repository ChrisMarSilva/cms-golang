package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type BdrEmpresaSetorRepository struct {
}

func NewBdrEmpresaSetorRepository() *BdrEmpresaSetorRepository {
	return &BdrEmpresaSetorRepository{}
}

func (repo BdrEmpresaSetorRepository) GetLista(db *gorm.DB, rows *[]entity.BdrEmpresaSetor) (err error) {
	//err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBBDR_EMPRESA_SETOR S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
