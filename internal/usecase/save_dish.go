package usecase

import (
	"github.com/janapc/order-restaurant/internal/entity"
)

type SaveDishInputDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type SaveDishOutputDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type SaveDishUseCase struct {
	DishRepository entity.DishRepositoryInterface
}

func NewSaveDishUseCase(repository entity.DishRepositoryInterface) *SaveDishUseCase {
	return &SaveDishUseCase{
		DishRepository: repository,
	}
}

func (s *SaveDishUseCase) Execute(input SaveDishInputDTO) (*SaveDishOutputDTO, error) {
	dish, err := entity.NewDish(input.Name, input.Description, input.Price)
	if err != nil {
		return nil, err
	}
	err = s.DishRepository.Save(dish)
	if err != nil {
		return nil, err
	}
	return &SaveDishOutputDTO{
		ID:          dish.ID,
		Name:        dish.Name,
		Description: dish.Description,
		Price:       dish.Price,
	}, nil
}
