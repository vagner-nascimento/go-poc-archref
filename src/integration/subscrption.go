package integration

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"reflect"
	"strings"
)

func SubscribeConsumers() error {
	var subs []subscription
	subs = append(subs, newUserSub()) // TODO: test with more subs
	amqSub := provider.AmqpSubscription()

	var subsSuccess []string
	for _, sub := range subs {
		err := amqSub.AddSubscriber(sub.getTopic(), sub.getConsumer(), sub.getHandler())
		if err != nil {
			infra.LogError("error on subscribe consumer", err)
		} else {
			subsSuccess = append(subsSuccess, reflect.TypeOf(sub).Elem().Name())
		}
	}

	if err := amqSub.SubscribeAll(); err != nil {
		return err
	}

	infra.LogInfo(fmt.Sprintf("successfully subscribed: %s", strings.Join(subsSuccess, ",")))
	return nil
}
