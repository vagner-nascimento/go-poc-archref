package integration

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
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
		subs := getAllSubs()
		var subsSuccess []string
		for _, sub := range subs {
			err := amqSub.AddSubscriber(sub.getTopic(), sub.getConsumer(), sub.getHandler())
			if err != nil {
				infra.LogError("error on subscribe consumer", err)
			} else {
				subsSuccess = append(subsSuccess, reflect.TypeOf(sub).Elem().Name())
			}
		}

		if err = amqSub.SubscribeAll(); err == nil {
			infra.LogInfo(fmt.Sprintf("successfully subscribed: %s", strings.Join(subsSuccess, ", ")))
		}
	}
	return err
}

func getAllSubs() (subs []subscription) {
	return append(subs,
		newUserSub(),
		newEnterpriseSub())
}
