package database

import (
	"context"
	"database/sql"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectionMongo() *mongo.Client {
	uri := os.Getenv("MONGO_URL")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

func ConnectionMysql() *sql.DB {
	mysqlUrl := os.Getenv("MYSQL_URL")
	database, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	return database
}
