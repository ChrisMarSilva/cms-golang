package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type CorretoraListaRepository struct {
}

func NewCorretoraListaRepository() *CorretoraListaRepository {
	return &CorretoraListaRepository{}
}

func (repo CorretoraListaRepository) GetLista(db *gorm.DB, rows *[]entity.CorretoraLista) (err error) {
	//err := s.bd.Raw("SELECT C.ID AS ID, C.NOME AS Nome FROM TBCORRETORA_LISTA C WHERE C.SITUACAO = 'A' ORDER BY C.NOME").Scan(&list).Error
	err = db.Find(rows).Error
	if err != nil {
		return err
	}
	return nil
}
