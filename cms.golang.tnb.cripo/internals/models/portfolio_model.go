package models

import (
	"github.com/shopspring/decimal"
)

type Portfolio struct {
	Items map[int]ItemPortfolio
}

func NewPortfolio() *Portfolio {
	return &Portfolio{
		Items: make(map[int]ItemPortfolio),
	}
}

func (p *Portfolio) Add(coin string, amount decimal.Decimal, value decimal.Decimal) {
	for i, item := range p.Items {
		if item.Coin == coin {
			p.Items[i].Amount.Add(amount)
			p.Items[i].AveragePrice.Add(value).Div(p.Items[i].Amount)
			return
		}
	}

	p.Items[len(p.Items)] = *NewItemPortfolioe(coin, amount, value.Div(amount))
}

type ItemPortfolio struct {
	Coin         string
	Amount       decimal.Decimal
	AveragePrice decimal.Decimal
}

func NewItemPortfolioe(coin string, amount decimal.Decimal, averagePrice decimal.Decimal) *ItemPortfolio {
	return &ItemPortfolio{
		Coin:         coin,
		Amount:       amount,
		AveragePrice: averagePrice,
	}
}
