package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	"gorm.io/gorm"
)

type LogMonitorRepositoryMSSQL struct {
	CdSistema       string
	CdSistemaOrigem string
	db              *gorm.DB
}

var FNumSequencial = 0

func NewLogMonitorRepositoryMSSQL(cdSistema string, cdSistemaOrigem string, db *gorm.DB) *LogMonitorRepositoryMSSQL {
	return &LogMonitorRepositoryMSSQL{
		CdSistema:       cdSistema,
		CdSistemaOrigem: cdSistemaOrigem,
		db:              db,
	}
}

func (repo LogMonitorRepositoryMSSQL) IncrementaErroSeq() {
	FNumSequencial++
	if FNumSequencial > 999 {
		FNumSequencial = 0
	}
}

func (repo LogMonitorRepositoryMSSQL) Inserir(tipo string, nivel string, erro string) (err error) {

	if len(tipo) > 20 {
		tipo = tipo[0:20]
	}

	if len(erro) > 3900 {
		erro = erro[0:3900]
	}

	for i := 0; i < 100; i++ { // Tentar inserir ateh conseguir ou no maximo 100x

		repo.IncrementaErroSeq()

		t := time.Now()
		// fmt.Sprintf("%02d%02d%02d", t.Year(), t.Month(), t.Day())
		// fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())

		var logMonitor entity.LogMonitor
		logMonitor.CodErro = "MO" + t.Format("060102150405") + strings.Trim(t.Format(".000"), ".") + fmt.Sprintf("%03d", FNumSequencial) // MOYYMMDDhhnnsszzz000 = 20 caracteres
		logMonitor.CodSistJD = repo.CdSistema
		logMonitor.Data = t.Format("20060102")
		logMonitor.Hora = t.Format("15:04:05")
		logMonitor.Visto = "N"
		logMonitor.Erro = erro
		logMonitor.NumMsg = tipo
		logMonitor.Prioridade = nivel
		logMonitor.CodSistJDOri = repo.CdSistemaOrigem
		logMonitor.CodErroUnico = "MO" + t.Format("060102150405") + strings.Trim(t.Format(".000"), ".") + fmt.Sprintf("%03d", FNumSequencial) // MOYYMMDDhhnnsszzz000 = 20 caracteres

		err = repo.db.Create(logMonitor).Error
		if err != nil {
			if strings.Contains(err.Error(), "PRIMARY KEY") || strings.Contains(err.Error(), "DUPLICATE KEY") || strings.Contains(err.Error(), "CHAVE DUPLICADA") { // "Violação da restrição PRIMARY KEY"
				log.Println("continue - #", i, "- Erro:", err)
				continue
			} else {
				log.Println("Errooooo:", err)
				return err
			}
		}

		return nil

	} // for i := 0; i < 3; i++ {

	return nil
}
