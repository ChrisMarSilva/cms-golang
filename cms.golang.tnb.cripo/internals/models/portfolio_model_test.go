package models_test

import (
	"testing"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/models"
	"github.com/shopspring/decimal"
)

func TestPortfolio(t *testing.T) {
	portfolio := models.NewPortfolio()

	// Test case 1
	portfolio.Add("BTC", decimal.NewFromFloat(1.5), decimal.NewFromFloat(5000))
	if len(portfolio.Items) != 1 {
		t.Errorf("Expected portfolio to have 1 item, but got %d", len(portfolio.Items))
	}

	// Test case 2
	portfolio.Add("ETH", decimal.NewFromFloat(2.3), decimal.NewFromFloat(200))
	if len(portfolio.Items) != 2 {
		t.Errorf("Expected portfolio to have 2 items, but got %d", len(portfolio.Items))
	}
}
