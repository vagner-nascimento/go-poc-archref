package integration

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"reflect"
	"strings"
)

type subscription interface {
	getTopic() string
	getConsumer() string
	getHandler() func([]byte)
}

func SubscribeConsumers() (err error) {
	if amqSub, err := provider.AmqpSubscription(); err == nil {
		subs := getSubscriptions()
		var subsSuccess []string
		for _, sub := range subs {
			err := amqSub.AddSubscriber(sub.getTopic(), sub.getConsumer(), sub.getHandler())
			if err != nil {
				logger.Error("error on subscribe consumer", err)
			} else {
				subsSuccess = append(subsSuccess, reflect.TypeOf(sub).Elem().Name())
			}
		}
		if err = amqSub.SubscribeAll(); err == nil {
			logger.Info(fmt.Sprintf("successfully subscribed: %s", strings.Join(subsSuccess, ", ")))
		}
	}
	return err
}

func getSubscriptions() (subs []subscription) {
	return append(subs,
		newUserSub(),
		newEnterpriseSub())
}
