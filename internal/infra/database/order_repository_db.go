package database

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/janapc/order-restaurant/internal/entity"
)

type OrderDB struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Dishes     []entity.DishOrder `bson:"dishes"`
	Tax        float64            `bson:"tax"`
	TotalPrice float64            `bson:"total_price,omitempty"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"created_at"`
	CanceledAt *time.Time         `bson:"canceled_at,omitempty"`
}

func (o *OrderRepositoryDB) ConvertEntityToDB(order *entity.Order) *OrderDB {
	return &OrderDB{
		Dishes:     order.Dishes,
		Tax:        order.Tax,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		CanceledAt: order.CanceledAt,
	}
}

func (o *OrderRepositoryDB) ConvertDBToEntity(orderDb *OrderDB) entity.Order {
	return entity.Order{
		ID:         orderDb.ID.Hex(),
		Dishes:     orderDb.Dishes,
		Tax:        orderDb.Tax,
		Status:     orderDb.Status,
		TotalPrice: orderDb.TotalPrice,
		CreatedAt:  orderDb.CreatedAt,
		CanceledAt: orderDb.CanceledAt,
	}
}

type OrderRepositoryDB struct {
	DB *mongo.Database
}

func NewOrderRepositoryDB(db *mongo.Database) *OrderRepositoryDB {
	return &OrderRepositoryDB{
		DB: db,
	}
}

var ERROR_ORDER_DB = "internal Server Error"

func (o *OrderRepositoryDB) Save(order *entity.Order) error {
	orderDb := o.ConvertEntityToDB(order)
	collectionName := os.Getenv("COLLECTION_MONGO_ORDERS")
	result, err := o.DB.Collection(collectionName).InsertOne(context.TODO(), orderDb)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_ORDER_DB)
	}
	order.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (o *OrderRepositoryDB) Cancel(id string) error {
	collectionName := os.Getenv("COLLECTION_MONGO_ORDERS")
	idDb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_ORDER_DB)
	}
	filter := bson.D{{Key: "_id", Value: idDb}}
	update := bson.D{{Key: "$set", Value: bson.M{"canceled_at": time.Now(), "status": "CANCELED"}}}
	_, err = o.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_ORDER_DB)
	}
	return nil
}

func (o *OrderRepositoryDB) FindAll() ([]entity.Order, error) {
	collectionName := os.Getenv("COLLECTION_MONGO_ORDERS")
	filter := bson.D{{}}
	cursor, err := o.DB.Collection(collectionName).Find(context.TODO(), filter)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_ORDER_DB)
	}
	var ordersDb []OrderDB
	if err = cursor.All(context.TODO(), &ordersDb); err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_ORDER_DB)
	}
	var orders []entity.Order
	for _, order := range ordersDb {
		orders = append(orders, o.ConvertDBToEntity(&order))
	}
	return orders, nil
}

func (o *OrderRepositoryDB) FindById(id string) (*entity.Order, error) {
	collectionName := os.Getenv("COLLECTION_MONGO_ORDERS")
	idDb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_ORDER_DB)
	}
	filter := bson.D{{Key: "_id", Value: idDb}}
	var order *entity.Order
	err = o.DB.Collection(collectionName).FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New(ERROR_ORDER_DB)
	}
	return order, nil
}

func (o *OrderRepositoryDB) Status(status string, id string) error {
	collectionName := os.Getenv("COLLECTION_MONGO_ORDERS")
	idDb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_ORDER_DB)
	}
	filter := bson.D{{Key: "_id", Value: idDb}}
	update := bson.D{{Key: "$set", Value: bson.M{"status": status}}}
	_, err = o.DB.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Error(err.Error())
		return errors.New(ERROR_ORDER_DB)
	}
	return nil
}
