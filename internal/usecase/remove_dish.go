package usecase

import (
	"github.com/janapc/order-restaurant/internal/entity"
)

type RemoveDishInputDTO struct {
	ID string `json:"id"`
}

type RemoveDishUseCase struct {
	DishRepository entity.DishRepositoryInterface
}

func NewRemoveDishUseCase(repository entity.DishRepositoryInterface) *RemoveDishUseCase {
	return &RemoveDishUseCase{
		DishRepository: repository,
	}
}

func (r *RemoveDishUseCase) Execute(input RemoveDishInputDTO) error {
	_, err := r.DishRepository.FindById(input.ID)
	if err != nil {
		return err
	}
	err = r.DishRepository.Remove(input.ID)
	if err != nil {
		return err
	}
	return nil
}
