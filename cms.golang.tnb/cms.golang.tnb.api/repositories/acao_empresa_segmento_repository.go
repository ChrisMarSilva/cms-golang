package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AcaoEmpresaSegmentoRepository struct {
}

func NewAcaoEmpresaSegmentoRepository() *AcaoEmpresaSegmentoRepository {
	return &AcaoEmpresaSegmentoRepository{}
}

func (repo AcaoEmpresaSegmentoRepository) GetLista(db *gorm.DB, rows *[]entity.AcaoEmpresaSegmento) (err error) {
	// err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBEMPRESA_SEGMENTO S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
