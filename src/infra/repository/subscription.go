package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type amqpSubscription struct {
	amqp data.AmqpHandler
}

func (rs *amqpSubscription) AddSubscriber(topicName string, consumerName string, messageHandler func(data []byte)) error {
	return rs.amqp.AddSubscriber(topicName, consumerName, messageHandler)
}

func (rs *amqpSubscription) SubscribeAll() (err error) {
	return rs.amqp.SubscribeAll()
}

func NewAmqpSubscription() (AmqpSubscriptionHandler, error) {
	handler, err := data.NewAmqpHandler()
	if err != nil {
		return nil, err
	}
	return &amqpSubscription{amqp: handler}, nil
}
