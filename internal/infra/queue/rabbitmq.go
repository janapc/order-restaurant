package queue

import (
	"context"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	Ch  *amqp.Channel
	Out chan amqp.Delivery
}

func NewRabbitmq(ch *amqp.Channel, out chan amqp.Delivery) *Rabbitmq {
	return &Rabbitmq{
		Ch:  ch,
		Out: out,
	}
}

func ConnectionRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	connectionUrl := os.Getenv("RABBITMQ_CONNECTION")
	conn, err := amqp.Dial(connectionUrl)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return conn, ch
}

func (r *Rabbitmq) Consumer(queueName string, exchange string) error {
	messages, err := r.Ch.Consume(queueName, exchange, false, false, false, false, nil)
	if err != nil {
		return err
	}
	for message := range messages {
		r.Out <- message
	}
	return nil
}

func (r *Rabbitmq) Publish(body string, exchange string, queue string) error {
	_, err := r.Ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = r.Ch.PublishWithContext(ctx,
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}
