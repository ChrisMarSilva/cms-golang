package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AlertaAssinaturaService struct {
	bd   *gorm.DB
	repo repository.AlertaAssinaturaRepository
}

func NewAlertaAssinaturaService(bd *gorm.DB, repo repository.AlertaAssinaturaRepository) *AlertaAssinaturaService {
	return &AlertaAssinaturaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AlertaAssinaturaService) GetLista() ([]dto.AlertaAssinaturaLocal, error) {

	var rows []entity.AlertaAssinatura
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.AlertaAssinaturaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.TipoAlerta
		list[i].Nomee = row.TipoAssinatura
	}

	return list, nil
}
