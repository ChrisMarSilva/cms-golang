package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type CorretoraListaService struct {
	bd   *gorm.DB
	repo repository.CorretoraListaRepository
}

func NewCorretoraListaService(bd *gorm.DB, repo repository.CorretoraListaRepository) *CorretoraListaService {
	return &CorretoraListaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *CorretoraListaService) GetLista() ([]dto.CorretoraListaLocal, error) {

	var rows []entity.CorretoraLista
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.CorretoraListaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
