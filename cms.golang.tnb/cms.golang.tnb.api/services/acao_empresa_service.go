package service

import (
	dto "github.com/ChrisMarSilva/cms.golang.tnb.api/dtos"
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AcaoEmpresaService struct {
	bd   *gorm.DB
	repo repository.AcaoEmpresaRepository
}

func NewAcaoEmpresaService(bd *gorm.DB, repo repository.AcaoEmpresaRepository) *AcaoEmpresaService {
	return &AcaoEmpresaService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AcaoEmpresaService) GetLista() ([]dto.AcaoEmpresaLocal, error) {

	var rows []entity.AcaoEmpresa
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	// var start time.Time = time.Now()
	// log.Println("FIM:", time.Since(start))

	list := make([]dto.AcaoEmpresaLocal, len(rows))
	for i, row := range rows {
		list[i].ID = row.ID
		list[i].Nome = row.Nome
	}

	return list, nil
}
