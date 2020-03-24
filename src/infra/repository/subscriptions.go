package repository

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type amqpSubscription struct {
	subscribers []data.RabbitSubscriber
}

func (rs *amqpSubscription) AddSubscriber(topicName string, consumerName string, messageHandler func(data []byte)) error {
	if topicName == "" || consumerName == "" || messageHandler == nil {
		return errors.New(fmt.Sprintf("invalid subscriber's data: topic: %s, consumer: %s, handler is nil: %t",
			topicName,
			consumerName,
			messageHandler == nil))
	}
	rs.subscribers = append(rs.subscribers, data.NewRabbitSubscriber(topicName, consumerName, messageHandler))
	return nil
}

func (rs *amqpSubscription) SubscribeAll() (err error) {
	if len(rs.subscribers) > 0 {
		err = data.SubscribeRabbitConsumers(rs.subscribers)
	} else {
		err = errors.New("there are no subscribers to consume topics")
	}
	return err
}

func NewAmqpSubscription() AmqpSubscriptionHandler {
	return &amqpSubscription{}
}
