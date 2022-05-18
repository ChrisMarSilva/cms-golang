package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AcaoEmpresaSegmentoService struct {
	bd   *gorm.DB
	repo repository.AcaoEmpresaSegmentoRepository
}

func NewAcaoEmpresaSegmentoService(bd *gorm.DB, repo repository.AcaoEmpresaSegmentoRepository) *AcaoEmpresaSegmentoService {
	return &AcaoEmpresaSegmentoService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AcaoEmpresaSegmentoService) GetLista() ([]dto.AcaoEmpresaSegmentoLocal, error) {

	var rows []entity.AcaoEmpresaSegmento
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.AcaoEmpresaSegmentoLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
