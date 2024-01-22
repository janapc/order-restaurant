package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDishValid(t *testing.T) {
	dish, err := NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	assert.NoError(t, err)
	assert.NotEmpty(t, dish)
	assert.NotEmpty(t, dish.ID)
	assert.Equal(t, dish.Name, "brigadeiro")
	assert.Equal(t, dish.Description, "delicioso brigadeiro artesanal")
	assert.Equal(t, dish.Price, 4.0)
}

func TestDishInvalid(t *testing.T) {
	dish, err := NewDish("brigadeiro", "delicioso brigadeiro artesanal", 0.0)
	assert.Error(t, err, "the price is required and has to be greater than 0")
	assert.Empty(t, dish)
}
