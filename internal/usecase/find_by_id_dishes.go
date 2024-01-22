package usecase

import (
	"github.com/janapc/order-restaurant/internal/entity"
)

type FindByIdInputDTO struct {
	ID string `json:"id"`
}

type FindByIdOutputDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type FindByIdUseCase struct {
	DishRepository entity.DishRepositoryInterface
}

func NewFindByIdUseCase(repository entity.DishRepositoryInterface) *FindByIdUseCase {
	return &FindByIdUseCase{
		DishRepository: repository,
	}
}

func (f *FindByIdUseCase) Execute(input FindByIdInputDTO) (*FindByIdOutputDTO, error) {
	dish, err := f.DishRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	return &FindByIdOutputDTO{
		ID:          dish.ID,
		Name:        dish.Name,
		Description: dish.Description,
		Price:       dish.Price,
	}, nil
}
