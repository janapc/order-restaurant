package usecase

import (
	"testing"

	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDish(t *testing.T) {
	repository := mock.NewDishRepositoryInMemory()
	dish, _ := entity.NewDish("bolinha de brigadeiro", "muito bom", 4.0)
	repository.Save(dish)
	usecase := NewRemoveDishUseCase(repository)
	input := RemoveDishInputDTO{
		ID: dish.ID,
	}
	err := usecase.Execute(input)
	assert.NoError(t, err)
}
