package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type FiiEmpresaTipoService struct {
	bd   *gorm.DB
	repo repository.FiiEmpresaTipoRepository
}

func NewFiiEmpresaTipoService(bd *gorm.DB, repo repository.FiiEmpresaTipoRepository) *FiiEmpresaTipoService {
	return &FiiEmpresaTipoService{
		bd:   bd,
		repo: repo,
	}
}

func (s *FiiEmpresaTipoService) GetLista() ([]dto.FiiEmpresaTipoLocal, error) {

	var rows []entity.FiiEmpresaTipo
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.FiiEmpresaTipoLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
