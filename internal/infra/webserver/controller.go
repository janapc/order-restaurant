package webserver

import (
	"encoding/json"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/infra/queue"
	"github.com/janapc/order-restaurant/internal/usecase"
)

type Controller struct {
	DishRepository  entity.DishRepositoryInterface
	OrderRepository entity.OrderRepositoryInterface
	Queue           queue.QueueInterface
}

func NewController(dishRepository entity.DishRepositoryInterface,
	orderRepository entity.OrderRepositoryInterface,
	queue queue.QueueInterface) *Controller {
	return &Controller{
		DishRepository:  dishRepository,
		OrderRepository: orderRepository,
		Queue:           queue,
	}
}

func (c *Controller) RegisterDish(ctx *fiber.Ctx) error {
	var input usecase.SaveDishInputDTO
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Body is wrong")
	}
	usecase := usecase.NewSaveDishUseCase(c.DishRepository)
	output, err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	json, _ := json.Marshal(output)
	return ctx.Status(fiber.StatusOK).Send(json)
}

func (c *Controller) UpdateDish(ctx *fiber.Ctx) error {
	var input usecase.UpdateDishInputDTO
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("Body is wrong")
	}
	id := ctx.Params("id")
	input.ID = id
	usecase := usecase.NewUpdateDishUseCase(c.DishRepository)
	err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *Controller) RemoveDish(ctx *fiber.Ctx) error {
	var input usecase.RemoveDishInputDTO
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	id := ctx.Params("id")
	input.ID = id
	usecase := usecase.NewRemoveDishUseCase(c.DishRepository)
	err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *Controller) FindAllDish(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	usecase := usecase.NewFindAllUseCase(c.DishRepository)
	output, err := usecase.Execute()
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	json, _ := json.Marshal(output)
	return ctx.Status(fiber.StatusOK).Send(json)
}

func (c *Controller) FindByIdDish(ctx *fiber.Ctx) error {
	var input usecase.FindByIdInputDTO
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	id := ctx.Params("id")
	input.ID = id
	usecase := usecase.NewFindByIdUseCase(c.DishRepository)
	output, err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	json, _ := json.Marshal(output)
	return ctx.Status(fiber.StatusOK).Send(json)
}

func (c *Controller) CreateOrder(ctx *fiber.Ctx) error {
	var input usecase.CreateOrderInputDTO
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	if err := ctx.BodyParser(&input); err != nil {
		slog.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).SendString("Body is wrong")
	}
	usecase := usecase.NewCreateOrderUseCase(c.OrderRepository, c.DishRepository, c.Queue)
	output, err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	json, _ := json.Marshal(output)
	return ctx.Status(fiber.StatusOK).Send(json)
}

func (c *Controller) CancelOrder(ctx *fiber.Ctx) error {
	var input usecase.CancelOrderInputDTO
	id := ctx.Params("id")
	input.ID = id
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	usecase := usecase.NewCancelOrderUseCase(c.OrderRepository)
	err := usecase.Execute(input)
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *Controller) FindAllOrders(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	usecase := usecase.NewFindAllOrdersUseCase(c.OrderRepository)
	output, err := usecase.Execute()
	if err != nil {
		if err.Error() == "internal Server Error" {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	json, _ := json.Marshal(output)
	return ctx.Status(fiber.StatusOK).Send(json)
}
