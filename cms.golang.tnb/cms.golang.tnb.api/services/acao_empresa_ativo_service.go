package service

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	repository "github.com/ChrisMarSilva/cms.golang.tnb.api/repositories"
	"gorm.io/gorm"
)

type AcaoEmpresaAtivoService struct {
	bd   *gorm.DB
	repo repository.AcaoEmpresaAtivoRepository
}

func NewAcaoEmpresaAtivoService(bd *gorm.DB, repo repository.AcaoEmpresaAtivoRepository) *AcaoEmpresaAtivoService {
	return &AcaoEmpresaAtivoService{
		bd:   bd,
		repo: repo,
	}
}

func (s *AcaoEmpresaAtivoService) GetLista() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetLista(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 2)
		list[i][0] = row.Codigo
		list[i][1] = row.ID
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompleto() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompleto(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 5)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
		list[i][2] = row.Tipo
		list[i][3] = row.ID
		//list[i][4] = row.Nome
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompletoAcao() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompletoAcao(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 2)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompletoFii() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompletoFii(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 3)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
		list[i][2] = row.ID
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompletoEtf() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompletoEtf(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 3)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
		list[i][2] = row.ID
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompletoBrd() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompletoBrd(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 2)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
	}

	return list, nil
}

func (s *AcaoEmpresaAtivoService) GetListaCodigoCompletoCripto() ([][]interface{}, error) {

	var rows []entity.AcaoEmpresaAtivo
	err := s.repo.GetListaCodigoCompletoCripto(s.bd, &rows)
	if err != nil {
		return nil, err
	}

	list := make([][]interface{}, len(rows))
	for i, row := range rows {
		list[i] = make([]interface{}, 4)
		list[i][0] = row.Codigo
		list[i][1] = row.Codigo
		list[i][2] = row.ID
		//list[i][3] = row.Nome
	}

	return list, nil
}
