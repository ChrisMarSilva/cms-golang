package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type AcaoEmpresaAtivoRepository struct {
}

func NewAcaoEmpresaAtivoRepository() *AcaoEmpresaAtivoRepository {
	return &AcaoEmpresaAtivoRepository{}
}

func (repo AcaoEmpresaAtivoRepository) GetLista(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	//err := s.Db.Raw("SELECT ID AS ID, CODIGO AS Codigo FROM TBEMPRESA_ATIVO ORDER BY CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompleto(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	// type ResultStruct struct {
	// 	ID     int64  `json:"id"`
	// 	Codigo string `json:"codigo"`
	// 	Tipo   string `json:"tipo"`
	// 	Nome   string `json:"nome"`
	// }

	// var rows []ResultStruct
	// err := s.bd.Raw(`
	// 	SELECT A.ID AS ID, A.CODIGO AS Codigo, 'ACAO' AS Tipo,   IFNULL(E.NOMRESUMIDO, E.NOME) AS Nome FROM TBEMPRESA_ATIVO A JOIN TBEMPRESA E ON ( E.ID = A.IDEMPRESA ) WHERE A.SITUACAO = 'A'
	// 	UNION
	// 	SELECT F.ID AS ID, F.CODIGO AS Codigo, 'FII' AS Tipo,    F.NOME AS Nome FROM TBFII_FUNDOIMOB F WHERE F.SITUACAO IN ('A','E')
	// 	UNION
	// 	SELECT F.ID AS ID, F.CODIGO AS Codigo, 'ETF' AS Tipo,    F.NOME AS Nome FROM TBETF_INDICE F WHERE F.SITUACAO IN ('A','E')
	// 	UNION
	// 	SELECT F.ID AS ID, F.CODIGO AS Codigo, 'BDR' AS Tipo,    F.NOME AS Nome FROM TBBDR_EMPRESA F WHERE F.SITUACAO IN ('A','E')
	// 	UNION
	// 	SELECT F.ID AS ID, F.CODIGO AS Codigo, 'CRIPTO' AS Tipo, F.NOME AS Nome FROM TBCRIPTO_EMPRESA F WHERE F.SITUACAO IN ('A','E')
	// 	ORDER BY CODIGO
	// `).Scan(&rows).Error
	// if err != nil {
	// 	return nil, err
	// }
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompletoAcao(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	// err := s.bd.Raw("SELECT E.CODIGO AS Codigo FROM TBEMPRESA_ATIVO E WHERE E.SITUACAO = 'A' ORDER BY E.CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompletoFii(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	// err := s.bd.Raw("SELECT E.ID AS ID, E.CODIGO AS Codigo FROM TBFII_FUNDOIMOB E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompletoEtf(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	// err := s.bd.Raw("SELECT E.ID AS ID, E.CODIGO AS Codigo FROM TBETF_INDICE E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompletoBrd(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	//err := s.bd.Raw("SELECT E.CODIGO AS Codigo FROM TBBDR_EMPRESA E WHERE E.SITUACAO = 'A' ORDER BY E.CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo AcaoEmpresaAtivoRepository) GetListaCodigoCompletoCripto(db *gorm.DB, rows *[]entity.AcaoEmpresaAtivo) (err error) {
	err = db.Find(rows).Error
	//err := s.bd.Raw("SELECT E.ID AS ID, E.CODIGO AS Codigo, E.NOME AS Nome FROM TBCRIPTO_EMPRESA E WHERE E.SITUACAO IN ('A', 'E') ORDER BY E.CODIGO").Scan(&rows).Error
	if err != nil {
		return err
	}
	return nil
}
