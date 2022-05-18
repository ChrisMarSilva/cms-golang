package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type BdrEmpresaRepository struct {
}

func NewBdrEmpresaRepository() *BdrEmpresaRepository {
	return &BdrEmpresaRepository{}
}

func (repo BdrEmpresaRepository) GetLista(db *gorm.DB, rows *[]entity.BdrEmpresa) (err error) {
	//err := s.bd.Raw("SELECT E.ID AS ID, E.NOME AS Nome FROM TBBDR_EMPRESA E WHERE E.SITUACAO = 'A' ORDER BY E.RAZAOSOCIAL").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
