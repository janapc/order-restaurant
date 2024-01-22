package mock

import (
	"database/sql"
	"log/slog"

	"github.com/janapc/order-restaurant/internal/entity"
)

type DishRepositoryInMemory struct {
	DB *sql.DB
}

func NewDishRepositoryInMemory() *DishRepositoryInMemory {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		slog.Error(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS dishes(id VARCHAR(150) NOT NULL,name VARCHAR(150) NOT NULL,description VARCHAR(150) NOT NULL,price FLOAT NOT NULL,PRIMARY KEY (id))")
	if err != nil {
		slog.Error(err.Error())
	}
	return &DishRepositoryInMemory{
		DB: db,
	}
}

func (d *DishRepositoryInMemory) Save(dish *entity.Dish) error {
	stmt, err := d.DB.Prepare("INSERT INTO dishes(id,name,description,price) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dish.ID, dish.Name, dish.Description, dish.Price)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepositoryInMemory) Update(dish *entity.Dish) error {
	stmt, err := d.DB.Prepare("UPDATE dishes SET name = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dish.Name, dish.Description, dish.Price, dish.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepositoryInMemory) FindById(id string) (*entity.Dish, error) {
	stmt, err := d.DB.Prepare("SELECT id, name, description, price FROM dishes WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var dish entity.Dish
	err = stmt.QueryRow(id).Scan(&dish.ID, &dish.Name, &dish.Description, &dish.Price)
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (d *DishRepositoryInMemory) Remove(id string) error {
	stmt, err := d.DB.Prepare("DELETE FROM dishes WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepositoryInMemory) FindAll() ([]entity.Dish, error) {
	rows, err := d.DB.Query("SELECT id, name, description, price FROM dishes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dishes []entity.Dish
	for rows.Next() {
		var dish entity.Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Description, &dish.Price)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}
