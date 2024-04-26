package utils_test

import (
	"log"
	"testing"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestParseFloat(t *testing.T) {
	// Arrange - Preparar o teste
	tests := []struct {
		input    string
		expected decimal.Decimal
	}{
		{"10.5", decimal.NewFromFloat(10.5)},
		{"-5.25", decimal.NewFromFloat(-5.25)},
		{"0", decimal.NewFromFloat(0)},
		{"3.14159", decimal.NewFromFloat(3.14159)},
		{"3,14159", decimal.NewFromFloat(3.14159)},
		{"0,05904694", decimal.NewFromFloat(0.05904694)},
		{"0,00029523", decimal.NewFromFloat(0.00029523)},
		{"0,05875171", decimal.NewFromFloat(0.05875171)},
		{"0,02016491", decimal.NewFromFloat(0.02016491)},
		{"0,00010082", decimal.NewFromFloat(0.00010082)},
		{"0,07881579", decimal.NewFromFloat(0.07881579)},
		{"0,31163303", decimal.NewFromFloat(0.31163303)},
		{"0,00155817", decimal.NewFromFloat(0.00155817)},
		{"0,38889066", decimal.NewFromFloat(0.38889066)},
		{"830,523840", decimal.NewFromFloat(830.523840)},
		{"4,152619", decimal.NewFromFloat(4.152619)},
		{"826,371221", decimal.NewFromFloat(826.371221)},
		{"26,18075191", decimal.NewFromFloat(26.18075191)},
		{"0,13090376", decimal.NewFromFloat(0.13090376)},
		{"26,04984815", decimal.NewFromFloat(26.04984815)},
		{"3,60227952", decimal.NewFromFloat(3.60227952)},
		{"0,01801140", decimal.NewFromFloat(0.01801140)},
		{"3,58426812", decimal.NewFromFloat(3.58426812)},
		{"4.247,01", decimal.NewFromFloat(4247.01)},
		{"12.752,39000000", decimal.NewFromFloat(12752.39000000)},
	}

	for _, test := range tests {
		// Act - Rodar o teste
		result, err := utils.ParseFloat(test.input)

		// Assert - Verificar as asserções
		assert.NotNil(t, err, "Unexpected error: %v", err)
		assert.Equal(t, 0, result.Cmp(test.expected), "Expected %s - %T, but got %s - %T", test.expected, test.expected, result, result)
	}
}

func TestAvg01(t *testing.T) {
	// Arrange - Preparar o teste

	amount := decimal.NewFromFloat(0)
	total := decimal.NewFromFloat(0)
	avg := decimal.NewFromFloat(0)

	// Act - Rodar o teste

	amount = amount.Add(decimal.NewFromFloat(10))
	total = total.Add(decimal.NewFromFloat(5000))
	avg = total.Div(amount)

	// Assert - Verificar as asserções

	log.Println("Amount:", amount.String())
	log.Println("Total:", total.String())
	log.Println("Avg:", avg.String())
}

func TestAvg02(t *testing.T) {
	// Arrange - Preparar o teste

	tests := []struct {
		kind   string
		amount decimal.Decimal
		value  decimal.Decimal
		profit decimal.Decimal
	}{
		// {"C", decimal.NewFromFloat(10), decimal.NewFromFloat(5000)},
		// {"C", decimal.NewFromFloat(20), decimal.NewFromFloat(10000)},

		// {"C", decimal.NewFromFloat(200), decimal.NewFromFloat(14.1)},              // 2.820,00
		// {"C", decimal.NewFromFloat(300), decimal.NewFromFloat(15.06666666666667)}, // 4.520,00 =  7.340,00
		// {"C", decimal.NewFromFloat(400), decimal.NewFromFloat(14.05)},             // 5.620,00 = 12.960,00

		// {"C", decimal.NewFromFloat(200), decimal.NewFromFloat(14.05)},    // 2.810,00
		// {"C", decimal.NewFromFloat(300), decimal.NewFromFloat(15.03334)}, // 4.510,00 = 7.320,00 / 500 = 14,64

		// {"C", decimal.NewFromFloat(100), decimal.NewFromFloat(5)},    // 500,00
		// {"C", decimal.NewFromFloat(200), decimal.NewFromFloat(5.75)}, // 1.150,00 = 1.650,00 / 500 = 5,50

		// {"C", decimal.NewFromFloat(10), decimal.NewFromFloat(50)}, // 500,00
		// {"C", decimal.NewFromFloat(15), decimal.NewFromFloat(45)}, // 675,00 = 1.175,00 / 25 = 47,00

		{"C", decimal.NewFromFloat(30), decimal.NewFromFloat(50), decimal.NewFromFloat(0)}, // 1.500,00
		{"C", decimal.NewFromFloat(30), decimal.NewFromFloat(40), decimal.NewFromFloat(0)}, // 1.200,00 = 2.700,00 /  60 = 45,00
		{"C", decimal.NewFromFloat(40), decimal.NewFromFloat(30), decimal.NewFromFloat(0)}, // 1.200,00 = 3.900,00 / 100 = 39,00
		{"V", decimal.NewFromFloat(20), decimal.NewFromFloat(60), decimal.NewFromFloat(0)}, //   780,00 = 3.120,00 /  80 = 39,00 -- 420,00 lucro
		{"V", decimal.NewFromFloat(20), decimal.NewFromFloat(30), decimal.NewFromFloat(0)}, //   480,00 = 2.640,00 /  60 = 39,00 -- 240,00 lucro
		{"V", decimal.NewFromFloat(50), decimal.NewFromFloat(10), decimal.NewFromFloat(0)}, //   230,00 = 2.310,00 / 110 = 39,00 -- 270,00 prejuízo
	}

	totalAmount := decimal.NewFromFloat(0)
	totalValue := decimal.NewFromFloat(0)
	totalprofit := decimal.NewFromFloat(0)
	avg := decimal.NewFromFloat(0)

	for i, test := range tests {
		// Act - Rodar o teste

		if test.kind == "C" {
			totalAmount = totalAmount.Add(test.amount)
			totalValue = totalValue.Add(test.amount.Mul(test.value))
			avg = totalValue.Div(totalAmount)
		} else {
			totalAmount = totalAmount.Sub(test.amount)
			totalValue = totalValue.Sub(test.amount.Mul(avg))
			//avg = totalValue.Div(totalAmount)
			totalprofit = totalprofit.Add(test.amount.Mul(test.value).Sub(avg.Mul(test.amount)))
		}

		log.Println("Nro:", i+1)
		log.Println("Amount:", totalAmount.Round(2).String())
		log.Println("Total:", totalValue.Round(2).String())
		log.Println("Avg:", avg.Round(2).String())
		if test.kind == "V" {
			log.Println("profit:", totalprofit.Round(2).String())
		}
		log.Println("")

		// Assert - Verificar as asserções
		//assert.NotNil(t, err, "Unexpected error: %v", err)
		//assert.Equal(t, 0, result.Cmp(test.expected), "Expected %s - %T, but got %s - %T", test.expected, test.expected, result, result)
	}

	// avg = totalValue.Div(totalAmount)

	// log.Println("Total Amount:", totalAmount.String())
	// log.Println("Total Value:", totalValue.String())
	// log.Println("Avg:", avg.String())
}

// func calcularPrecoMedio(precos []decimal.Decimal) decimal.Decimal {
//     total := decimal.NewFromInt(0)
//     for _, preco := range precos {
//         total = total.Add(preco)
//     }
//     media := total.Div(decimal.NewFromInt(int64(len(precos))))
//     return media
// }
