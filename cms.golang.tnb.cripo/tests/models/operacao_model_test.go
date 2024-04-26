package models_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/models"
	"github.com/shopspring/decimal"
)

func TestOperacionString(t *testing.T) {
	op := models.Operacion{
		DataHora:      time.Now(),
		Tipo:          "Compra",
		CriptoDe:      "BRL",
		CriptoPara:    "ETH",
		Qtd:           decimal.NewFromFloat(1.5),
		AmountTax:     decimal.NewFromFloat(0.01),
		AmountTotal:   decimal.NewFromFloat(1.51),
		VlrPreco:      decimal.NewFromFloat(5000),
		TotVlrPreco:   decimal.NewFromFloat(7500),
		ValueTotalTax: decimal.NewFromFloat(75),
		TotVlr:        decimal.NewFromFloat(7575),
		ValueBalance:  decimal.NewFromFloat(10000),
	}

	// expected := "DateTime: " + op.DateTime.String() +
	// 	", Type: " + op.Type +
	// 	", OriginCoin: " + op.OriginCoin +
	// 	", DestinationCoin: " + op.DestinationCoin +
	// 	", QuantidadeMoeda: " + op.QuantidadeMoeda.String() +
	// 	", QuantidadeTaxa: " + op.QuantidadeTaxa.String() +
	// 	", QuantidadeTotal: " + op.QuantidadeTotal.String() +
	// 	", ValorPrecoMoeda: " + op.ValorPrecoMoeda.String() +
	// 	", ValorTotalMoeda: " + op.ValorTotalMoeda.String() +
	// 	", ValorTotalTaxa: " + op.ValorTotalTaxa.String() +
	// 	", ValorTotal: " + op.ValorTotal.String() +
	// 	", ValorSaldo: " + op.ValorSaldo.String()

	expected := fmt.Sprintf(
		"{%s: %s-%s-%s; QtdOp:%s; QtdTx:%s; QtdTot:%s; Price:%s; VlrTotCoin:%s; VlrTotTx:%s; VlrTot:%s; VlrSaldo:%s;}",
		op.DataHora.Format("02/01/2006 15:04:05"),
		op.Tipo,
		op.CriptoDe,
		op.CriptoPara,
		op.Qtd.String(),
		op.AmountTax.String(),
		op.AmountTotal.String(),
		op.VlrPreco.String(),
		op.TotVlrPreco.String(),
		op.ValueTotalTax.String(),
		op.TotVlr.String(),
		op.ValueBalance.String(),
	)

	result := op.String()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestNewOperacion(t *testing.T) {
	op := models.NewOperacion(
		[]string{"F7MYQKV7FYJ7AB", "23/04/2021 23:47:57", "Trade", "BRL", "-752,99", "0,00", "4.247,01", "Preço do Ativo (ETH/BRL): R$         12.752,39"},
		[]string{"F7MYQKV7FYJ7AB", "23/04/2021 23:47:57", "Trade", "ETH", "0,05904694", "0,00029523", "0,05875171", "Preço do Ativo (ETH/BRL): R$         12.752,39000000"},
	)

	expectedDateTime, _ := time.Parse("02/01/2006 15:04:05", "23/04/2021 23:47:57")

	expected := models.Operacion{
		DataHora:      expectedDateTime,
		Tipo:          "Compra",
		CriptoDe:      "BRL",
		CriptoPara:    "ETH",
		Qtd:           decimal.NewFromFloat(0.05904694),
		AmountTax:     decimal.NewFromFloat(0.00029523),
		AmountTotal:   decimal.NewFromFloat(0.05875171),
		VlrPreco:      decimal.NewFromFloat(12752.39000000),
		TotVlrPreco:   decimal.NewFromFloat(752.9896071866),
		ValueTotalTax: decimal.NewFromFloat(3.7648880997),
		TotVlr:        decimal.NewFromFloat(749.2247190869),
		ValueBalance:  decimal.NewFromFloat(4247.010392814), // 4247.01 // 4247.010392814
	}

	if reflect.ValueOf(op) == reflect.ValueOf(expected) { // if *op != expected {
		t.Errorf("Expected %+v, but got %+v", expected, *op)
	}
}
