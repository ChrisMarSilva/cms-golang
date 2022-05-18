package entity

import (
	"fmt"
)

type LogMonitor struct {
	CodErro      string `gorm:"column:COD_ERRO" gorm:"primary_key;"` // TBMONITOR.COD_ERRO // TBJDLOG_ERROS.NumErroID
	CodSistJD    string `gorm:"column:COD_SIST_JD"`                  // TBMONITOR.COD_SIST_JD // TBJDLOG_ERROS.CdSistema
	Data         string `gorm:"column:DATA_ERRO"`                    // TBMONITOR.DATA_ERRO // TBJDLOG_ERROS.DtErro
	Hora         string `gorm:"column:HORA_ERRO"`                    // TBMONITOR.HORA_ERRO // TBJDLOG_ERROS.HrErro
	Visto        string `gorm:"column:VISTO"`                        // TBMONITOR.VISTO // TBJDLOG_ERROS.CdErro
	Erro         string `gorm:"column:ERRO"`                         // TBMONITOR.ERRO // TBJDLOG_ERROS.Descricao
	NumMsg       string `gorm:"column:NUMERO_MENSAGEM"`              // TBMONITOR.NUMERO_MENSAGEM // TBJDLOG_ERROS.NumIdent
	Prioridade   string `gorm:"column:PRIORIDADE"`                   // TBMONITOR.PRIORIDADE // TBJDLOG_ERROS.NivelErro
	CodSistJDOri string `gorm:"column:Cod_Sist_JD_Ori"`              // TBMONITOR.Cod_Sist_JD_Ori // TBJDLOG_ERROS.CdOrigem
	CodErroUnico string `gorm:"column:Cod_Erro_Unico"`               // TBMONITOR.Cod_Erro_Unico // TBJDLOG_ERROS.NumErro
}

func (LogMonitor) TableName() string {
	return "TBMONITOR" // TBMONITOR // TBJDLOG_ERROS
}

func (row LogMonitor) ToString() string {
	return fmt.Sprintf("CodErro: %s; Data: %s; Hora: %s; CodSistJD: %s; CodSistJDOri: %s; NumMsg: %s; Erro: %s", row.CodErro, row.Data, row.Hora, row.CodSistJD, row.CodSistJDOri, row.NumMsg, row.Erro)
}
