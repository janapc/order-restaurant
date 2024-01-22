package database

import (
	"database/sql"
	"log/slog"
	"testing"

	"github.com/janapc/order-restaurant/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func connectDBInMemory() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		slog.Error(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS dishes(id VARCHAR(150) NOT NULL,name VARCHAR(150) NOT NULL,description VARCHAR(150) NOT NULL,price FLOAT NOT NULL,PRIMARY KEY (id))")
	if err != nil {
		slog.Error(err.Error())
	}
	return db
}

func TestSaveDishRepository(t *testing.T) {
	input, _ := entity.NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	db := connectDBInMemory()
	repository := NewDishRepositoryDB(db)
	err := repository.Save(input)
	assert.NoError(t, err)
}

func TestFindByIdDishRepository(t *testing.T) {
	input, _ := entity.NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	db := connectDBInMemory()
	repository := NewDishRepositoryDB(db)
	_ = repository.Save(input)
	dish, _ := repository.FindById(input.ID)
	assert.NotEmpty(t, dish)
	assert.Equal(t, dish.ID, input.ID)
}

func TestUpdateDishRepository(t *testing.T) {
	input, _ := entity.NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	db := connectDBInMemory()
	repository := NewDishRepositoryDB(db)
	_ = repository.Save(input)
	input.Name = "Brigadeiro 2"
	err := repository.Update(input)
	assert.NoError(t, err)
	dish, _ := repository.FindById(input.ID)
	assert.NotEmpty(t, dish)
	assert.Equal(t, dish.Name, "Brigadeiro 2")
}

func TestRemoveDishRepository(t *testing.T) {
	input, _ := entity.NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	db := connectDBInMemory()
	repository := NewDishRepositoryDB(db)
	_ = repository.Save(input)
	err := repository.Remove(input.ID)
	assert.NoError(t, err)
	dish, _ := repository.FindById(input.ID)
	assert.Empty(t, dish)
}

func TestFindAllDishesRepository(t *testing.T) {
	input, _ := entity.NewDish("brigadeiro", "delicioso brigadeiro artesanal", 4.0)
	input2, _ := entity.NewDish("torta", "deliciosa torta artesanal", 8.0)
	db := connectDBInMemory()
	repository := NewDishRepositoryDB(db)
	_ = repository.Save(input)
	_ = repository.Save(input2)
	dishes, err := repository.FindAll()
	assert.NoError(t, err)
	assert.Len(t, dishes, 2)
}
