package usecase

import (
	"errors"
	"slices"

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
	order, err := c.OrderRepository.FindById(input.ID)
	statusInvalid := []string{"PROCESSING", "SENT"}
	if err != nil {
		return err
	}
	if slices.Contains(statusInvalid, order.Status) {
		return errors.New("the order cannot be canceled")
	}
	err = c.OrderRepository.Cancel(input.ID)
	if err != nil {
		return err
	}
	return nil
}
