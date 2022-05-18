package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type BdrEmpresaSegmentoRepository struct {
}

func NewBdrEmpresaSegmentoRepository() *BdrEmpresaSegmentoRepository {
	return &BdrEmpresaSegmentoRepository{}
}

func (repo BdrEmpresaSegmentoRepository) GetLista(db *gorm.DB, rows *[]entity.BdrEmpresaSegmento) (err error) {
	//err := s.bd.Raw("SELECT S.ID AS ID, S.DESCRICAO AS Descricao FROM TBBDR_EMPRESA_SEGMENTO S WHERE S.SITUACAO = 'A' ORDER BY S.DESCRICAO").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
