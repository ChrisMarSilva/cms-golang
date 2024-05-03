package utils

import (
	"strings"

	"github.com/shopspring/decimal"
)

func ParseFloat(value string) (decimal.Decimal, error) {
	value = strings.Replace(value, "R$", "", -1)

	temPonto := strings.Contains(value, ".")
	temVirgula := strings.Contains(value, ",")

	if temPonto && temVirgula {
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", ".", -1)
	} else if !temPonto && temVirgula {
		value = strings.Replace(value, ",", ".", -1)
	}

	return decimal.NewFromString(value)
}
