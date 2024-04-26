package models_test

import (
	"fmt"
	"testing"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/models"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestPortfolio(t *testing.T) {
	// Arrange - Preparar o teste

	portfolio := models.NewPortfolio()

	// Act - Rodar o teste

	// portfolio.Add("BTC", "C", decimal.NewFromFloat(1), decimal.NewFromFloat(30000))
	// portfolio.Add("ETH", "C", decimal.NewFromFloat(1), decimal.NewFromFloat(10000))
	// portfolio.Add("ETH", "C", decimal.NewFromFloat(2), decimal.NewFromFloat(5000))

	// portfolio.Add("LINK", "C", decimal.NewFromFloat(200), decimal.NewFromFloat(14.1))
	// portfolio.Add("LINK", "C", decimal.NewFromFloat(300), decimal.NewFromFloat(15.06666666666667))
	// portfolio.Add("LINK", "C", decimal.NewFromFloat(400), decimal.NewFromFloat(14.05))

	// portfolio.Add("LINK", "C", decimal.NewFromFloat(200), decimal.NewFromFloat(14.05))    // 2.810,00
	// portfolio.Add("LINK", "C", decimal.NewFromFloat(300), decimal.NewFromFloat(15.03334)) // 4.510,00 = 7.320,00 / 500 = 14,64

	// portfolio.Add("LINK", "C", decimal.NewFromFloat(100), decimal.NewFromFloat(5))    // 500,00
	// portfolio.Add("LINK", "C", decimal.NewFromFloat(200), decimal.NewFromFloat(5.75)) //  1.150,00 = 1.650,00 / 500 = 5,50

	// portfolio.Add("LINK", "C", decimal.NewFromFloat(10), decimal.NewFromFloat(50)) // 500,00
	// portfolio.Add("LINK", "C", decimal.NewFromFloat(15), decimal.NewFromFloat(45)) // 675,00 = 1.175,00 / 25 = 47,00

	portfolio.Add("LINK", "C", decimal.NewFromFloat(30), decimal.NewFromFloat(50)) // 1.500,00 / 30 = 50,00
	portfolio.Add("LINK", "C", decimal.NewFromFloat(30), decimal.NewFromFloat(40)) // 1.200,00 = 2.700,00 /  60 = 45,00
	portfolio.Add("LINK", "C", decimal.NewFromFloat(40), decimal.NewFromFloat(30)) // 1.200,00 = 3.900,00 / 100 = 39,00
	portfolio.Add("LINK", "V", decimal.NewFromFloat(20), decimal.NewFromFloat(60)) //   780,00 = 3.120,00 /  80 = 39,00 -- 420,00 lucro
	portfolio.Add("LINK", "V", decimal.NewFromFloat(30), decimal.NewFromFloat(30)) // 1.170,00 = 1.950,00 /  50 = 39,00 -- 270,00 prejuízo
	portfolio.Add("LINK", "V", decimal.NewFromFloat(50), decimal.NewFromFloat(10)) //	1.950,00 = 2.600,00 /  0 = 39,00 -- 1.450,00 prejuízo

	/*

		20 * 39 = 780,00  comprado
		20 * 60 = 1.200,00 vendido
		780,00 total compra - 1200,00 total venda = 420,00 lucro

		30 * 39 = 1.170,00  comprado
		30 * 30 = 900,00 vendido
		1.170,00 total compra - 900,00 total venda = 270,00 prejuízo

		50 * 39 = 1.950,00  comprado
		50 * 10 = 500,00 vendido
		1.950,00 total compra - 500,00 total venda = 1.450,00 prejuízo

		total compra = 3.900,00
		total venda = 1.200,00 + 900,00 + 500,00 = 2.600,00
		total lucro/prejuizo = 3.900,00 - 2.600,00 = 1.300,00 prejuízo

	*/

	// Assert - Verificar as asserçõess

	assert.NotNil(t, portfolio, "portfolio is nil")
	assert.NotNil(t, portfolio.Items, "items of the portfolio is nil")
	assert.Equal(t, 1, len(portfolio.Items), "Expected portfolio to have 2 items, but got %d", len(portfolio.Items))

	for _, item := range portfolio.Items {
		fmt.Printf("Coin: %s, Amount: %s, Total: %s, Average: %s, Profit: %s\n", item.Coin, item.Amount.Round(2).String(), item.Total.Round(2).String(), item.AveragePrice.Round(2).String(), item.Profit.Round(2).String())

		assert.NotNil(t, item, "item of the portfolio is nil")
		assert.NotEmpty(t, item.Coin, "coin of the item is empty")
		assert.NotEqual(t, decimal.Zero, item.Amount, "amount of the item is zero")
		assert.NotEqual(t, decimal.Zero, item.Total, "total of the item is zero")
		assert.NotEqual(t, decimal.Zero, item.AveragePrice, "average price of the item is zero")

		//switch item.Coin {
		// case "BTC":
		// 	assert.Equal(t, decimal.NewFromFloat(1), item.Amount, fmt.Sprintf("Expected amount of %s to be 1, but got %s", item.Coin, item.Amount.String()))
		// 	assert.Equal(t, decimal.NewFromFloat(13760).Round(2), item.Total.Round(2), fmt.Sprintf("Expected total price of %s to be 6666.67, but got %s", item.Coin, item.Amount.String()))
		// 	assert.Equal(t, decimal.NewFromFloat(30000).Round(2), item.Average.Round(2), fmt.Sprintf("Expected average price of BTC to be 30000, but got %s", item.Coin, item.Amount.String()))
		// case "ETH":
		// 	assert.Equal(t, decimal.NewFromFloat(3), item.Amount, fmt.Sprintf("Expected amount of %s to be 3, but got %s", item.Amount.String()))
		// 	assert.Equal(t, decimal.NewFromFloat(13760).Round(2), item.Total.Round(2), fmt.Sprintf("Expected total price of %s to be 6666.67, but got %s", item.Coin, item.Amount.String()))
		// 	assert.Equal(t, decimal.NewFromFloat(6666.67).Round(2), item.Average.Round(2), fmt.Sprintf("Expected average price of %s to be 6666.67, but got %s", item.Coin, item.Amount.String()))
		// case "LINK":
		// 	assert.Equal(t, decimal.NewFromFloat(200+400), item.Amount, fmt.Sprintf("Expected amount of %s to be 900, but got %s", item.Coin, item.Amount.String()))
		// 	assert.Equal(t, decimal.NewFromFloat(2820+6420).Round(2), item.Total.Round(2), fmt.Sprintf("Expected total price of %s to be 6666.67, but got %s", item.Coin, item.Amount.String()))
		// 	//assert.Equal(t, decimal.NewFromFloat(6666.67).Round(2), item.Average.Round(2), fmt.Sprintf("Expected average price of ETH to be 6666.67, but got %s", item.Amount.String()))
		// default:
		// 	t.Errorf("Unexpected coin %s", item.Coin)
		// }
	}
}
