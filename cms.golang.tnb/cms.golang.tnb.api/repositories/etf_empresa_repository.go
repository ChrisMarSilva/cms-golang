package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type EtfEmpresaRepository struct {
}

func NewEtfEmpresaRepository() *EtfEmpresaRepository {
	return &EtfEmpresaRepository{}
}

func (repo EtfEmpresaRepository) GetLista(db *gorm.DB, rows *[]entity.EtfEmpresa) (err error) {
	//err := s.bd.Raw("SELECT E.ID AS ID, E.NOME AS Nome, E.NOME AS Nomee FROM TBETF_INDICE E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
