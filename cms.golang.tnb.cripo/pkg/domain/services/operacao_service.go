package services

import (
	"log"
	"strings"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.domain/models"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

type OperacionService struct {
	Filename string
	Sheet    string
}

func NewOperacionService(filename string, sheet string) *OperacionService {
	return &OperacionService{
		Filename: filename,
		Sheet:    sheet,
	}
}

func (op *OperacionService) readFile() ([][]string, error) {
	f, err := excelize.OpenFile(op.Filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	return f.GetRows(op.Sheet)
}

func (op *OperacionService) ProcessFile() error {
	rows, err := op.readFile()
	if err != nil {
		log.Println("op.ReadFile():", err)
		return err
	}

	totLines := len(rows)
	operacoes := make(map[int]models.Operacao)

	for idx := 0; idx < totLines; idx++ {
		lineOne := rows[idx]
		typeOperation := strings.TrimSpace(lineOne[2])
		coin := strings.TrimSpace(lineOne[3])

		if typeOperation == "descrição" {
			continue
		}

		if typeOperation == "Trade" && coin == "BRL" {
			if idx+1 < totLines {
				idx++
				lineTwo := rows[idx]
				operacoes[len(operacoes)] = *models.NewOperacao(lineOne, lineTwo)
			}
		} else {
			log.Println(idx, "\tTYPE NOT DEFINED\t", lineOne)
		}
	}

	portfolio := models.NewPortfolio()
	totLines = len(operacoes)
	qtdAtual := decimal.NewFromFloat(0)
	vlrPrecoMedioAtual := decimal.NewFromFloat(0)

	for idx := 0; idx < totLines; idx++ {
		oper := operacoes[idx]

		qtdAtual, vlrPrecoMedioAtual = oper.CalcVlrPrecoMedio(qtdAtual, vlrPrecoMedioAtual)
		oper.VlrValorizacao = oper.CalcVlrValorizacao()
		oper.PercValorizacao = oper.CalcPercValorizacao()

		//portfolio.Add(oper) // coin, kind string, amount, value, tax decimal.Decimal
		//log.Println(oper)
	}

	for _, item := range portfolio.Items {
		log.Printf("Coin: %s; Amount: %s; Total: %s; Average Price: %s; Profit: %s; \n", item.Coin, item.Amount.Round(8).String(), item.Total.Round(8).String(), item.AveragePrice.Round(8).String(), item.Profit.Round(8).String())
	}

	return nil
}
