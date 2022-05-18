package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type BdrEmpresaSubSetorRepository struct {
}

func NewBdrEmpresaSubSetorRepository() *BdrEmpresaSubSetorRepository {
	return &BdrEmpresaSubSetorRepository{}
}

func (repo BdrEmpresaSubSetorRepository) GetLista(db *gorm.DB, rows *[]entity.BdrEmpresaSubSetor) (err error) {
	//err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBBDR_EMPRESA_SUBSETOR S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
