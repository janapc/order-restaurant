package usecase

import (
	"github.com/janapc/order-restaurant/internal/entity"
)

type CancelOrderInputDTO struct {
	ID string `json:"id"`
}

type CancelOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCancelOrderUseCase(orderRepository entity.OrderRepositoryInterface) *CancelOrderUseCase {
	return &CancelOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CancelOrderUseCase) Execute(input CancelOrderInputDTO) error {
	err := c.OrderRepository.Cancel(input.ID)
	if err != nil {
		return err
	}
	return nil
}
