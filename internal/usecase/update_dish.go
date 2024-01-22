package usecase

import (
	"github.com/janapc/order-restaurant/internal/entity"
)

type UpdateDishInputDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type UpdateDishUseCase struct {
	DishRepository entity.DishRepositoryInterface
}

func NewUpdateDishUseCase(repository entity.DishRepositoryInterface) *UpdateDishUseCase {
	return &UpdateDishUseCase{
		DishRepository: repository,
	}
}

func (u *UpdateDishUseCase) Execute(input UpdateDishInputDTO) error {
	dish, err := u.DishRepository.FindById(input.ID)
	if err != nil {
		return err
	}
	if input.Name != "" {
		dish.Name = input.Name
	}
	if input.Description != "" {
		dish.Description = input.Description
	}
	if input.Price != 0.0 {
		dish.Price = input.Price
	}
	err = u.DishRepository.Update(dish)
	if err != nil {
		return err
	}
	return nil
}
