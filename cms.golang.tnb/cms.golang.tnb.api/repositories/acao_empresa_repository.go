package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AcaoEmpresaRepository struct {
}

func NewAcaoEmpresaRepository() *AcaoEmpresaRepository {
	return &AcaoEmpresaRepository{}
}

func (repo AcaoEmpresaRepository) GetLista(db *gorm.DB, rows *[]entity.AcaoEmpresa) (err error) {
	// 	err := s.bd.Raw("SELECT E.ID AS ID, E.NOMRESUMIDO AS Nome FROM TBEMPRESA E WHERE E.SITUACAO = 'A' ORDER BY E.RAZAOSOCIAL").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
