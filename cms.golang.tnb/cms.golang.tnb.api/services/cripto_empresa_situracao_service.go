package service

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type CriptoEmpresaSituacaoService struct {
	bd   *gorm.DB
	repo repository.CriptoEmpresaSituacaoRepository
}

func NewCriptoEmpresaSituacaoService(bd *gorm.DB, repo repository.CriptoEmpresaSituacaoRepository) *CriptoEmpresaSituacaoService {
	return &CriptoEmpresaSituacaoService{
		bd:   bd,
		repo: repo,
	}
}

func (s *CriptoEmpresaSituacaoService) GetSituacoes() ([]entity.CriptoEmpresaSituacao, error) {
	var list []entity.CriptoEmpresaSituacao
	var situacaoRepo repository.CriptoEmpresaSituacaoRepository
	err := situacaoRepo.GetSituacoes(s.bd, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *CriptoEmpresaSituacaoService) GetSituacao(codigo string) (entity.CriptoEmpresaSituacao, error) {
	var row entity.CriptoEmpresaSituacao
	var situacaoRepo repository.CriptoEmpresaSituacaoRepository
	err := situacaoRepo.GetSituacao(s.bd, &row, codigo)
	if err != nil {
		return entity.CriptoEmpresaSituacao{}, err
	}
	return row, nil
}
