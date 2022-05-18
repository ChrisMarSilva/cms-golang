package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type CriptoEmpresaService struct {
	bd   *gorm.DB
	repo repository.CriptoEmpresaRepository
}

func NewCriptoEmpresaService(bd *gorm.DB, repo repository.CriptoEmpresaRepository) *CriptoEmpresaService {
	return &CriptoEmpresaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *CriptoEmpresaService) GetLista() ([]dto.CriptoEmpresaLocal, error) {

	var rows []entity.CriptoEmpresa
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.CriptoEmpresaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
