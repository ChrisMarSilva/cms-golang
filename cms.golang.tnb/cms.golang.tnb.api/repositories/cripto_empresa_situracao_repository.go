package repository

import (
	entity "github.com/ChrisMarSilva/cms.golang.tnb.api/entities"
	"gorm.io/gorm"
)

type CriptoEmpresaSituacaoRepository struct {
}

func NewCriptoEmpresaSituacaoRepository() *CriptoEmpresaSituacaoRepository {
	return &CriptoEmpresaSituacaoRepository{}
}

func (repo CriptoEmpresaSituacaoRepository) GetSituacoes(db *gorm.DB, rows *[]entity.CriptoEmpresaSituacao) (err error) {

	err = db.Find(rows).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo CriptoEmpresaSituacaoRepository) GetSituacao(db *gorm.DB, row *entity.CriptoEmpresaSituacao, codigo string) (err error) {

	db.Raw("SELECT CODIGO, DESCRICAO FROM TBCRIPTO_EMPRESA_ST WHERE CODIGO = ?", codigo).Scan(&row)
	if err != nil {
		return err
	}

	return nil
}
