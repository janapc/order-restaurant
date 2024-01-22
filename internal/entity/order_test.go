package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	dishes := []DishOrder{
		{
			ID: "123", Quantity: 2, Price: 4.0,
		},
		{
			ID: "1234", Quantity: 1, Price: 6.0,
		},
	}
	order := NewOrder(dishes, 10.0)
	assert.NotEmpty(t, order.CreatedAt)
	assert.NotEmpty(t, order.Dishes)
	assert.Equal(t, order.Tax, 10.0)
	assert.Empty(t, order.TotalPrice)
	assert.Empty(t, order.CanceledAt)
}

func TestCalculateOrder(t *testing.T) {
	dishes := []DishOrder{
		{
			ID: "123", Quantity: 2, Price: 4.0,
		},
		{
			ID: "1234", Quantity: 1, Price: 6.0,
		},
	}
	order := NewOrder(dishes, 0.05)
	assert.NotEmpty(t, order)
	order.CalculateTotalPrice()
	assert.NotEmpty(t, order.TotalPrice)
	assert.Equal(t, order.TotalPrice, 14.05)
}
