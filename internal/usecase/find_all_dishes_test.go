package usecase

import (
	"testing"

	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestFindAllDishes(t *testing.T) {
	repository := mock.NewDishRepositoryInMemory()
	dish, _ := entity.NewDish("bolinha de brigadeiro", "muito bom", 4.0)
	dish2, _ := entity.NewDish("torta de maça", "muito bom", 5.0)
	repository.Save(dish)
	repository.Save(dish2)
	usecase := NewFindAllUseCase(repository)
	dishes, err := usecase.Execute()
	assert.NoError(t, err)
	assert.Len(t, dishes, 2)
	assert.Equal(t, dishes[0].Name, dish.Name)
	assert.Equal(t, dishes[1].Name, dish2.Name)
}
