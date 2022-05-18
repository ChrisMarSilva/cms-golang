package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AcaoEmpresaSubSetorService struct {
	bd   *gorm.DB
	repo repository.AcaoEmpresaSubSetorRepository
}

func NewAcaoEmpresaSubSetorService(bd *gorm.DB, repo repository.AcaoEmpresaSubSetorRepository) *AcaoEmpresaSubSetorService {
	return &AcaoEmpresaSubSetorService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AcaoEmpresaSubSetorService) GetLista() ([]dto.AcaoEmpresaSubSetorLocal, error) {

	var rows []entity.AcaoEmpresaSubSetor
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.AcaoEmpresaSubSetorLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
