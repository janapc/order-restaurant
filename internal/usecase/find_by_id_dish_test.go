package usecase

import (
	"testing"

	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestFindByIdDish(t *testing.T) {
	repository := mock.NewDishRepositoryInMemory()
	input, _ := entity.NewDish("bolinha de brigadeiro", "muito bom", 4.0)
	repository.Save(input)
	usecase := NewFindByIdUseCase(repository)
	dish, err := usecase.Execute(FindByIdInputDTO{ID: input.ID})
	assert.NoError(t, err)
	assert.NotEmpty(t, dish)
	assert.Equal(t, dish.ID, input.ID)
}
