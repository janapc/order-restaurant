package usecase

import "github.com/janapc/order-restaurant/internal/entity"

type FindAllOutputDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type FindAllUseCase struct {
	DishRepository entity.DishRepositoryInterface
}

func NewFindAllUseCase(repository entity.DishRepositoryInterface) *FindAllUseCase {
	return &FindAllUseCase{
		DishRepository: repository,
	}
}

func (f *FindAllUseCase) Execute() ([]FindAllOutputDTO, error) {
	dishes, err := f.DishRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var output []FindAllOutputDTO
	for _, dish := range dishes {
		output = append(output, FindAllOutputDTO{
			ID:          dish.ID,
			Name:        dish.Name,
			Description: dish.Description,
			Price:       dish.Price,
		})
	}
	return output, nil
}
