package repository

import "github.com/vagner-nascimento/go-poc-archref/infra/data"

func SubscribeAllConsumers() error {
	var subs []data.AmqSubscriber

	subs = append(subs, newCustomerSubscriber())
	// TODO: append other consumers to test multiple queue readers

	err := data.SubscribeConsumers(subs)

	if err != nil {
		return err
	}

	return nil
}
