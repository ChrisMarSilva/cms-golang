package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AlertaAssinaturaRepository struct {
}

func NewAlertaAssinaturaRepository() *AlertaAssinaturaRepository {
	return &AlertaAssinaturaRepository{}
}

func (repo AlertaAssinaturaRepository) GetLista(db *gorm.DB, rows *[]entity.AlertaAssinatura) (err error) {
	//err := s.bd.Raw("SELECT U.ID AS ID, U.NOME AS Nome, U.NOME AS Nomee FROM TBUSUARIO U WHERE EXISTS( SELECT 1 FROM TBALERTA_ASSINATURA AA WHERE AA.IDUSUARIO = U.ID ) ORDER BY U.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
