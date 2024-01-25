package entity

type DishRepositoryInterface interface {
	Save(dish *Dish) error
	Update(dish *Dish) error
	Remove(id string) error
	FindAll() ([]Dish, error)
	FindById(id string) (*Dish, error)
}

type OrderRepositoryInterface interface {
	Save(order *Order) error
	Cancel(id string) error
	FindAll() ([]Order, error)
	FindById(id string) (*Order, error)
	Status(status string, id string) error
}
