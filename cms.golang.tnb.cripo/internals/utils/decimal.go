package utils

import (
	"strings"

	"github.com/shopspring/decimal"
)

func ParseFloat(value string) (decimal.Decimal, error) {
	temPonto := strings.Contains(value, ".")
	temVirgula := strings.Contains(value, ",")

	value = strings.Replace(value, "R$", "", -1)
	//log.Println("value:", value, "temPonto:", temPonto, "temVirgula:", temVirgula)

	if temPonto && temVirgula {
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", ".", -1)
	} else if !temPonto && temVirgula {
		value = strings.Replace(value, ",", ".", -1)
	} // else if temPonto && !temVirgula {
	// value = strings.Replace(value, ".", "", -1)
	//}

	//log.Println("value:", value)
	return decimal.NewFromString(value)
}
