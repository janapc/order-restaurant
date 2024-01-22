package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janapc/order-restaurant/internal/entity"
)

type WebServer struct {
	DishRepository  entity.DishRepositoryInterface
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebServer(dishRepository entity.DishRepositoryInterface,
	orderRepository entity.OrderRepositoryInterface) *WebServer {
	return &WebServer{
		DishRepository:  dishRepository,
		OrderRepository: orderRepository,
	}
}

func (w *WebServer) Start() {
	controller := NewController(w.DishRepository, w.OrderRepository)
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
