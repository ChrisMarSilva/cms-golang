package models_test

import (
	"testing"
	"time"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/models"
	"github.com/shopspring/decimal"
)

func TestNewBalance(t *testing.T) {
	balance := models.NewBalance()

	// Assert that the balance is initialized correctly
	if balance.TotalDeposit.Cmp(decimal.Zero) != 0 {
		t.Errorf("Expected TotalDeposit to be zero, got %s", balance.TotalDeposit.String())
	}
	if balance.TotalWithdrawal.Cmp(decimal.Zero) != 0 {
		t.Errorf("Expected TotalWithdrawal to be zero, got %s", balance.TotalWithdrawal.String())
	}
	if len(balance.Items) != 0 {
		t.Errorf("Expected Items to be empty, got %d", len(balance.Items))
	}
}

func TestAddDeposit(t *testing.T) {
	// Add a deposit to the balance
	datahora := "01/01/2022 10:00:00"
	valor := "100.00"
	saldo := "100.00"

	balance := models.NewBalance()
	balance.AddDeposit(datahora, valor, saldo)

	// Assert that the deposit is added correctly
	if balance.TotalDeposit.Cmp(decimal.RequireFromString(valor)) != 0 {
		t.Errorf("Expected TotalDeposit to be %s, got %s", valor, balance.TotalDeposit.String())
	}
	if len(balance.Items) != 1 {
		t.Errorf("Expected Items to have 1 element, got %d", len(balance.Items))
	}

	item := balance.Items[0]
	if item.DateTime.Format("02/01/2006 15:04:05") != datahora {
		t.Errorf("Expected DateTime to be %s, got %s", datahora, item.DateTime.Format("02/01/2006 15:04:05"))
	}
	if item.Type != "Depósito" {
		t.Errorf("Expected Type to be 'Depósito', got %s", item.Type)
	}
	if item.Value.Cmp(decimal.RequireFromString(valor)) != 0 {
		t.Errorf("Expected Value to be %s, got %s", valor, item.Value.String())
	}
	if item.Total.Cmp(decimal.RequireFromString(saldo)) != 0 {
		t.Errorf("Expected Total to be %s, got %s", saldo, item.Total.String())
	}
}

func TestAddWithdrawal(t *testing.T) {
	balance := models.NewBalance()

	// Add a withdrawal to the balance
	datahora := "01/01/2022 10:00:00"
	valor := "50.00"
	saldo := "50.00"
	balance.AddWithdrawal(datahora, valor, saldo)

	// Assert that the withdrawal is added correctly
	if balance.TotalWithdrawal.Cmp(decimal.RequireFromString(valor)) != 0 {
		t.Errorf("Expected TotalWithdrawal to be %s, got %s", valor, balance.TotalWithdrawal.String())
	}
	if len(balance.Items) != 1 {
		t.Errorf("Expected Items to have 1 element, got %d", len(balance.Items))
	}

	item := balance.Items[0]
	if item.DateTime.Format("02/01/2006 15:04:05") != datahora {
		t.Errorf("Expected DateTime to be %s, got %s", datahora, item.DateTime.Format("02/01/2006 15:04:05"))
	}
	if item.Type != "Retirada" {
		t.Errorf("Expected Type to be 'Retirada', got %s", item.Type)
	}
	if item.Value.Cmp(decimal.RequireFromString(valor)) != 0 {
		t.Errorf("Expected Value to be %s, got %s", valor, item.Value.String())
	}
	if item.Total.Cmp(decimal.RequireFromString(saldo)) != 0 {
		t.Errorf("Expected Total to be %s, got %s", saldo, item.Total.String())
	}
}

func TestNewItemBalance(t *testing.T) {
	datahora := "01/01/2022 10:00:00"
	tipo := "Depósito"
	valor := "100.00"
	saldo := "100.00"

	item := models.NewItemBalance(datahora, tipo, valor, saldo)

	// Assert that the item is created correctly
	expectedDateTime, _ := time.Parse("02/01/2006 15:04:05", datahora)
	if !item.DateTime.Equal(expectedDateTime) {
		t.Errorf("Expected DateTime to be %s, got %s", expectedDateTime.Format("02/01/2006 15:04:05"), item.DateTime.Format("02/01/2006 15:04:05"))
	}
	if item.Type != tipo {
		t.Errorf("Expected Type to be %s, got %s", tipo, item.Type)
	}
	if item.Value.Cmp(decimal.RequireFromString(valor)) != 0 {
		t.Errorf("Expected Value to be %s, got %s", valor, item.Value.String())
	}
	if item.Total.Cmp(decimal.RequireFromString(saldo)) != 0 {
		t.Errorf("Expected Total to be %s, got %s", saldo, item.Total.String())
	}
}
