package entity

type DishRepositoryInterface interface {
	Save(dish *Dish) error
	Update(dish *Dish) error
	Remove(id string) error
	FindAll() ([]Dish, error)
	FindById(id string) (*Dish, error)
}

type OrderRepositoryInterface interface {
	Create(order *Order) error
	Cancel(id string) error
	FindAll() ([]Order, error)
}
