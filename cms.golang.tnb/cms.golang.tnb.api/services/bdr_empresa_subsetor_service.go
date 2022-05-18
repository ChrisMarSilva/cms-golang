package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type BdrEmpresaSubSetorService struct {
	bd   *gorm.DB
	repo repository.BdrEmpresaSubSetorRepository
}

func NewBdrEmpresaSubSetorService(bd *gorm.DB, repo repository.BdrEmpresaSubSetorRepository) *BdrEmpresaSubSetorService {
	return &BdrEmpresaSubSetorService{
		bd:   bd,
		repo: repo,
	}
}

func (s *BdrEmpresaSubSetorService) GetLista() ([]dto.BdrEmpresaSubSetorLocal, error) {

	var rows []entity.BdrEmpresaSubSetor
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([]dto.BdrEmpresaSubSetorLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Descricao = row.Descricao
	}

	return list, nil
}
