package main

import (
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/janapc/order-restaurant/internal/infra/database"
	"github.com/janapc/order-restaurant/internal/infra/queue"
	"github.com/janapc/order-restaurant/internal/infra/webserver"
	"github.com/janapc/order-restaurant/internal/usecase"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		slog.Error(err.Error())
	}

	mysqlDB := database.ConnectionMysql()
	mongoDB := database.ConnectionMongo()
	defer mysqlDB.Close()

	dishRepository := database.NewDishRepositoryDB(mysqlDB)
	mongoDBName := os.Getenv("MONGO_DATABASE")
	orderRepository := database.NewOrderRepositoryDB(mongoDB.Database(mongoDBName))

	conn, ch := queue.ConnectionRabbitMQ()
	defer conn.Close()
	defer ch.Close()
	outCh := make(chan amqp.Delivery)
	rabbitMq := queue.NewRabbitmq(ch, outCh)
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	usecaseProcessOrder := usecase.NewProcessOrder(orderRepository)
	go rabbitMq.Consumer(queueName, "")
	go Worker(outCh, usecaseProcessOrder)

	server := webserver.NewWebServer(dishRepository, orderRepository, rabbitMq)
	server.Start()
}

func Worker(msgCh chan amqp.Delivery, uc *usecase.ProcessOrder) {
	for msg := range msgCh {
		err := uc.Execute(msg.Body)
		if err == nil {
			msg.Ack(false)
		}
	}
}
