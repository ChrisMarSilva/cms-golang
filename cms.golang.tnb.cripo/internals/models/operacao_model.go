package models

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/utils"
	"github.com/shopspring/decimal"
)

type Operacion struct {
	DateTime        time.Time
	Type            string
	OriginCoin      string
	DestinationCoin string
	QuantidadeMoeda decimal.Decimal
	QuantidadeTaxa  decimal.Decimal
	QuantidadeTotal decimal.Decimal
	ValorPrecoMoeda decimal.Decimal
	ValorTotalMoeda decimal.Decimal
	ValorTotalTaxa  decimal.Decimal
	ValorTotal      decimal.Decimal
	ValorSaldo      decimal.Decimal
}

func NewOperacion(lineOne []string, lineTwo []string) *Operacion {
	// Line One Data
	lineOneCoin := strings.TrimSpace(lineOne[3])
	lineOneValue := strings.TrimSpace(lineOne[4])
	LineOneBalance := strings.TrimSpace(lineOne[6])

	// Line Two Data
	lineTwoDateTime := strings.TrimSpace(lineTwo[1])
	lineTwoCoin := strings.TrimSpace(lineTwo[3])
	lineTwoValue := strings.TrimSpace(lineTwo[4])
	lineTwoTax := strings.TrimSpace(lineTwo[5])
	LineTwoBalance := strings.TrimSpace(lineTwo[6])
	lineTwoData := ""
	if len(lineTwo) > 7 {
		lineTwoData = strings.ReplaceAll(strings.TrimSpace(lineTwo[7]), "\n", "")
	}

	datetime, err := utils.ParseTime(lineTwoDateTime)
	if err != nil {
		log.Println("datahora", err)
		return nil
	}

	typeOperacion := "Compra" // Compra, Venda // se for negativo = compra // se for positivo = venda // valor total da operacao
	balanceOperacion, err := utils.ParseFloat(lineOneValue)
	if err != nil {
		log.Println("balance", err)
		return nil
	}
	if balanceOperacion.GreaterThan(decimal.NewFromFloat(0)) {
		typeOperacion = "Venda"
	}

	// quantidade de compra ou venda
	quantidadeMoeda, err := utils.ParseFloat(lineTwoValue)
	if err != nil {
		log.Println("quantidadeMoeda", err)
		return nil
	}

	// quantidade de taxa a debitar
	quantidadeTaxa, err := utils.ParseFloat(lineTwoTax)
	if err != nil {
		log.Println("quantidadeTaxa", err)
		return nil
	}

	// quantidade real da operacao
	quantidadeTotal, err := utils.ParseFloat(LineTwoBalance)
	if err != nil {
		log.Println("quantidadeTotal", err)
		return nil
	}

	// valor da moeda na hora da operacao
	if strings.Contains(lineTwoData, "Preço") {
		re := regexp.MustCompile("[^0-9.,]")
		lineTwoData = re.ReplaceAllString(lineTwoData, "")
		lineTwoData = strings.TrimSpace(lineTwoData)
	}

	valorPrecoMoeda, err := utils.ParseFloat(lineTwoData)
	if err != nil {
		log.Println("valorPrecoMoeda", err)
		return nil
	}

	// saldo restante
	valorSaldo, err := utils.ParseFloat(LineOneBalance)
	if err != nil {
		log.Println("saldo", err)
		return nil
	}

	//-------------------------------------------------------------------------
	//-------------------------------------------------------------------------

	//log.Println(1, "\t\t\t\t", linhaOne)
	//log.Println(2, "\t\t\t\t", linhaTwo)
	// log.Println("linhaTwo[6]:", linhaTwo[6])
	// log.Println("linhaTwoSaldo:", linhaTwoSaldo)
	// log.Println("quantidadeTotal:", quantidadeTotal)

	// fmt.Printf("reflect.linhaTwo[6]: %T\n", linhaTwo[6])
	// fmt.Printf("reflect.linhaTwoSaldo: %T\n", linhaTwoSaldo)
	// fmt.Printf("reflectquantidadeTotal: %T\n", quantidadeTotal)

	// linhaTwoValor := strings.TrimSpace(linhaTwo[4])
	// linhaTwoTaxa := strings.TrimSpace(linhaTwo[5])
	// linhaTwoSaldo := strings.TrimSpace(linhaTwo[6])

	//-------------------------------------------------------------------------
	//-------------------------------------------------------------------------

	return &Operacion{
		DateTime:        datetime,
		Type:            typeOperacion,
		OriginCoin:      lineOneCoin,
		DestinationCoin: lineTwoCoin,
		QuantidadeMoeda: quantidadeMoeda,
		QuantidadeTaxa:  quantidadeTaxa,
		QuantidadeTotal: quantidadeTotal,
		ValorPrecoMoeda: valorPrecoMoeda,
		ValorTotalMoeda: quantidadeMoeda.Mul(valorPrecoMoeda),
		ValorTotalTaxa:  quantidadeTaxa.Mul(valorPrecoMoeda),
		ValorTotal:      quantidadeTotal.Mul(valorPrecoMoeda),
		ValorSaldo:      valorSaldo,
	}
}

func (op Operacion) String() string {
	return fmt.Sprintf(
		"{%s: %s-%s-%s; QtdOp:%s; QtdTx:%s; QtdTot:%s; Price:%s; VlrTotCoin:%s; VlrTotTx:%s; VlrTot:%s; VlrSaldo:%s;}",
		op.DateTime.Format("02/01/2006 15:04:05"),
		op.Type,
		op.OriginCoin,
		op.DestinationCoin,
		op.QuantidadeMoeda.String(),
		op.QuantidadeTaxa.String(),
		op.QuantidadeTotal.String(),
		op.ValorPrecoMoeda.String(),
		op.ValorTotalMoeda.String(),
		op.ValorTotalTaxa.String(),
		op.ValorTotal.String(),
		op.ValorSaldo.String(),
	)
}
