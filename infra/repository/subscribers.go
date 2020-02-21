package repository

import "github.com/vagner-nascimento/go-poc-archref/infra/data"

func SubscribeAllConsumers() error {
	// TODO: append other consumers to test multiple queue readers
	var subs []data.AmqSubscriber

	subs = append(subs, newUserSubscriber())
	err := data.SubscribeConsumers(subs)

	if err != nil {
		return err
	}

	return nil
}
