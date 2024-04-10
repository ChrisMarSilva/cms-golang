package utils_test

import (
	"testing"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/utils"
	"github.com/shopspring/decimal"
)

func TestParseFloat(t *testing.T) {
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
		// //{"invalid", decimal.NewFromFloat(0)},
	}

	for _, test := range tests {
		result, err := utils.ParseFloat(test.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// log.Println("result:", result, "expected:", test.expected, "result.Cmp:", result.Cmp(test.expected))

		if result.Cmp(test.expected) != 0 { // if result.String() != test.expected.String() {
			t.Errorf("Expected %s - %T, but got %s - %T", test.expected, test.expected, result, result)
		}
	}
}
