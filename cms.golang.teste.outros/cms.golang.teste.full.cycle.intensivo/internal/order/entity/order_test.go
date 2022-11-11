package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test ./...
// go test
// go test -v
// go test -run TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError -v

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "", Price: 0.0, Tax: 0.0}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 0.0, Tax: 0.0}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 100.0, Tax: 0.0}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAVAlidParams_WhenICallNewOder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 100.0, 10.0)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 100.0, 10.0)
	assert.NoError(t, err)
	assert.NoError(t, order.CalculateFinalPrice())
	assert.Equal(t, 110.0, order.FinalPrice)
}
