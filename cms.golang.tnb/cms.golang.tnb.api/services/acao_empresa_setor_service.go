package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AcaoEmpresaSetorService struct {
	bd   *gorm.DB
	repo repository.AcaoEmpresaSetorRepository
}

func NewAcaoEmpresaSetorService(bd *gorm.DB, repo repository.AcaoEmpresaSetorRepository) *AcaoEmpresaSetorService {
	return &AcaoEmpresaSetorService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AcaoEmpresaSetorService) GetLista() ([]dto.AcaoEmpresaSetorLocal, error) {

	var rows []entity.AcaoEmpresaSetor
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.AcaoEmpresaSetorLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
