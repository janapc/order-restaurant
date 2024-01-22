package entity

import (
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
}

func NewOrder(dishes []DishOrder, tax float64) *Order {
	order := &Order{
		Dishes:     dishes,
		Tax:        tax,
		CreatedAt:  time.Now(),
		Status:     "CREATED",
		CanceledAt: nil,
	}
	return order
}

func (o *Order) CalculateTotalPrice() {
	total := 0.0
	for _, dish := range o.Dishes {
		total += dish.Price * float64(dish.Quantity)
	}
	o.TotalPrice = total + o.Tax
}
