package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Dish struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewDish(name string, description string, price float64) (*Dish, error) {
	dish := Dish{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
	}
	err := dish.isValid()
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (m *Dish) isValid() error {
	if m.Name == "" {
		return errors.New("the name is required")
	}
	if m.Description == "" {
		return errors.New("the description is required")
	}
	if m.Price <= 0.0 {
		return errors.New("the price is required and has to be greater than 0")
	}
	return nil
}
