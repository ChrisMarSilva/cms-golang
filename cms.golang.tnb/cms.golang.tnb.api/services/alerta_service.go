package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AlertaService struct {
	bd   *gorm.DB
	repo repository.AlertaRepository
}

func NewAlertaService(bd *gorm.DB, repo repository.AlertaRepository) *AlertaService {
	return &AlertaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AlertaService) GetLista() ([]dto.AlertaLocal, error) {

	var rows []entity.Alerta
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.AlertaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.SituacaoTelegram
		list[i].Nomee = row.SituacaoEmai
	}

	return list, nil
}
