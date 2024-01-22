package main

import (
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janapc/order-restaurant/internal/infra/database"
	"github.com/janapc/order-restaurant/internal/infra/webserver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		slog.Error(err.Error())
	}
	mysqlDB, err := database.ConnectionMysql()
	if err != nil {
		panic(err.Error())
	}
	mongoDB, err := database.ConnectionMongo()
	if err != nil {
		panic(err.Error())
	}
	defer mysqlDB.Close()
	dishRepository := database.NewDishRepositoryDB(mysqlDB)
	orderRepository := database.NewOrderRepositoryDB(mongoDB)
	server := webserver.NewWebServer(dishRepository, orderRepository)
	server.Start()
}
