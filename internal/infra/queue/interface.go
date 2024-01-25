package queue

type QueueInterface interface {
	Consumer(queueName string, exchange string) error
	Publish(body string, exchange string, queue string) error
}
