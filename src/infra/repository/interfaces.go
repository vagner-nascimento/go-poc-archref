package repository

type AmqpSubscriptionHandler interface {
	AddSubscriber(topicName string, consumerName string, messageHandler func(data []byte)) error
	SubscribeAll() error
}
