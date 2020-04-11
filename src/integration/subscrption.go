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
	amqSub, err := provider.AmqpSubscription()
	if err == nil {
		subs := getSubscriptions()

		var subSuccess []string
		for _, sub := range subs {
			err := amqSub.AddSubscriber(sub.getTopic(), sub.getConsumer(), sub.getHandler())
			if err != nil {
				// TODO: INTEGRATION AMQ SUB: if someone fails error, it will be never reconnected
				// it will happen only if topic name. consumer or handler were empty or null
				logger.Error("error on subscribe consumer", err)
			} else {
				subSuccess = append(subSuccess, reflect.TypeOf(sub).Elem().Name())
			}
		}

		if err = amqSub.SubscribeAll(); err == nil {
			logger.Info(fmt.Sprintf("successfully subscribed: %s", strings.Join(subSuccess, ", ")))
		}
	}

	return err
}

func getSubscriptions() (subs []subscription) {
	return append(subs,
		newUserSub(),
		newEnterpriseSub())
}
