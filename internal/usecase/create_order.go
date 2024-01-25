package usecase

import (
	"os"
	"time"

	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/infra/queue"
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
	Queue           queue.QueueInterface
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepositoryInterface,
	dishRepository entity.DishRepositoryInterface,
	queue queue.QueueInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
		DishRepository:  dishRepository,
		Queue:           queue,
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
			ID:       dish.ID,
			Quantity: dish.Quantity,
			Price:    dishDB.Price,
		})
	}
	input.Dishes = dishes
	order, err := entity.NewOrder(input.Dishes, input.Tax, "PENDING")
	if err != nil {
		return nil, err
	}
	order.CalculateTotalPrice()
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	c.Queue.Publish(order.ID, "", queueName)
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
