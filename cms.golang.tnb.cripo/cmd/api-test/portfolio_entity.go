package main

// import (
// 	"log"

// 	"github.com/shopspring/decimal"
// )

// type Portfolio struct {
// 	Items map[int]ItemPortfolio
// }

// func NewPortfolio() *Portfolio {
// 	return &Portfolio{
// 		Items: make(map[int]ItemPortfolio),
// 	}
// }

// func (p *Portfolio) Add(coin, kind string, amount, value, tax decimal.Decimal) {
// 	totalAmount := decimal.NewFromFloat(0)
// 	totalValue := decimal.NewFromFloat(0)
// 	totalAverage := decimal.NewFromFloat(0)
// 	totalProfit := decimal.NewFromFloat(0)

// 	for i, item := range p.Items {
// 		if item.Coin == coin {
// 			if kind == "C" {
// 				totalAmount = item.Amount.Add(amount)
// 				totalValue = item.Total.Add(amount.Mul(value)).Add(tax)
// 				totalAverage = totalValue.Div(totalAmount)
// 				totalProfit = item.Profit
// 			} else {
// 				totalAmount = item.Amount.Sub(amount)
// 				totalValue = item.Total.Sub(amount.Mul(item.AveragePrice)) // .Sub(tax)
// 				totalAverage = item.AveragePrice
// 				totalProfit = item.Profit.Add(amount.Mul(value).Sub(totalAverage.Mul(amount))).Sub(tax)

// 				if totalAmount.Cmp(decimal.NewFromFloat(0)) == 0 {
// 					totalAverage = decimal.NewFromFloat(0)
// 				}
// 			}

// 			log.Printf(
// 				" ---> coin: %s; kind: %s; amount: %s; value: %s; tax: %s; totalAmount: %s; totalValue: %s; totalAverage: %s; totalProfit: %s; \n",
// 				coin, kind, amount.Round(8).String(), value.Round(8).String(), tax.Round(8).String(), totalAmount.Round(8).String(), totalValue.Round(8).String(), totalAverage.Round(8).String(), totalProfit.Round(8).String(),
// 			)

// 			p.Items[i] = *NewItemPortfolioe(coin, totalAmount, totalValue, totalAverage, totalProfit)
// 			return
// 		}
// 	}

// 	totalAmount = amount
// 	totalValue = amount.Mul(value).Add(tax)
// 	totalAverage = totalValue.Div(totalAmount)
// 	totalProfit = decimal.NewFromFloat(0)

// 	log.Printf(
// 		" ---> coin: %s; kind: %s; amount: %s; value: %s; tax: %s; totalAmount: %s; totalValue: %s; totalAverage: %s; totalProfit: %s; \n",
// 		coin, kind, amount.Round(8).String(), value.Round(8).String(), tax.Round(8).String(), totalAmount.Round(8).String(), totalValue.Round(8).String(), totalAverage.Round(8).String(), totalProfit.Round(8).String(),
// 	)

// 	p.Items[len(p.Items)] = *NewItemPortfolioe(coin, totalAmount, totalValue, totalAverage, totalProfit)
// }

// type ItemPortfolio struct {
// 	Coin         string
// 	Amount       decimal.Decimal
// 	Total        decimal.Decimal
// 	AveragePrice decimal.Decimal
// 	Profit       decimal.Decimal
// }

// func NewItemPortfolioe(coin string, amount, total, averagePrice, profit decimal.Decimal) *ItemPortfolio {
// 	return &ItemPortfolio{
// 		Coin:         coin,
// 		Amount:       amount,
// 		Total:        total,
// 		AveragePrice: averagePrice,
// 		Profit:       profit,
// 	}
// }
