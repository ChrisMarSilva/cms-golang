package repository

import (
	"github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AlertaRepository struct {
}

func NewAlertaRepository() *AlertaRepository {
	return &AlertaRepository{}
}

func (repo AlertaRepository) GetLista(db *gorm.DB, rows *[]entity.Alerta) (err error) {
	//err := s.bd.Raw("SELECT U.ID AS ID, U.NOME AS Nome, U.NOME AS Nomee FROM TBUSUARIO U WHERE EXISTS( SELECT 1 FROM TBALERTA AL WHERE AL.IDUSUARIO = U.ID ) ORDER BY U.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
