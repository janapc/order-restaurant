package usecase

import (
	"time"

	"github.com/janapc/order-restaurant/internal/entity"
)

type CreateOrderInputDTO struct {
	Dishes []entity.DishOrder `json:"dishes"`
	Tax    float64            `json:"tax"`
}

type CreateOrderOutputDTO struct {
	ID         string             `json:"id"`
	Dishes     []entity.DishOrder `json:"dishes"`
	Status     string             `json:"status"`
	Tax        float64            `json:"tax"`
	TotalPrice float64            `json:"total_price"`
	CreatedAt  time.Time          `json:"created_at"`
	CanceledAt *time.Time         `json:"canceled_at"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	DishRepository  entity.DishRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepositoryInterface, dishRepository entity.DishRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
		DishRepository:  dishRepository,
	}
}

func (c *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (*CreateOrderOutputDTO, error) {
	var dishes []entity.DishOrder
	for _, dish := range input.Dishes {
		dishDB, err := c.DishRepository.FindById(dish.ID)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, entity.DishOrder{
			ID:       dishDB.ID,
			Quantity: dish.Quantity,
			Price:    dishDB.Price,
		})
	}
	input.Dishes = dishes
	order := entity.NewOrder(input.Dishes, input.Tax)
	order.CalculateTotalPrice()
	err := c.OrderRepository.Create(order)
	if err != nil {
		return nil, err
	}
	return &CreateOrderOutputDTO{
		ID:         order.ID,
		Dishes:     order.Dishes,
		Status:     order.Status,
		Tax:        order.Tax,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		CanceledAt: order.CanceledAt,
	}, nil
}
