package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type BdrEmpresaSetorService struct {
	bd   *gorm.DB
	repo repository.BdrEmpresaSetorRepository
}

func NewBdrEmpresaSetorService(bd *gorm.DB, repo repository.BdrEmpresaSetorRepository) *BdrEmpresaSetorService {
	return &BdrEmpresaSetorService{
		bd:   bd,
		repo: repo,
	}
}

func (s *BdrEmpresaSetorService) GetLista() ([]dto.BdrEmpresaSetorLocal, error) {

	var rows []entity.BdrEmpresaSetor
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.BdrEmpresaSetorLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
