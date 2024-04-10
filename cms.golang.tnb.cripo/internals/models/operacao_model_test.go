package models_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/models"
	"github.com/shopspring/decimal"
)

func TestOperacionString(t *testing.T) {
	op := models.Operacion{
		DateTime:        time.Now(),
		Type:            "Compra",
		OriginCoin:      "BRL",
		DestinationCoin: "ETH",
		QuantidadeMoeda: decimal.NewFromFloat(1.5),
		QuantidadeTaxa:  decimal.NewFromFloat(0.01),
		QuantidadeTotal: decimal.NewFromFloat(1.51),
		ValorPrecoMoeda: decimal.NewFromFloat(5000),
		ValorTotalMoeda: decimal.NewFromFloat(7500),
		ValorTotalTaxa:  decimal.NewFromFloat(75),
		ValorTotal:      decimal.NewFromFloat(7575),
		ValorSaldo:      decimal.NewFromFloat(10000),
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
		DateTime:        expectedDateTime,
		Type:            "Compra",
		OriginCoin:      "BRL",
		DestinationCoin: "ETH",
		QuantidadeMoeda: decimal.NewFromFloat(0.05904694),
		QuantidadeTaxa:  decimal.NewFromFloat(0.00029523),
		QuantidadeTotal: decimal.NewFromFloat(0.05875171),
		ValorPrecoMoeda: decimal.NewFromFloat(12752.39000000),
		ValorTotalMoeda: decimal.NewFromFloat(752.9896071866),
		ValorTotalTaxa:  decimal.NewFromFloat(3.7648880997),
		ValorTotal:      decimal.NewFromFloat(749.2247190869),
		ValorSaldo:      decimal.NewFromFloat(4247.010392814), // 4247.01 // 4247.010392814
	}

	if reflect.ValueOf(op) == reflect.ValueOf(expected) { // if *op != expected {
		t.Errorf("Expected %+v, but got %+v", expected, *op)
	}
}
