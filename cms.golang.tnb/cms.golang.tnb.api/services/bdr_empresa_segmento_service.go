package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type BdrEmpresaSegmentoService struct {
	bd   *gorm.DB
	repo repository.BdrEmpresaSegmentoRepository
}

func NewBdrEmpresaSegmentoService(bd *gorm.DB, repo repository.BdrEmpresaSegmentoRepository) *BdrEmpresaSegmentoService {
	return &BdrEmpresaSegmentoService{
		bd:   bd,
		repo: repo,
	}
}

func (s *BdrEmpresaSegmentoService) GetLista() ([]dto.BdrEmpresaSegmentoLocal, error) {

	var rows []entity.BdrEmpresaSegmento
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.BdrEmpresaSegmentoLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
