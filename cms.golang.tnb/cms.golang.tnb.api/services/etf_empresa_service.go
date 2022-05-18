package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type EtfEmpresaService struct {
	bd   *gorm.DB
	repo repository.EtfEmpresaRepository
}

func NewEtfEmpresaService(bd *gorm.DB, repo repository.EtfEmpresaRepository) *EtfEmpresaService {
	return &EtfEmpresaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *EtfEmpresaService) GetLista() ([]dto.EtfEmpresaLocal, error) {

	var rows []entity.EtfEmpresa
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.EtfEmpresaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
