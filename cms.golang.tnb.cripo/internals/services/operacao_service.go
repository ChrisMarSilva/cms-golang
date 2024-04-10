package services

import (
	"log"
	"strings"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/models"
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
	balance := models.NewBalance()
	portfolio := models.NewPortfolio()

	//log.Println("")
	//log.Println("OPERATIONS...")
	for idx := 0; idx < totLines; idx++ {
		lineOne := rows[idx]
		datahora := strings.TrimSpace(lineOne[1])
		descricao := strings.TrimSpace(lineOne[2])
		moeda := strings.TrimSpace(lineOne[3])
		valor := strings.TrimSpace(lineOne[4])
		saldo := strings.TrimSpace(lineOne[6])

		if descricao == "descrição" || descricao == "EarnApplication" {
			continue
		}

		if descricao == "Depósito" && moeda == "BRL" {
			balance.AddDeposit(datahora, valor, saldo)
		} else if descricao == "Retirada" && moeda == "BRL" {
			balance.AddWithdrawal(datahora, valor, saldo)
		} else if descricao == "Trade" && moeda == "BRL" {
			if idx+1 < totLines {
				idx++
				lineTwo := rows[idx]
				op := models.NewOperacion(lineOne, lineTwo)
				portfolio.Add(op.DestinationCoin, op.QuantidadeTotal, op.ValorPrecoMoeda)
				//log.Println("Lines:", idx, "and", idx+1, "\t", "Operacion:", op)
			}
		} else {
			log.Println(idx, "\tTYPE NOT DEFINED\t", lineOne)
		}
	}

	// log.Println("")
	// log.Println("TOTAL BALANCE...")
	// log.Println("TOTAL DEPOSIT\t: R$", balance.TotalDeposit.String())
	// log.Println("TOTAL WITHDRAWAL\t: R$", balance.TotalWithdrawal.String())
	// log.Println("CURRENT BALANCE\t: R$", balance.TotalDeposit.Sub(balance.TotalWithdrawal).String())

	log.Println("")
	log.Println("PORTFOLIO...")
	for _, item := range portfolio.Items {
		log.Println("Coin:", item.Coin, "\t\tAmount:", item.Amount.String(), "\tAverage Price:", item.AveragePrice.String())
	}

	return nil
}
