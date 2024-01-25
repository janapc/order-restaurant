package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/infra/queue"
)

type WebServer struct {
	DishRepository  entity.DishRepositoryInterface
	OrderRepository entity.OrderRepositoryInterface
	Queue           queue.QueueInterface
}

func NewWebServer(dishRepository entity.DishRepositoryInterface,
	orderRepository entity.OrderRepositoryInterface,
	queue queue.QueueInterface) *WebServer {
	return &WebServer{
		DishRepository:  dishRepository,
		OrderRepository: orderRepository,
		Queue:           queue,
	}
}

func (w *WebServer) Start() {
	controller := NewController(w.DishRepository, w.OrderRepository, w.Queue)
	app := fiber.New()
	apiDish := app.Group("/api/dish")
	apiDish.Post("/", controller.RegisterDish)
	apiDish.Patch("/:id", controller.UpdateDish)
	apiDish.Delete("/:id", controller.RemoveDish)
	apiDish.Get("/:id", controller.FindByIdDish)
	apiDish.Get("/", controller.FindAllDish)

	apiOrder := app.Group("/api/order")
	apiOrder.Post("/", controller.CreateOrder)
	apiOrder.Delete("/:id", controller.CancelOrder)
	apiOrder.Get("/", controller.FindAllOrders)

	app.Listen(":3000")
}
