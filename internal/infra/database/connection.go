package database

import (
	"context"
	"database/sql"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectionMongo() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URL")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database("order_restaurant"), nil
}

func ConnectionMysql() (*sql.DB, error) {
	mysqlUrl := os.Getenv("MYSQL_URL")
	database, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return nil, err
	}
	return database, nil
}
