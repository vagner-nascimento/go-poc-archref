package repository

import "github.com/vagner-nascimento/go-poc-archref/src/infra/data"

func SubscribeAllConsumers() error {
	// TODO: append other consumers to test multiple queue readers
	var subs []data.AmqSubscriber
	subs = append(subs, newUserSubscriber())

	return data.SubscribeConsumers(subs)
}
