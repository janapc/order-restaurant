package entity

import (
	"errors"
	"time"
)

type DishOrder struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	ID         string      `json:"id"`
	Dishes     []DishOrder `json:"dishes"`
	Tax        float64     `json:"tax"`
	Status     string      `json:"status"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	CanceledAt *time.Time  `json:"canceled_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func NewOrder(dishes []DishOrder, tax float64, status string) (*Order, error) {
	order := &Order{
		Dishes:     dishes,
		Tax:        tax,
		Status:     status,
		CreatedAt:  time.Now(),
		CanceledAt: nil,
		UpdatedAt:  time.Now(),
	}
	if err := order.IsValid(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.Status == "" {
		return errors.New("the status is required")
	}
	if len(o.Dishes) == 0 {
		return errors.New("the dishes is required")
	}
	if o.Tax < 0.0 {
		return errors.New("the tax is required and has to be greater than 0 or equal 0")
	}
	return nil
}

func (o *Order) CalculateTotalPrice() {
	total := 0.0
	for _, dish := range o.Dishes {
		total += dish.Price * float64(dish.Quantity)
	}
	o.TotalPrice = total + o.Tax
}
