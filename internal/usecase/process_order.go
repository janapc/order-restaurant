package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/janapc/order-restaurant/internal/entity"
)

type ProcessOrder struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewProcessOrder(orderRepository entity.OrderRepositoryInterface) *ProcessOrder {
	return &ProcessOrder{
		OrderRepository: orderRepository,
	}
}

func (p *ProcessOrder) Execute(data []byte) error {
	id := string(data)
	order, err := p.OrderRepository.FindById(id)
	if err != nil || order.Status == "CANCELED" {
		msg := fmt.Sprintf("error processing this message: %s", id)
		return errors.New(msg)
	}
	err = p.OrderRepository.Status("PROCESSING", id)
	if err != nil {
		return err
	}
	fmt.Println("prepare to send...")
	time.Sleep(8 * time.Second)
	err = p.OrderRepository.Status("SENT", id)
	if err != nil {
		return err
	}
	fmt.Println("sent ...")
	return nil
}
