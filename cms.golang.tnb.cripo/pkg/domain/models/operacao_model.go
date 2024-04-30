package models

// import (
// 	"fmt"
// 	"log"
// 	"regexp"
// 	"strings"
// 	"time"

// 	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/utils"
// 	"github.com/shopspring/decimal"
// )

// type Operacion struct {
// 	DataHora        time.Time       // DateTime
// 	Tipo            string          // Type
// 	Cripto          string          // OriginCoin // DestinationCoin
// 	Qtd             decimal.Decimal // AmountCoin
// 	VlrTaxa         decimal.Decimal // ValueTax
// 	VlrPreco        decimal.Decimal // ValuePrice
// 	TotVlrPreco     decimal.Decimal // ValueTotalCoin
// 	TotVlr          decimal.Decimal // ValueTotal
// 	VlrCusto        decimal.Decimal
// 	VlrPrecoMedio   decimal.Decimal
// 	VlrValorizacao  decimal.Decimal
// 	PercValorizacao decimal.Decimal
// }

// func NewOperacion(lineOne, lineTwo []string) *Operacion {
// 	op := new(Operacion)

// 	dataHora := op.GetDataHora(lineOne)
// 	//log.Println("dataHora", dataHora)

// 	cripto := op.GetCripto(lineOne, lineTwo)
// 	//log.Println("cripto", cripto)

// 	qtd := op.GetQtd(lineTwo)
// 	//log.Println("qtd", qtd)

// 	tipo := op.GetTipo(lineTwo, qtd)
// 	//log.Println("tipo", tipo)

// 	qtd = qtd.Abs()

// 	vlrPreco := op.GetVlrPreco(lineTwo)
// 	//log.Println("vlrPreco", vlrPreco)

// 	vlrTaxa := op.GetVlrTaxa(lineOne, lineTwo, tipo, vlrPreco)
// 	//log.Println("vlrTaxa", vlrTaxa)

// 	totVlrPreco := op.GetTotVlrPreco(qtd, vlrPreco)
// 	//log.Println("totVlrPreco", totVlrPreco)

// 	totVlr := op.GetTotVlr(tipo, totVlrPreco, vlrTaxa)
// 	//log.Println("totVlr", totVlr)

// 	vlrCusto := op.GetVlrCusto(qtd, totVlr)
// 	// log.Println("vlrCusto", vlrCusto)

// 	vlrPrecoMedio := decimal.NewFromFloat(0)
// 	vlrValorizacao := decimal.NewFromFloat(0)
// 	percValorizacao := decimal.NewFromFloat(0)

// 	//log.Println("qtd", qtd, "totVlrPreco", totVlrPreco, "totVlr", totVlr, "vlrCusto", vlrCusto)

// 	return &Operacion{
// 		DataHora:        dataHora,
// 		Tipo:            tipo,
// 		Cripto:          cripto,
// 		Qtd:             qtd,
// 		VlrTaxa:         vlrTaxa,
// 		VlrPreco:        vlrPreco,
// 		TotVlrPreco:     totVlrPreco,
// 		TotVlr:          totVlr,
// 		VlrCusto:        vlrCusto,
// 		VlrPrecoMedio:   vlrPrecoMedio,
// 		VlrValorizacao:  vlrValorizacao,
// 		PercValorizacao: percValorizacao,
// 	}
// }

// func (Operacion) GetDataHora(line []string) time.Time {
// 	dataHoraStr := strings.TrimSpace(line[1])

// 	dataHora, err := utils.ParseTime(dataHoraStr)
// 	if err != nil {
// 		log.Panic("dataHora", err)
// 	}

// 	return dataHora
// }

// func (Operacion) GetTipo(line []string, qtd decimal.Decimal) string {
// 	tipo := "C"
// 	if qtd.LessThan(decimal.NewFromFloat(0)) {
// 		tipo = "V"
// 	}

// 	return tipo
// }

// func (Operacion) GetCripto(lineOne, lineTwo []string) string {
// 	criptoDe := strings.TrimSpace(lineOne[3])
// 	criptoPara := strings.TrimSpace(lineTwo[3])

// 	return fmt.Sprintf("%s/%s", criptoDe, criptoPara)
// }

// func (Operacion) GetQtd(line []string) decimal.Decimal {
// 	qtdStr := strings.TrimSpace(line[4])
// 	qtdTaxaStr := strings.TrimSpace(line[5])

// 	qtdTaxa, err := utils.ParseFloat(qtdTaxaStr)
// 	if err != nil {
// 		log.Panic("qtd", err)
// 	}

// 	qtd, err := utils.ParseFloat(qtdStr)
// 	if err != nil {
// 		log.Panic("qtd", err)
// 	}

// 	return qtd.Sub(qtdTaxa)
// }

// func (Operacion) GetVlrTaxa(lineOne, lineTwo []string, tipo string, vlrPreco decimal.Decimal) decimal.Decimal {
// 	if tipo == "C" {
// 		qtdTaxaStr := strings.TrimSpace(lineTwo[5])

// 		qtdTaxa, err := utils.ParseFloat(qtdTaxaStr)
// 		if err != nil {
// 			log.Panic("qtdTaxa", err)
// 		}

// 		vlrTaxa := qtdTaxa.Mul(vlrPreco)
// 		return vlrTaxa
// 	} else {
// 		vlrTaxStr := strings.TrimSpace(lineOne[5])

// 		vlrTaxa, err := utils.ParseFloat(vlrTaxStr)
// 		if err != nil {
// 			log.Panic("vlrTaxa", err)
// 		}

// 		return vlrTaxa
// 	}
// }

// func (Operacion) GetVlrPreco(lineTwo []string) decimal.Decimal {
// 	vlrPrecoStr := ""
// 	if len(lineTwo) > 7 {
// 		vlrPrecoStr = strings.ReplaceAll(strings.TrimSpace(lineTwo[7]), "\n", "")
// 		if strings.Contains(vlrPrecoStr, "Preço") {
// 			re := regexp.MustCompile("[^0-9.,]")
// 			vlrPrecoStr = re.ReplaceAllString(vlrPrecoStr, "")
// 			vlrPrecoStr = strings.TrimSpace(vlrPrecoStr)
// 		}
// 	}

// 	vlrPreco, err := utils.ParseFloat(vlrPrecoStr)
// 	if err != nil {
// 		log.Panic("vlrPreco", err)
// 	}

// 	return vlrPreco
// }

// func (Operacion) GetTotVlrPreco(qtd, vlrPreco decimal.Decimal) decimal.Decimal {
// 	totVlrPreco := decimal.NewFromFloat(0)

// 	if qtd.GreaterThan(decimal.NewFromFloat(0)) && vlrPreco.GreaterThan(decimal.NewFromFloat(0)) {
// 		totVlrPreco = qtd.Mul(vlrPreco)
// 	}

// 	return totVlrPreco
// }

// func (Operacion) GetTotVlr(tipo string, totVlrPreco, vlrTaxa decimal.Decimal) decimal.Decimal {
// 	totVlr := decimal.NewFromFloat(0)

// 	if tipo == "C" {
// 		totVlr = totVlrPreco.Add(vlrTaxa)
// 	} else {
// 		totVlr = totVlrPreco.Sub(vlrTaxa)
// 	}

// 	return totVlr
// }

// func (Operacion) GetVlrCusto(qtd, totVlr decimal.Decimal) decimal.Decimal {
// 	vlrCusto := decimal.NewFromFloat(0)

// 	if qtd.GreaterThan(decimal.NewFromFloat(0)) && totVlr.GreaterThan(decimal.NewFromFloat(0)) {
// 		vlrCusto = totVlr.Div(qtd)
// 	}

// 	return vlrCusto
// }

// func (op *Operacion) CalcVlrPrecoMedio(qtdAtual, vlrPrecoMedioAtual decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
// 	op.VlrPrecoMedio = decimal.NewFromFloat(0)

// 	if op.Tipo == "C" {
// 		totVlrCusto := op.Qtd.Mul(op.VlrCusto)
// 		if qtdAtual.GreaterThan(decimal.NewFromFloat(0)) && vlrPrecoMedioAtual.GreaterThan(decimal.NewFromFloat(0)) {
// 			totVlrCusto = totVlrCusto.Add(qtdAtual.Mul(vlrPrecoMedioAtual))
// 		}
// 		qtdAtual = qtdAtual.Add(op.Qtd)
// 		if qtdAtual.GreaterThan(decimal.NewFromFloat(0)) && totVlrCusto.GreaterThan(decimal.NewFromFloat(0)) {
// 			vlrPrecoMedioAtual = totVlrCusto.Div(qtdAtual)
// 		}
// 	} else if op.Tipo == "V" {
// 		qtdAtual = qtdAtual.Sub(op.Qtd)
// 	}

// 	op.VlrPrecoMedio = vlrPrecoMedioAtual
// 	return qtdAtual, vlrPrecoMedioAtual
// }

// func (op Operacion) CalcVlrValorizacao() decimal.Decimal {
// 	vlrValorizacao := decimal.NewFromFloat(0)

// 	if op.Tipo == "V" && op.Qtd.GreaterThan(decimal.NewFromFloat(0)) && op.VlrPrecoMedio.GreaterThan(decimal.NewFromFloat(0)) {
// 		vlrValorizacao = op.TotVlr.Sub(op.Qtd.Mul(op.VlrPrecoMedio))
// 	}

// 	return vlrValorizacao
// }

// func (op Operacion) CalcPercValorizacao() decimal.Decimal {
// 	vlrZero := decimal.NewFromFloat(0)
// 	percValorizacao := decimal.NewFromFloat(0)

// 	if op.Tipo == "V" && op.Qtd.GreaterThan(vlrZero) && op.VlrPrecoMedio.GreaterThan(vlrZero) && !op.VlrValorizacao.Equal(vlrZero) {
// 		// perc_valorizacao = (tot_vlr_valorizacao / (vlr_preco_medio * quant)) * 100
// 		percValorizacao = op.VlrValorizacao.Div(op.Qtd.Mul(op.VlrPrecoMedio)).Mul(decimal.NewFromFloat(100))
// 		//log.Println("Qtd", op.Qtd, "; PrcMed", op.VlrPrecoMedio, "; VlrVlrz", op.VlrValorizacao, "; %", percValorizacao)
// 	}

// 	return percValorizacao
// }

// func (op Operacion) String() string {
// 	texto := fmt.Sprintf(
// 		"%s-%s; %s; Qtd:%s; Prç:%s; TotPrç:%s; Tx:%s; Tot:%s; Cst: %s; PrçMed: %s; ",
// 		op.DataHora.Format("02/01/2006 15:04:05"),
// 		op.Tipo,
// 		op.Cripto,
// 		op.Qtd.Round(8).String(),
// 		op.VlrPreco.Round(8).String(),
// 		op.TotVlrPreco.Round(8).String(),
// 		op.VlrTaxa.Round(8).String(),
// 		op.TotVlr.Round(8).String(),
// 		op.VlrCusto.Round(8).String(),
// 		op.VlrPrecoMedio.Round(8).String(),
// 	)

// 	if op.Tipo == "V" {
// 		texto += fmt.Sprintf(
// 			"Vlrzç: %s (%s%s);",
// 			op.VlrValorizacao.Round(8).String(),
// 			op.PercValorizacao.Round(8).String(),
// 			"%",
// 		)
// 	}

// 	return texto
// }
