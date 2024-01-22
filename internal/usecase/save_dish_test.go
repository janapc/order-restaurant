package usecase

import (
	"testing"

	"github.com/janapc/order-restaurant/internal/mock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSaveDish(t *testing.T) {
	input := SaveDishInputDTO{
		Name:        "brigadeiro",
		Description: "delicioso brigadeiro artesanal",
		Price:       4.0,
	}
	repository := mock.NewDishRepositoryInMemory()
	usecase := NewSaveDishUseCase(repository)
	output, err := usecase.Execute(input)
	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.NotEmpty(t, output.ID)
}
