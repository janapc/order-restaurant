package database

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/janapc/order-restaurant/internal/entity"
)

type DishRepositoryDB struct {
	DB *sql.DB
}

func NewDishRepositoryDB(db *sql.DB) *DishRepositoryDB {
	return &DishRepositoryDB{
		DB: db,
	}
}

var ERROR_DISH_DB = "internal Server Error"

func (d *DishRepositoryDB) Save(dish *entity.Dish) error {
	stmt, err := d.DB.Prepare("INSERT INTO dishes(id,name,description,price) VALUES(?,?,?,?)")
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	defer stmt.Close()
	_, err = stmt.Exec(dish.ID, dish.Name, dish.Description, dish.Price)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	return nil
}

func (d *DishRepositoryDB) Update(dish *entity.Dish) error {
	stmt, err := d.DB.Prepare("UPDATE dishes SET name = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	defer stmt.Close()
	_, err = stmt.Exec(dish.Name, dish.Description, dish.Price, dish.ID)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	return nil
}

func (d *DishRepositoryDB) FindById(id string) (*entity.Dish, error) {
	stmt, err := d.DB.Prepare("SELECT id, name, description, price FROM dishes WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("dish is not found")
	}
	defer stmt.Close()
	var dish entity.Dish
	err = stmt.QueryRow(id).Scan(&dish.ID, &dish.Name, &dish.Description, &dish.Price)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_DISH_DB)
	}
	return &dish, nil
}

func (d *DishRepositoryDB) Remove(id string) error {
	stmt, err := d.DB.Prepare("DELETE FROM dishes WHERE id = ?")
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_DISH_DB)
	}
	return nil
}

func (d *DishRepositoryDB) FindAll() ([]entity.Dish, error) {
	rows, err := d.DB.Query("SELECT id, name, description, price FROM dishes")
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_DISH_DB)
	}
	defer rows.Close()
	var dishes []entity.Dish
	for rows.Next() {
		var dish entity.Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Description, &dish.Price)
		if err != nil {
			slog.Error(err.Error())
			return nil, errors.New(ERROR_DISH_DB)
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}
