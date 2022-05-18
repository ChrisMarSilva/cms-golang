package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type FiiEmpresaService struct {
	bd   *gorm.DB
	repo repository.FiiEmpresaRepository
}

func NewFiiEmpresaService(bd *gorm.DB, repo repository.FiiEmpresaRepository) *FiiEmpresaService {
	return &FiiEmpresaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *FiiEmpresaService) GetLista() ([]dto.FiiEmpresaLocal, error) {

	var rows []entity.FiiEmpresa
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.FiiEmpresaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
