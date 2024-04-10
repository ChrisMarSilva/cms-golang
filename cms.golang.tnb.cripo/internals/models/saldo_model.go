package models

import (
	"log"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/utils"
	"github.com/shopspring/decimal"
)

type Balance struct {
	TotalDeposit    decimal.Decimal
	TotalWithdrawal decimal.Decimal
	Items           map[int]ItemBalance
}

func NewBalance() *Balance {
	return &Balance{
		TotalDeposit:    decimal.NewFromFloat(0),
		TotalWithdrawal: decimal.NewFromFloat(0),
		Items:           make(map[int]ItemBalance),
	}
}

func (b *Balance) AddDeposit(datahora string, valor string, saldo string) {
	value, err := utils.ParseFloat(valor)
	if err != nil {
		log.Println("fValor", err)
		return
	}

	b.TotalDeposit = b.TotalDeposit.Add(value)
	b.Items[len(b.Items)] = *NewItemBalance(datahora, "Depósito", valor, saldo)
}

func (b *Balance) AddWithdrawal(datahora string, valor string, saldo string) {
	value, err := utils.ParseFloat(valor)
	if err != nil {
		log.Println("fValor", err)
		return
	}

	b.TotalWithdrawal = b.TotalWithdrawal.Add(value.Abs())
	b.Items[len(b.Items)] = *NewItemBalance(datahora, "Retirada", valor, saldo)
}

type ItemBalance struct {
	DateTime time.Time
	Type     string
	Value    decimal.Decimal
	Total    decimal.Decimal
}

func NewItemBalance(datahora string, tipo string, valor string, saldo string) *ItemBalance {
	datetime, _ := utils.ParseTime(datahora)
	value, _ := utils.ParseFloat(valor)
	total, _ := utils.ParseFloat(saldo)

	return &ItemBalance{
		DateTime: datetime,
		Type:     tipo,
		Value:    value,
		Total:    total,
	}
}
