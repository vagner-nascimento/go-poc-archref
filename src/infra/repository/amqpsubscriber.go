package repository

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
)

type Subscription interface {
	GetTopic() string
	GetConsumer() string
	GetHandler() func([]byte)
}

// SubscribeConsumers - subscribes the consumers into amqp server and retry subscribe if connection gets down
// while it is not lost forever (connection is lost forever when cannot reconnect on retry parameters limits)
func SubscribeConsumers(subs []Subscription) error {
	if err := subscribeConsumers(subs); err != nil {
		return err
	}

	go func(subs []Subscription) {
		connStatus := make(chan bool)
		if err := data.ListenToRabbitConnectionStatus(&connStatus); err != nil {
			fmt.Println("error on listen to amqp connection status")
			return
		}

		for isConnected := range connStatus {
			if !isConnected {
				subscribeAllWhenReestablishConnection(&connStatus, subs)
			}
		}

		return
	}(subs)

	return nil
}

func subscribeConsumers(subs []Subscription) error {
	subsFailed := 0
	for _, sub := range subs {
		if err := data.SubscribeConsumer(sub.GetTopic(), sub.GetConsumer(), sub.GetHandler()); err != nil {
			logger.Error(fmt.Sprintf("error on subscribe consumer %s", sub.GetConsumer()), err)
			subsFailed = subsFailed + 1
		} else {
			logger.Info(fmt.Sprintf("consumer %s subscried on topic %s", sub.GetConsumer(), sub.GetTopic()))
		}
	}

	if subsFailed == len(subs) {
		return errors.New("all subscriptions failed to consume topics")
	}

	return nil
}

// TODO: Realise why dont subscribe again on reconnect
func subscribeAllWhenReestablishConnection(connStatus *chan bool, subs []Subscription) {
	for isConnected := range *connStatus {
		if isConnected {
			if err := subscribeConsumers(subs); err != nil {
				logger.Error("error try to re-subscribe consumers", err)
			}
		}
	}
}
