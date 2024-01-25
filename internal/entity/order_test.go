package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	dishes := []DishOrder{
		{
			ID: "123", Quantity: 2,
		},
		{
			ID: "1234", Quantity: 1,
		},
	}
	order, err := NewOrder(dishes, 10.0, "PENDING")
	assert.NoError(t, err)
	assert.NotEmpty(t, order.CreatedAt)
	assert.NotEmpty(t, order.Dishes)
	assert.Equal(t, order.Tax, 10.0)
	assert.Empty(t, order.TotalPrice)
	assert.Empty(t, order.CanceledAt)
}

func TestInvalidOrder(t *testing.T) {
	dishes := []DishOrder{
		{
			ID: "123", Quantity: 2,
		},
		{
			ID: "1234", Quantity: 1,
		},
	}
	order, err := NewOrder(dishes, -0.1, "PENDING")
	assert.NotEmpty(t, err)
	assert.Empty(t, order)
	assert.Equal(t, err.Error(), "the tax is required and has to be greater than 0 or equal 0")
}

func TestCalculateOrder(t *testing.T) {
	dishes := []DishOrder{
		{
			ID: "123", Quantity: 2, Price: 2.0,
		},
		{
			ID: "1234", Quantity: 1, Price: 10.0,
		},
	}
	order, err := NewOrder(dishes, 0.05, "PENDING")
	assert.NoError(t, err)
	assert.NotEmpty(t, order)
	order.CalculateTotalPrice()
	assert.NotEmpty(t, order.TotalPrice)
	assert.Equal(t, order.TotalPrice, 14.05)
}
