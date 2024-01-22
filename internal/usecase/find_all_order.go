package usecase

import (
	"time"

	"github.com/janapc/order-restaurant/internal/entity"
)

type FindAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type FindAllOrdersOutputDTO struct {
	ID         string             `json:"id"`
	Dishes     []entity.DishOrder `json:"dishes"`
	Status     string             `json:"status"`
	Tax        float64            `json:"tax"`
	TotalPrice float64            `json:"total_price"`
	CreatedAt  time.Time          `json:"created_at"`
	CanceledAt *time.Time         `json:"canceled_at"`
}

func NewFindAllOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *FindAllOrdersUseCase {
	return &FindAllOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (f *FindAllOrdersUseCase) Execute() ([]FindAllOrdersOutputDTO, error) {
	orders, err := f.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var output []FindAllOrdersOutputDTO
	for _, order := range orders {
		output = append(output, FindAllOrdersOutputDTO{
			ID:         order.ID,
			Dishes:     order.Dishes,
			Tax:        order.Tax,
			Status:     order.Status,
			TotalPrice: order.TotalPrice,
			CreatedAt:  order.CreatedAt,
			CanceledAt: order.CanceledAt,
		})
	}
	return output, nil
}
