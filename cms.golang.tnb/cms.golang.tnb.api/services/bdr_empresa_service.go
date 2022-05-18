package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type BdrEmpresaService struct {
	bd   *gorm.DB
	repo repository.BdrEmpresaRepository
}

func NewBdrEmpresaService(bd *gorm.DB, repo repository.BdrEmpresaRepository) *BdrEmpresaService {
	return &BdrEmpresaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *BdrEmpresaService) GetLista() ([]dto.BdrEmpresaLocal, error) {

	var rows []entity.BdrEmpresa
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.BdrEmpresaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
