package usecase

import (
	"testing"

	"github.com/janapc/order-restaurant/internal/entity"
	"github.com/janapc/order-restaurant/internal/mock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUpdateDish(t *testing.T) {
	repository := mock.NewDishRepositoryInMemory()
	dish, _ := entity.NewDish("bolinha de brigadeiro", "muito bom", 4.0)
	repository.Save(dish)
	input := UpdateDishInputDTO{
		ID:   dish.ID,
		Name: "brigadeiro",
	}
	usecase := NewUpdateDishUseCase(repository)
	err := usecase.Execute(input)
	assert.NoError(t, err)
}
